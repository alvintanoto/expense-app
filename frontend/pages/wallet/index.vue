<script setup>
const layout = "client";

const wallets = [
  {
    name: "Kantong USD",
    currency: "USD",
    createdAt: "2 September 2023",
  },
];
</script>

<template>
  <NuxtLayout :name="layout">
    <div class="w-full rounded-lg px-4">
      <div class="text-xl m-2 my-4">Wallet Management</div>
      <nuxt-link :to="{ path: '/wallet/form', query: { action: 'create' } }">
        <button
          class="mx-2 mb-4 px-4 py-2 bg-rp-dawn-iris/20 dark:bg-rp-moon-pine/80 hover:bg-rp-dawn-iris/40 hover:dark:bg-rp-moon-pine rounded-md"
        >
          Create Wallet
        </button>
      </nuxt-link>
      <!-- Loop this for list of wallet -->
      <div class="flex flex-row flex-wrap">
        <div
          v-for="(item, index) in wallets"
          :key="index"
          class="p-4 m-2 w-[320px] rounded-md cursor-pointer bg-rp-dawn-subtle/10 dark:bg-rp-moon-subtle/10 text-rp-dawn-subtle dark:text-rp-moon-subtle"
          :class="item"
        >
          <nuxt-link
            :to="{
              path: '/wallet/form',
              query: { action: 'update', wallet_id: index },
            }"
          >
            <div class="font-bold min-h-[24px]">{{ item.name }}</div>
            <div class="font-light min-h-[24px] mb-8">{{ item.currency }}</div>
            <div class="text-sm text-right min-h-[24px] italic">
              {{ item.createdAt }}
            </div>
          </nuxt-link>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>
