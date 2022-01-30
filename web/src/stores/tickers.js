import { writable, derived } from 'svelte/store';
import axios from 'axios';
import Fuse from 'fuse.js';

const cryptoPath = '/data/daily-crypto-symbol-list.json';
const stockPath = '/data/daily-symbol-list.json';

const getSetstore = async (path, setFn) => {
	try {
		const resp = await axios.get(path);
		setFn(resp.data);
	} catch (err) {
		console.log('axios error');
	}
};

const symbolStore = (path) => {
	const init = {};
	const { subscribe, set, update } = writable(init);

	getSetstore(path, set);

	return {
		subscribe,
		set,
		update
	};
};

export const stocks = symbolStore(stockPath);
export const crypto = symbolStore(cryptoPath);

export const symbols = derived([stocks, crypto], ([$stocks, $crypto]) => {
	let agg = [];
	if (!!$stocks.data && Array.isArray($stocks.data)) {
		$stocks.data.forEach((s) => {
			agg.push({
				type: 'stock',
				data: s
			});
		});
	}
	if (!!$crypto.data && Array.isArray($crypto.data)) {
		$crypto.data.forEach((s) => {
			agg.push({
				type: 'crypto',
				data: s
			});
		});
	}

	return agg;
});

export const count = derived(symbols, ($symbols) => {
	return $symbols.length;
});

const searchFn = (v, symbols, limit = 5) => {
	const options = {
		includeScore: true,
		threshold: 0.25,
		keys: [
			{
				name: 'data.symbol',
				weight: 0.7
			},
			{
				name: 'data.name',
				weight: 0.3
			}
		]
	};

	// Create a new instance of Fuse
	const fuse = new Fuse(symbols, options);

	// Now search for 'Man'
	const result = fuse.search(v);
	return result.slice(0, limit);
};

export const search = (() => {
	const _query = writable('');
	const _limit = writable(5);
	const _results = derived([symbols, _query, _limit], ([$symbols, $query, $limit]) =>
		searchFn($query, $symbols, $limit)
	);

	return {
		query: _query.set,
		limit: _limit.set,
		..._results
	};
})();
