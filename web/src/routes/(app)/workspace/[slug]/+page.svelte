<script lang="ts">
	import { t } from '$lib/i18n';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const ws = $derived(data.workspace);

	const dateFmt = new Intl.DateTimeFormat('id-ID', { day: 'numeric', month: 'short', year: 'numeric' });
	const fmtDate = (iso: string) => dateFmt.format(new Date(iso));
</script>

<svelte:head><title>{ws.name} · {t('brand.name')}</title></svelte:head>

<div class="mx-auto w-full max-w-3xl px-6 py-8">
	<header>
		<h1 class="truncate text-2xl font-semibold tracking-[-0.02em]">{ws.name}</h1>
		<p
			class="mt-1.5 flex flex-wrap items-center gap-x-2 gap-y-0.5 font-mono text-xs text-muted"
		>
			<span>{ws.slug}</span>
			<span aria-hidden="true">·</span>
			<span>{t('ws.detail.created')} {fmtDate(ws.created_at)}</span>
			<span aria-hidden="true">·</span>
			<span>{t('ws.detail.updated')} {fmtDate(ws.updated_at)}</span>
		</p>
	</header>

	{#if ws.description}
		<p class="mt-5 max-w-[65ch] text-[0.9375rem] leading-relaxed text-muted text-pretty">
			{ws.description}
		</p>
	{/if}
</div>
