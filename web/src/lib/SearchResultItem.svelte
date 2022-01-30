<script>

  export let type = 'stock'
  export let symbol = 'N/A'
  export let name = 'N/A'

  $: iconPath = type === 'stock' ? 'images/stock-icon.svg' : 'images/crypto-icon.svg'

</script>

<div class="px-4 py-5 bg-gray-800 shadow rounded-full overflow-hidden sm:p-6 bg-no-repeat bg-auto" style="background-image: url({iconPath})">
  <dt class="mt-1 text-3xl font-semibold text-gray-300">
    {symbol}
  </dt>
  <dd class="text-sm font-medium text-gray-500 truncate">
    {name}
  </dd>
</div>