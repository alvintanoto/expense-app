<script setup></script>

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

      <div class="mt-16 container-md mx-auto">
        <div class="text-lg">Login</div>
        <div class="flex flex-col mt-4">
          <label for="username"> Username </label>
          <input
            type="text"
            name="username"
            id="username"
            placeholder="Username"
            class="mt-2 appearance-none border-2 text-white border-rp-dawn-surface dark:border-rp-moon-surface rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
          />
        </div>
        <div class="flex flex-col mt-4">
          <label for="password"> Password </label>
          <input
            type="password"
            name="password"
            id="password"
            placeholder="Password"
            class="mt-2 appearance-none border-2 text-white border-rp-dawn-surface dark:border-rp-moon-surface rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
          />
        </div>
        <div class="float-right mt-8">
            <button class="px-8 py-2 rounded-md 
            bg-rp-dawn-surface dark:bg-rp-moon-surface
            hover:bg-rp-dawn-overlay dark:hover:bg-rp-moon-overlay"> Login </button>
        </div>
      </div>
    </div>
  </main>
</template>
