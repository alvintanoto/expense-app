<script setup>
const router = useRouter();

const onLogoutClicked = async () => {
    // TODO: LOGOUT PROCESS
    router.push({path: "/login"})
};
</script>

<template>
  <div class="flex h-[64px] justify-end items-center mx-4">
    <button
      @click="onLogoutClicked"
      class="bg-rp-dawn-love/30 hover:bg-rp-dawn-love/40 dark:bg-rp-moon-love/25 hover:dark:bg-rp-moon-love/50 my-auto text-white font-bold py-2 px-4 rounded"
    >
      Logout
    </button>
  </div>
</template>
