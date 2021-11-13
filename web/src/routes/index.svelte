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
			download="symbol-list-{fileDate}.json"
			class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md transition-all duration-500 text-gray-200 bg-white hover:bg-gray-50 sm:w-auto bg-gradient-to-br hover:from-green-400 hover:to-blue-500 from-pink-600 via-pink-500 to-yellow-500 "
		>
			Download JSON dataset ({fileDate})
		</a>
		<a
			href="/data/daily-symbol-list.csv"
			download="symbol-list-{fileDate}.csv"
			class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md transition-all duration-500 text-gray-200 bg-white hover:bg-gray-50 sm:w-auto bg-gradient-to-br hover:from-green-400 hover:to-blue-500 from-pink-600 via-pink-500 to-yellow-500 "
		>
			Download CSV dataset ({fileDate})
		</a>
	</div>
</div>

<div class="bg-gray-900">
  <div class="max-w-7xl mx-auto py-8 px-4 sm:py-10 sm:px-6 lg:px-8">
    <div class="lg:max-w-2xl lg:mx-auto lg:text-center">
      <h2 class="text-3xl font-extrabold tracking-tight text-white sm:text-4xl">
        Frequently asked questions
      </h2>
      <p class="mt-4 text-gray-400">Hopefully these answer your questions, feel free to reach out on twitter <a class="font-semibold text-gray-300 hover:text-gray-200" href="https://twitter.com/jonfriesen" target="_blank">@jonfriesen</a></p>
			
    </div>
    <div class="mt-20">
      <dl class="space-y-10 lg:space-y-0 lg:grid lg:grid-cols-2 lg:gap-x-8 lg:gap-y-10">
				<div>
          <dt class="font-semibold text-white">
            Why was this project created?
          </dt>
          <dd class="mt-3 text-gray-400">
            Finding reliable and accurate bulk symbols lists that didn't break the bank was a huge challenge. I've written a few small tools to help me manage my portfolio and analyze the market. One of the public projects is <a class="font-semibold text-gray-300 hover:text-gray-200" href="https://buythe.top" target="_blank">🤖 buy the top bot</a> which uses these symbols to return stock data in Telegram channels.
          </dd>
        </div>
        <div>
          <dt class="font-semibold text-white">
            Where does this data come from?
          </dt>
          <dd class="mt-3 text-gray-400">
            In it's current iteration this data is pulled directly from the nasdaq and tsx exchanges. Fortunately, the nasdaq includes a full list of other listed symbols which includes the NYSE. Here are the sources:
						<ul class="list-inside list-disc">
							<li><a class="font-semibold text-gray-300 hover:text-gray-200" href="ftp://ftp.nasdaqtrader.com/symboldirectory/nasdaqlisted.txt" target="_blank">nasdaq symbols (ftp)</a></li>
							<li><a class="font-semibold text-gray-300 hover:text-gray-200" href="ftp://ftp.nasdaqtrader.com/symboldirectory/otherlisted.txt" target="_blank">nyse (and other) symbols (ftp)</a></li>
							<li><a class="font-semibold text-gray-300 hover:text-gray-200" href="https://www.tsx.com/json/company-directory/search/tsx/^*" target="_blank">tsx</a></li>
							<li><a class="font-semibold text-gray-300 hover:text-gray-200" href="https://www.tsx.com/json/company-directory/search/tsxv/^*" target="_blank">tsxv</a></li>
						</ul>
          </dd>
        </div>
				<div>
          <dt class="font-semibold text-white">
            How does it work?
          </dt>
          <dd class="mt-3 text-gray-400">
            The data is pulled from the exchanges and compiled into a single JSON & CSV file. The data is then stored in a public GitHub repository. The data is updated daily using a scheduled GitHub Action workflow. Once the action to get the latest symbol list is run, the new data is commited to the same repository. One new commits a DigitalOcean App Platform triggers and builds the static site with the latest dataset.
          </dd>
        </div>
				<div>
          <dt class="font-semibold text-white">
            Who are you?
          </dt>
          <dd class="mt-3 text-gray-400">
            👋&nbsp; I'm <a class="font-semibold text-gray-300 hover:text-gray-200" href="https://jonfriesen.ca" target="_blank">Jon!</a> I'm a software engineer who mainly works on cloud platform as a service products.
          </dd>
        </div>
      </dl>
    </div>
  </div>
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