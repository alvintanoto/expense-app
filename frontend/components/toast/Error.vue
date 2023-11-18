<script setup>
import { ref } from "vue";

const router = useRouter();
const currentRoute = router.currentRoute;

const globalStore = useGlobalStore();
const errMessageRef = ref("");

if (globalStore.errorMessage != "") {
  if (globalStore.forceLogout && currentRoute.value.name != "login") {
    router.push({ path: "/login" });
    globalStore.forceLogout = false;
  }

  errMessageRef.value = globalStore.errorMessage;
  setInterval(() => {
    globalStore.forceLogout = false;
    globalStore.errorMessage = "";
    errMessageRef.value = "";
  }, 3000);
}
</script>

<template>
  <div
    v-if="errMessageRef"
    id="toast-top-right"
    class="bg-rp-moon-love text-[color:white] dark:text-white dark:bg-rp-moon-love text-white fixed flex items-center w-full max-w-xs p-4 space-x-4 divide-x rtl:divide-x-reverse divide-gray-200 rounded-lg shadow top-5 right-5"
    role="alert"
  >
    <div class="text-sm font-normal">{{ errMessageRef }}</div>
  </div>
</template>
