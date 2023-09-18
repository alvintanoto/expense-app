<script setup>
const router = useRouter();
const currentRoute = router.currentRoute;

const getLeftNavigatorColor = (value) => {
  if (currentRoute.value.name === value) {
    return "text-rp-dawn-text dark:text-rp-moon-text hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay";
  }

  return "text-rp-dawn-subtle/50 dark:text-rp-moon-subtle hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay";
};

const getProfileNavigationColor = () => {
  if (currentRoute.value.name !== "profile") {
    return "text-rp-dawn-gold/80 dark:text-rp-moon-gold/50 bg-rp-dawn-overlay dark:bg-rp-moon-overlay hover:bg-rp-dawn-subtle hover:dark:bg-rp-moon-subtle";
  }

  return "text-rp-dawn-gold dark:text-rp-moon-gold bg-rp-dawn-overlay dark:bg-rp-moon-overlay hover:bg-rp-dawn-subtle hover:dark:bg-rp-moon-subtle";
};
</script>

<template>
  <div
    class="flex flex-row min-h-screen bg-gray-200 bg-rp-dawn-overlay dark:bg-rp-moon-base text-rp-dawn-text dark:text-rp-moon-text"
  >
    <!-- body -->
    <div class="flex flex-row flex-1">
      <!-- left navigation -->
      <div
        class="w-[80px] bg-rp-dawn-surface dark:bg-rp-moon-surface flex flex-col drop-shadow-xl"
      >
        <div class="flex flex-col flex-1">
          <div class="h-[64px] text-center font-bold"></div>
          <NuxtLink
            to="/transaction"
            :class="getLeftNavigatorColor('transaction')"
            class="h-[80px] w-[80px] text-center text-xs flex flex-col"
          >
            <div class="m-auto">
              <IconTransaction class="text-center mx-auto" />
              Transaction
            </div>
          </NuxtLink>
          <NuxtLink
            to="/report"
            :class="getLeftNavigatorColor('report')"
            class="h-[80px] w-[80px] text-center text-xs flex flex-col"
          >
            <div class="m-auto">
              <IconReport class="text-center mx-auto" />
              Report
            </div>
          </NuxtLink>
          <NuxtLink
            to="/wallet"
            :class="getLeftNavigatorColor('wallet')"
            class="h-[80px] w-[80px] text-center text-xs flex flex-col"
          >
            <div class="m-auto">
              <IconWallet class="text-center mx-auto" />
              Wallet
            </div>
          </NuxtLink>
          <NuxtLink
            to="/settings"
            :class="getLeftNavigatorColor('settings')"
            class="h-[80px] w-[80px] text-center text-xs flex flex-col"
          >
            <div class="m-auto">
              <IconSetting class="text-center mx-auto" />
              Settings
            </div>
          </NuxtLink>
        </div>
        <NuxtLink
          to="/profile"
          :class="getProfileNavigationColor()"
          class="h-[80px] w-[80px] text-center text-xs flex flex-col"
        >
          <div class="m-auto">
            <IconProfile class="text-center mx-auto" />
            Profile
          </div>
        </NuxtLink>
      </div>

      <!-- container -->

      <div class="flex-1 overflow-auto flex flex-col">
        <div class="h-[64px] w-full bg-rp-dawn-surface dark:bg-rp-moon-surface">
          <slot name="header" />
        </div>
        <div class="flex container-xl mx-8 mt-8">
          <div
            class="w-max min-w-[786px] container-md mx-auto bg-rp-dawn-surface dark:bg-rp-moon-surface drop-shadow-xl"
          >
            <slot name="container" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
