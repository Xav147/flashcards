import { derived, writable } from 'svelte/store';

export type Deck = {
	name: string;
	size: number;
};

export const apiData = writable<Deck[]>([]);

export const deckNamesStore = derived(apiData, ($apiData) =>
	$apiData.map((deck): [string, number] => [deck.name, deck.size])
);
