<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import Trash2Icon from '@lucide/svelte/icons/trash-2';
	import {
		Table,
		TableBody,
		TableCell,
		TableHead,
		TableHeader,
		TableRow
	} from '$lib/components/ui/table';
	import { onMount } from 'svelte';
	import { apiData, deckNamesStore } from './store';

	function deleteDeck(deckName: string) {
		console.log(`Delete deck: ${deckName}`);
	}

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
			{#if $deckNamesStore.length > 0}
				<Table>
					<TableHeader>
						<TableRow>
							<TableHead>Deck</TableHead>
							<TableHead class="text-right">Cards</TableHead>
							<TableHead class="text-right">Actions</TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						{#each $deckNamesStore as deckName (deckName)}
							<TableRow>
								<TableCell class="font-medium">{deckName[0]}</TableCell>
								<TableCell class="text-right">{deckName[1]}</TableCell>
								<TableCell class="text-right">
									<Button
										variant="destructive"
										size="icon-sm"
										aria-label={`Delete ${deckName[0]}`}
										onclick={() => deleteDeck(deckName[0])}
									>
										<Trash2Icon />
									</Button>
								</TableCell>
							</TableRow>
						{/each}
					</TableBody>
				</Table>
			{:else}
				<p>No decks available</p>
			{/if}

			<Button class="rounded-full px-5" type="submit" href="./create_deck">Add deck</Button>
		</div>
	</div>
</div>
