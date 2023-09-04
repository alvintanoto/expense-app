<script setup>
const router = useRouter();
const currentRoute = router.currentRoute;

const getLeftNavigatorColor = (value) => {
  if (currentRoute.value.name === value) {
    return "text-rp-dawn-text dark:text-rp-moon-text";
  }

  return "text-rp-dawn-subtle dark:text-rp-moon-subtle";
};
</script>

<template>
  <div
    class="flex flex-row min-h-screen bg-gray-200 text-white bg-rp-dawn-base dark:bg-rp-moon-base text-rp-dawn-text dark:text-rp-moon-text"
  >

    <!-- body -->
    <div class="flex flex-row flex-1">

      <!-- left navigation -->
      <div
        class="w-[240px] bg-rp-dawn-surface dark:bg-rp-moon-surface flex flex-col"
      >
        <div class="flex flex-col flex-1">
          <div class="p-2 h-[56px] text-center font-bold my-4">Expense App</div>
          <div class="p-2">
            <NuxtLink
              to="/dashboard"
              class="p-2 cursor-pointer"
              :class="getLeftNavigatorColor('dashboard')"
            >
              Dashboard
            </NuxtLink>
          </div>
          <div class="p-2">
            <NuxtLink
              to="/wallet"
              class="p-2 cursor-pointer"
              :class="getLeftNavigatorColor('wallet')"
            >
              Wallet
            </NuxtLink>
          </div>
          <div class="p-2">
            <NuxtLink
              to="/transaction"
              class="p-2 cursor-pointer"
              :class="getLeftNavigatorColor('transaction')"
            >
              Transaction
            </NuxtLink>
          </div>
        </div>
        <NuxtLink
        to="/profile"
        class="p-6 text-center bg-rp-dawn-overlay dark:bg-rp-moon-overlay">
           [Username]
        </NuxtLink>
      </div>

      <!-- container -->
      <div class="bg-slate-800 flex-1 overflow-auto flex flex-col">
        <div class="flex-1 p-4"><slot /></div>
        <div class="h-[24px] mx-auto m-4 p-8">&copy;2023 - AT</div>
      </div>
    </div>
  </div>
</template>
