<script setup>
definePageMeta({
  middleware: ["authenticated"],
});
</script>

<template>
  <main
    class="min-h-screen p-4 bg-rp-dawn-base text-rp-dawn-text dark:bg-rp-moon-base dark:text-rp-moon-text"
  >
    <div class="container mx-auto p-4">
      <div class="flex flex-row items-center justify-between">
        <div class="text-xl">Expense App</div>
        <div class="flex flex-row items-center">
          <NuxtLink
            to="/login"
            class="px-4 py-1 cursor-pointer mr-1 hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay rounded-md"
          >
            Login
          </NuxtLink>
          <NuxtLink
            to="/register"
            class="px-4 py-1 cursor-pointer hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay rounded-md"
          >
            Register
          </NuxtLink>
        </div>
      </div>
    </div>
  </main>
</template>
