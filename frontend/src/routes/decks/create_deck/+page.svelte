<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { toast } from 'svelte-sonner';

	let deckName = $state('');

	async function createDeck(event: SubmitEvent) {
		const endpoint = 'http://localhost:8010/proxy/create_deck';
		event.preventDefault();

		const deck = {
			name: deckName.trim(),
			size: 0
		};

		try {
			const response = await fetch(endpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(deck)
			});

			if (response.status === 409) {
				toast.warning('A deck with this name already exists.');
				return;
			}

			if (!response.ok) {
				throw new Error(`Response status: ${response.status}`);
			}

			toast.success('Deck created successfully.');
		} catch (error: unknown) {
			toast.error('Could not create deck. Please try again.');

			if (error instanceof Error) {
				console.error(error.message);
			} else {
				console.error(error);
			}
		}
	}
</script>

<svelte:head>
	<title>Create Deck</title>
</svelte:head>

<div class="min-h-screen bg-stone-100 px-6 py-16 text-stone-900">
	<div class="mx-auto flex max-w-2xl flex-col gap-8">
		<div class="space-y-3">
			<p class="text-sm font-semibold tracking-[0.24em] text-stone-500 uppercase">Flashcards</p>
			<h1 class="text-4xl font-bold tracking-tight text-stone-950">Create Deck</h1>
			<p class="max-w-xl text-base text-stone-600">Add the basics for a new flashcard deck.</p>
		</div>

		<form
			class="space-y-5 rounded-3xl border border-stone-200 bg-white p-6 shadow-sm"
			onsubmit={createDeck}
		>
			<div class="space-y-2">
				<label class="text-sm font-medium text-stone-700" for="deck-name">Deck name</label>
				<input
					id="deck-name"
					class="w-full rounded-xl border border-stone-300 px-4 py-3 text-base transition outline-none focus:border-stone-500 focus:ring-2 focus:ring-stone-200"
					type="text"
					bind:value={deckName}
					placeholder="Italian verbs"
					required
				/>
			</div>

			<div class="flex items-center justify-between gap-3">
				<Button class="rounded-full px-5" href="./list_decks" variant="outline">Cancel</Button>
				<Button class="rounded-full px-5" type="submit" disabled={!deckName.trim()}>
					Create deck
				</Button>
			</div>
		</form>
	</div>
</div>
