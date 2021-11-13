<script>
	const today = new Date();
	let fileDate = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getUTCDate();

	// Use yesterdays archive if it's too early for todays
	if (today.getUTCHours() < 2) {
		const yesterday = new Date(today);
		yesterday.setDate(yesterday.getUTCDate() - 1);
		fileDate =
			yesterday.getFullYear() + '-' + (yesterday.getMonth() + 1) + '-' + yesterday.getDate();
	}
</script>

<div>
<header class="bg-gray-900">
	<nav class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8" aria-label="Top">
		<div class="w-full py-6 flex items-center justify-end border-none">
			<a
				href="https://github.com/jonfriesen/symbol-list/tree/main/data"
				target="_blank"
				class="inline-flex  py-2 px-4 items-center justify-center border border-transparent rounded-md text-base font-medium transition duration-500 text-gray-200 hover:bg-gray-400 "
			>
				Archive
			</a>
			<div class="space-x-4">
				<a
					href="https://github.com/jonfriesen/symbol-list"
					target="_blank"
					class="inline-block py-2 px-4 border border-transparent rounded-md text-base font-medium transition duration-500 text-gray-600 hover:bg-gray-400"
				>
					<img
						class="h-8"
						src="/images/octocat.png"
						alt="GitHub Octocat link to project repository"
					/>
				</a>
			</div>
		</div>
	</nav>
</header>

<div class="bg-gray-900">
	<div class="max-w-2xl mx-auto text-center py-16 px-4 sm:py-20 sm:px-6 lg:px-8">
		<h2 class="text-3xl font-extrabold text-white sm:text-4xl">
			<span class="block">Symbol List</span>
			<span class="block">Daily market symbol lists</span>
		</h2>
		<p class="mt-4 text-lg leading-6 text-gray-200">
			NASDAQ, NYSE, TSX, & TSXV symbol lists pulled and compiled every night directly from the
			exchanges.
		</p>
		<a
			href="/data/daily-symbol-list.json"
			download="symbol-list.json"
			class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md transition-all duration-500 text-gray-200 bg-white hover:bg-gray-50 sm:w-auto bg-gradient-to-br hover:from-green-400 hover:to-blue-500 from-pink-600 via-pink-500 to-yellow-500 "
		>
			Download JSON dataset ({fileDate})
		</a>
	</div>
</div>

<div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
	<!-- todo content -->
</div>

<footer class="bg-gray-800 h-screen">
	<div
		class="max-w-7xl mx-auto py-12 px-4 sm:px-6 md:flex md:items-center md:justify-between lg:px-8 text-sm"
	>
		<div class="flex order-1 text-gray-200">
			Sponsored by
			<a href="https://planetside.co" target="_blank" class="px-4">
				<img class="h-4" src="/images/logo.svg" alt="Planetside Software logo" />
			</a>
		</div>
		<div class="flex mt-8 md:mt-0  text-gray-200">&copy; 2021 Jon Friesen All rights reserved.</div>
	</div>
</footer>
</div>