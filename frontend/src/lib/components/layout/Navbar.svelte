<script lang="ts">
	import { HomeIcon, PlayingCardsFan, type Icon as IconType } from '@lucide/svelte';
	const items = $derived<
		{
			name: string;
			url: string;
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
	<a
		bind:this={links[index]}
		href={item.url}
		aria-label={item.name}
		class=""
	>
		<div class="grid grid-rows-1 grid-cols-1">
			{#key item.name}
				<div
					class="mr-2 w-full row-start-1 col-start-1 h-full flex items-center justify-center gap-3"
				>
					<div class="relative">
						<Icon />
					</div>
					<p
						class="font-medium hidden hyphens-auto break-all md:flex min-w-0"
					>
						{item.name}
					</p>
				</div>
			{/key}
		</div>
	</a>
{/snippet}

<div bind:this={container} class="flex items-center gap-6 bg-gray-500 min-h-10">
    {#each items as item, i (item.url)}
        {@render link(item, i)}
    {/each}
</div>
