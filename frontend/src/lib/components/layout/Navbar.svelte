<script lang="ts">
	import { resolve } from '$app/paths';
	import type { Pathname } from '$app/types';
	import { HomeIcon, PlayingCardsFan, type Icon as IconType } from '@lucide/svelte';
	const items = $derived<
		{
			name: string;
			url: Pathname;
			activeMatch: (pathname: string) => boolean;
			icon: typeof IconType;
		}[]
	>([
		{
			name: 'Home',
			url: '/',
			activeMatch: (pathname) => pathname === '/',
			icon: HomeIcon
		},
		{
			name: 'Decks',
			url: '/decks',
			activeMatch: (pathname) => pathname === '/decks',
			icon: PlayingCardsFan
		}
	]);

	let links = $state<HTMLAnchorElement[]>([]);
	let container = $state<HTMLDivElement>();
</script>

{#snippet link(item: (typeof items)[0], index: number)}
	{@const Icon = item.icon}
	<a bind:this={links[index]} href={resolve(item.url)} aria-label={item.name} class="">
		<div class="grid grid-cols-1 grid-rows-1">
			{#key item.name}
				<div
					class="col-start-1 row-start-1 mr-2 flex h-full w-full items-center justify-center gap-3"
				>
					<div class="relative">
						<Icon />
					</div>
					<p class="hidden min-w-0 font-medium break-all hyphens-auto md:flex">
						{item.name}
					</p>
				</div>
			{/key}
		</div>
	</a>
{/snippet}

<div bind:this={container} class="flex min-h-10 items-center gap-6 bg-gray-500">
	{#each items as item, i (item.url)}
		{@render link(item, i)}
	{/each}
</div>
