import { redirect } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { exchangeSSO, isSSOProvider } from '$lib/server/api';
import { takeOAuthState } from '$lib/server/oauth-state';
import { setSession } from '$lib/server/session';

// Provider redirected back with ?code&state. Verify state (CSRF), exchange, open session.
export const GET: RequestHandler = async ({ params, url, cookies }) => {
	if (!isSSOProvider(params.provider)) redirect(303, '/login?sso_error=provider');

	const expected = takeOAuthState(cookies);
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	// User declined consent at the provider — intentional, not a failure. Return quietly.
	if (url.searchParams.get('error')) redirect(303, '/login?sso_error=cancelled');

	// Missing code or CSRF state mismatch — a real failure.
	if (!code || !state || !expected || state !== expected) {
		redirect(303, '/login?sso_error=state');
	}

	const res = await exchangeSSO(params.provider, code);
	if (!res.ok) {
		redirect(303, `/login?sso_error=${res.status === 409 ? 'conflict' : 'failed'}`);
	}

	setSession(cookies, res.data);
	redirect(303, '/');
};
