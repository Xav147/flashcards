<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';

	let name = $state('');
	let content = $state('');
	let submittedCard = $state<{ name: string; content: string } | null>(null);

	function handleSubmit(event: SubmitEvent) {
		event.preventDefault();
		submittedCard = {
			name: name.trim(),
			content: content.trim()
		};
	}
</script>

<svelte:head>
	<title>Flashcards Home</title>
</svelte:head>

<div class="min-h-screen bg-stone-100 px-6 py-16 text-stone-900">
	<div class="mx-auto flex max-w-2xl flex-col gap-8">
		<div class="space-y-3">
			<p class="text-sm font-semibold tracking-[0.24em] text-stone-500 uppercase">Flashcards</p>
			<h1 class="text-4xl font-bold tracking-tight text-stone-950">Create a card</h1>
			<p class="max-w-xl text-base text-stone-600">
				Add a simple name and some content to start drafting a new flashcard.
			</p>
		</div>

		<form
			class="space-y-5 rounded-3xl border border-stone-200 bg-white p-6 shadow-sm"
			onsubmit={handleSubmit}
		>
			<div class="space-y-2">
				<label class="text-sm font-medium text-stone-700" for="card-name">Name</label>
				<input
					id="card-name"
					name="name"
					bind:value={name}
					class="w-full rounded-2xl border border-stone-300 bg-stone-50 px-4 py-3 text-base transition outline-none focus:border-stone-500 focus:bg-white"
					placeholder="Photosynthesis"
					type="text"
				/>
			</div>

			<div class="space-y-2">
				<label class="text-sm font-medium text-stone-700" for="card-content">Content</label>
				<textarea
					id="card-content"
					name="content"
					bind:value={content}
					class="min-h-36 w-full rounded-2xl border border-stone-300 bg-stone-50 px-4 py-3 text-base transition outline-none focus:border-stone-500 focus:bg-white"
					placeholder="Plants convert light energy into chemical energy."
				></textarea>
			</div>

			<Button class="rounded-full px-5" type="submit">Save draft</Button>
		</form>

		{#if submittedCard}
			<section class="rounded-3xl border border-stone-200 bg-white p-6 shadow-sm">
				<p class="text-sm font-semibold tracking-[0.24em] text-stone-500 uppercase">
					Submitted preview
				</p>
				<h2 class="mt-4 text-2xl font-semibold text-stone-950">
					{submittedCard.name || 'Untitled card'}
				</h2>
				<p class="mt-3 whitespace-pre-wrap text-stone-700">
					{submittedCard.content || 'No content added yet.'}
				</p>
			</section>
		{/if}
	</div>
</div>
