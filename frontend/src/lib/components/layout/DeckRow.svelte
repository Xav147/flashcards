<script lang="ts">
	let {
		deckName,
		deckSize
	}: {
		deckName: string;
		deckSize: number;
	} = $props();
</script>

<div class="grid grid-cols-3 items-center gap-4 bg-gray-400 px-4 py-2">
	<span class="min-w-0 truncate">{deckName}</span>
	<span aria-hidden="true"></span>
	<span class="text-right">{deckSize}</span>
</div>
