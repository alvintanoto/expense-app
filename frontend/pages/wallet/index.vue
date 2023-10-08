<script setup>
definePageMeta({
  middleware: ["authenticated"],
});
const walletStore = useWalletStore();
const layout = "client";

onMounted(async () => {
  await walletStore.getUserWallets()
})
</script>

<template>
  <NuxtLayout :name="layout">
    <template #header> </template>
    <template #container>
      <div
        class="w-max min-w-[480px] rounded-md container-md mx-auto my-4 bg-rp-dawn-surface dark:bg-rp-moon-surface drop-shadow-xl"
      >
        <div class="p-4 font-bold border-b-[1px]">Your wallets</div>
        <div class="p-4">
          <div
            v-for="(item, index) in walletStore.walletList"
            :key="index"
            class="m-2 p-4 flex flex-row cursor-pointer items-center hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay rounded-md"
          >
            <div class="mr-4"><IconWallet /></div>
            <div class="flex flex-col">
              <div class="font-bold">
                {{ item.wallet.wallet_name }}
              </div>
              <div class="text-sm">
                {{ item.currency.currency_code }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </NuxtLayout>
</template>
