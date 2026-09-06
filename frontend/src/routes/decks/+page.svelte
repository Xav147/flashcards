<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { onMount } from 'svelte';
	import { apiData, deckNamesStore } from './store';
	onMount(async () => {
		fetch('http://localhost:8010/proxy/list_decks')
			.then((response) => response.json())
			.then((data) => {
				console.log(data);
				apiData.set(data);
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<svelte:head>
	<title>Flashcards Home</title>
</svelte:head>

<div class="min-h-screen bg-stone-100 px-6 py-16 text-stone-900">
	<div class="mx-auto flex max-w-2xl flex-col gap-8">
		<div class="space-y-3">
			<p class="text-sm font-semibold tracking-[0.24em] text-stone-500 uppercase">Flashcards</p>
			<h1 class="text-4xl font-bold tracking-tight text-stone-950">Available Decks</h1>
			<p class="max-w-xl text-base text-stone-600">Here you can see all of the available decks</p>
		</div>

		<div class="space-y-5 rounded-3xl border border-stone-200 bg-white p-6 shadow-sm">
			<ul>
				{#each $deckNamesStore as deckName (deckName)}
					<li>{deckName[0]} - Cards: {deckName[1]}</li>
				{/each}
			</ul>

			<Button class="rounded-full px-5" type="submit">Add deck</Button>
		</div>
	</div>
</div>
