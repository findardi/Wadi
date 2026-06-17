<script lang="ts">
	import { t } from '$lib/i18n';

	type Provider = 'google' | 'github';

	let pending = $state<Provider | null>(null);

	const start = (p: Provider) => () => {
		pending = p;
	};
</script>

<div class="flex justify-center gap-3" class:pointer-events-none={pending !== null}>
	<a
		href="/auth/sso/google"
		data-sveltekit-reload
		class="btn btn-square btn-outline size-10 border-base-300"
		class:opacity-60={pending !== null && pending !== 'google'}
		aria-label={t('sso.google')}
		aria-busy={pending === 'google'}
		title={t('sso.google')}
		onclick={start('google')}
	>
		{#if pending === 'google'}
			<span class="loading loading-spinner loading-sm"></span>
		{:else}
			<svg width="18" height="18" viewBox="0 0 18 18" aria-hidden="true">
				<path
					fill="#4285F4"
					d="M17.64 9.2c0-.64-.06-1.25-.16-1.84H9v3.48h4.84a4.14 4.14 0 0 1-1.8 2.72v2.26h2.92c1.71-1.57 2.68-3.88 2.68-6.62z"
				/>
				<path
					fill="#34A853"
					d="M9 18c2.43 0 4.47-.8 5.96-2.18l-2.92-2.26c-.81.54-1.84.86-3.04.86-2.34 0-4.32-1.58-5.03-3.7H.96v2.33A9 9 0 0 0 9 18z"
				/>
				<path
					fill="#FBBC05"
					d="M3.97 10.72a5.4 5.4 0 0 1 0-3.44V4.95H.96a9 9 0 0 0 0 8.1l3.01-2.33z"
				/>
				<path
					fill="#EA4335"
					d="M9 3.58c1.32 0 2.5.45 3.44 1.35l2.58-2.59C13.46.89 11.43 0 9 0A9 9 0 0 0 .96 4.95l3.01 2.34C4.68 5.16 6.66 3.58 9 3.58z"
				/>
			</svg>
		{/if}
	</a>

	<a
		href="/auth/sso/github"
		data-sveltekit-reload
		class="btn btn-square btn-outline size-10 border-base-300"
		class:opacity-60={pending !== null && pending !== 'github'}
		aria-label={t('sso.github')}
		aria-busy={pending === 'github'}
		title={t('sso.github')}
		onclick={start('github')}
	>
		{#if pending === 'github'}
			<span class="loading loading-spinner loading-sm"></span>
		{:else}
			<svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
				<path
					d="M12 .5C5.37.5 0 5.87 0 12.5c0 5.3 3.44 9.8 8.21 11.39.6.11.82-.26.82-.58 0-.29-.01-1.05-.02-2.06-3.34.73-4.04-1.61-4.04-1.61-.55-1.39-1.34-1.76-1.34-1.76-1.09-.75.08-.73.08-.73 1.21.09 1.84 1.24 1.84 1.24 1.07 1.84 2.81 1.31 3.5 1 .11-.78.42-1.31.76-1.61-2.67-.3-5.47-1.33-5.47-5.93 0-1.31.47-2.38 1.24-3.22-.13-.3-.54-1.52.12-3.17 0 0 1.01-.32 3.3 1.23a11.5 11.5 0 0 1 6.01 0c2.29-1.55 3.3-1.23 3.3-1.23.66 1.65.25 2.87.12 3.17.77.84 1.23 1.91 1.23 3.22 0 4.61-2.81 5.62-5.49 5.92.43.37.81 1.1.81 2.22 0 1.6-.01 2.89-.01 3.29 0 .32.21.7.82.58A12.01 12.01 0 0 0 24 12.5C24 5.87 18.63.5 12 .5z"
				/>
			</svg>
		{/if}
	</a>
</div>

<!-- Announce the redirect for screen readers (visual users see the spinner). -->
<p class="sr-only" role="status" aria-live="polite">{pending ? t('sso.redirecting') : ''}</p>
