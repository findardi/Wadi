import { redirect } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { getSSOAuthUrl, isSSOProvider } from '$lib/server/api';
import { setOAuthState } from '$lib/server/oauth-state';

// Mint a CSRF `state`, persist it httpOnly, then bounce to the provider consent screen.
export const GET: RequestHandler = async ({ params, cookies }) => {
	if (!isSSOProvider(params.provider)) redirect(303, '/login?sso_error=provider');

	const state = crypto.randomUUID();
	const res = await getSSOAuthUrl(params.provider, state);
	if (!res) redirect(303, '/login?sso_error=unavailable');

	setOAuthState(cookies, state);
	redirect(302, res.url);
};
