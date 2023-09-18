<script setup>
import { ref } from "vue";
const errorMessage = ref("");

const handleLogin = (evt) => {
  const username = event.target.username.value.trim();
  const password = event.target.password.value.trim();

  if (username === "") {
    errorMessage.value = "Please input your username.";
    return;
  }

  if (password === "") {
    errorMessage.value = "Please input your password.";
    return;
  }

  navigateTo("/transaction");
};
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
            to="/register"
            class="px-4 py-1 cursor-pointer hover:bg-rp-dawn-overlay hover:dark:bg-rp-moon-overlay rounded-md"
          >
            Register
          </NuxtLink>
        </div>
      </div>

      <div class="mt-16 container-md max-w-[768px] mx-auto">
        <form @submit.prevent="handleLogin($event)">
          <div class="text-lg font-bold">Login</div>
          <div
            v-if="errorMessage"
            class="my-4 p-4 border-2 rounded-md bg-rp-dawn-love/10 dark:bg-rp-dawn-love/10 border-rp-dawn-love dark:border-rp-moon-love text-rp-dawn-love dark:text-rp-moon-love"
          >
            This is error message
          </div>
          <div class="flex flex-col mt-4">
            <label for="username"> Username </label>
            <input
              type="text"
              name="username"
              id="username"
              placeholder="Username"
              class="mt-2 appearance-none border-2 text-white border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
            />
          </div>
          <div class="flex flex-col mt-4">
            <label for="password"> Password </label>
            <input
              type="password"
              name="password"
              id="password"
              placeholder="Password"
              class="mt-2 appearance-none border-2 text-white border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
            />
          </div>
          <div class="float-right mt-8">
            <button
              @click="submit"
              class="px-8 py-2 border-2 border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded-md bg-rp-dawn-surface dark:bg-rp-moon-surface hover:bg-rp-dawn-overlay dark:hover:bg-rp-moon-overlay"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  </main>
</template>
