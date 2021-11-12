<script>

  const today = new Date();
  let fileDate = today.getFullYear() + '-' + (today.getMonth() + 1) + '-' + today.getDate();

  // Use yesterdays archive if it's too early for todays
  if (today.getUTCHours() < 3) {
    const yesterday = new Date(today)
    yesterday.setDate(yesterday.getDate() - 1)
    fileDate = yesterday.getFullYear() + '-' + (yesterday.getMonth() + 1) + '-' + yesterday.getDate();
  }

  

</script>

<header class="bg-indigo-700">
  <nav class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8" aria-label="Top">
    <div class="w-full py-6 flex items-center justify-between border-b border-indigo-500 lg:border-none">
      <div class="flex items-center">
      </div>
      <div class="ml-10 space-x-4">
        <a href="https://github.com/jonfriesen/symbol-list" target="_blank" class="inline-block py-2 px-4 border border-transparent rounded-md text-base font-medium text-indigo-600 hover:bg-indigo-400">
          <img class="h-8" src="/images/octocat.png" alt="GitHub Octocat link to project repository" />
        </a>
      </div>
    </div>
  </nav>
</header>


<div class="bg-indigo-700">
  <div class="max-w-2xl mx-auto text-center py-16 px-4 sm:py-20 sm:px-6 lg:px-8">
    <h2 class="text-3xl font-extrabold text-white sm:text-4xl">
      <span class="block">Symbol List</span>
      <span class="block">Daily market symbol lists</span>
    </h2>
    <p class="mt-4 text-lg leading-6 text-indigo-200">NASDAQ, NYSE, TSX, & TSXV symbol lists pulled and compiled every night directly from the exchanges.</p>
    <a href="/data/{fileDate}.json" download="symbol-list.json" class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-indigo-600 bg-white hover:bg-indigo-50 sm:w-auto">
      Download JSON dataset
    </a>
    <a href="https://github.com/jonfriesen/symbol-list/tree/main/data" target="_blank" class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-indigo-600 bg-white hover:bg-indigo-50 sm:w-auto">
      Archive
    </a>
  </div>
</div>

<div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
  <!-- todo content -->
</div>

<footer class="bg-white">
  <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 md:flex md:items-center md:justify-between lg:px-8">
    <div class="flex mt-8 md:mt-0 md:order-1 text-gray-400">
      
      &copy; 2021
      <a href="https://planetside.co" target="_blank" class="px-4">
        <img class="h-4 mt-1" src="/images/logo.svg" alt="Planetside Software logo" />
      </a>
      All rights reserved.
      
  </div>
</footer>
