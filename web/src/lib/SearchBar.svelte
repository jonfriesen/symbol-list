<script>
	import { search } from '../stores/tickers';

	let timer;
	let v = '';
	const keyPress = () => {
		clearTimeout(timer);
		timer = setTimeout(() => {
			search.query(v);
		}, 10);
	};
</script>

<div>
	<div class="mt-1">
		<input
			type="text"
			name="symbol"
			id="symbol"
			bind:value={v}
			on:input={keyPress}
			class="shadow-sm w-full hover:ring-2 focus:ring-purple-500 hover:ring-pink-500 block text-gray-200 md:text-xl sm:text-sm bg-gray-800 border-gray-800 px-4  py-4 rounded-full"
			placeholder="Start typing to search for a ticker..."
		/>
	</div>
</div>
