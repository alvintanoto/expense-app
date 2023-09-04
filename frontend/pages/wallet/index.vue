<script setup>
const layout = "client";

const colors = [
  "bg-rp-dawn-love/10 dark:bg-rp-moon-love/10 text-rp-dawn-love dark:text-rp-moon-love",
  "bg-rp-dawn-gold/10 dark:bg-rp-moon-gold/10 text-rp-dawn-gold dark:text-rp-moon-gold",
  "bg-rp-dawn-rose/10 dark:bg-rp-moon-rose/10 text-rp-dawn-rose dark:text-rp-moon-rose",
  "bg-rp-dawn-pine/10 dark:bg-rp-moon-pine/10 text-rp-dawn-pine dark:text-rp-moon-pine",
  "bg-rp-dawn-foam/10 dark:bg-rp-moon-foam/10 text-rp-dawn-foam dark:text-rp-moon-foam",
  "bg-rp-dawn-iris/10 dark:bg-rp-moon-iris/10 text-rp-dawn-iris dark:text-rp-moon-iris",
];
</script>

<template>
  <NuxtLayout :name="layout">
    <div
      class="w-full rounded-lg px-4"
    >
      <div class="text-xl p-4 my-2">Wallet Management</div>
      <button class="mx-2 mb-4 px-4 py-2 bg-rp-dawn-pine/80 dark:bg-rp-moon-pine/80 hover:bg-rp-dawn-pine hover:dark:bg-rp-moon-pine rounded-md">Create Wallet</button>
      <!-- Loop this for list of wallet -->
      <div class="flex flex-row flex-wrap">
        <div
          v-for="(item, index) in colors"
          :key="index"
          class="p-4 m-2 w-[280px] rounded-md"
          :class=item
        >
          <div class="font-bold">Wallet Name</div>
          <div class="font-light">USD</div>
          <div class="text-sm text-right">2 September 2023</div>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>
