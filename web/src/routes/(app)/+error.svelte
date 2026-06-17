<script lang="ts">
	import { page } from '$app/state';
	import { t } from '$lib/i18n';
</script>

<svelte:head><title>{page.status} · {t('brand.name')}</title></svelte:head>

<div class="grid min-h-full place-items-center px-6 py-16">
	<section class="flex max-w-md flex-col items-center text-center">
		<p class="font-mono text-sm text-muted">{page.status}</p>
		<h1 class="mt-2 text-[1.375rem] font-semibold tracking-[-0.015em] text-balance">
			{page.error?.message ?? t('error.title')}
		</h1>
		<a href="/" class="btn btn-primary mt-6">{t('error.home')}</a>
	</section>
</div>
