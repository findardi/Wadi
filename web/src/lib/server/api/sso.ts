import type { LoginData } from '$lib/types';
import { API_URL, get, post } from './client';

const PROVIDERS = ['google', 'github'] as const;
export type SSOProvider = (typeof PROVIDERS)[number];

export function isSSOProvider(p: string): p is SSOProvider {
	return (PROVIDERS as readonly string[]).includes(p);
}

// BFF step 1: we own the CSRF `state`, the backend only builds the authorize URL.
export async function getSSOAuthUrl(
	provider: SSOProvider,
	state: string
): Promise<{ url: string } | null> {
	if (!API_URL) return null; // no backend wired (stub mode) → SSO unavailable
	const res = await get<{ url: string }>(
		`/auth/sso/${provider}/url?state=${encodeURIComponent(state)}`
	);
	return res.ok ? res.data : null;
}

// BFF step 2: swap the provider `code` for our own session tokens.
// 409 = email already on a local (password) account — no auto-link (backend policy).
export async function exchangeSSO(
	provider: SSOProvider,
	code: string
): Promise<{ ok: true; data: LoginData } | { ok: false; status: number }> {
	const res = await post<LoginData>(`/auth/sso/${provider}/exchange`, { code });
	return res.ok ? { ok: true, data: res.data } : { ok: false, status: res.status };
}
