<script setup>
const router = useRouter();

const authStore = useAuthStore();
const currencyStore = useCurrencyStore();
const walletStore = useWalletStore();

const onLogoutClicked = async () => {
  // TODO: MOVE LOGOUT PROCESS TO COMPOSABLES
  currencyStore.$reset();
  authStore.$reset();
  walletStore.$reset();

  router.push({ path: "/login" });
};
</script>

<template>
  <div class="flex h-[64px] justify-end items-center mx-4">
    <button
      @click="onLogoutClicked"
      class="bg-rp-dawn-love/30 hover:bg-rp-dawn-love/40 dark:bg-rp-moon-love/50 hover:dark:bg-rp-moon-love/75 my-auto text-white font-bold py-2 px-4 rounded"
    >
      Logout
    </button>
  </div>
</template>
