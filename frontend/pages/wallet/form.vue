<script setup>
const layout = "client";
</script>

<template>
  <NuxtLayout :name="layout">
    <div class="w-full rounded-lg px-4">
      <div class="text-xl p-4 my-2">[Action] Wallet</div>

      <form @submit.prevent="">
        <div class="flex flex-col mt-4 p-4">
          <label for="name"> Wallet Name </label>
          <input
            type="text"
            name="name"
            id="name"
            placeholder="Wallet Name"
            class="mt-2 appearance-none border-2 text-white border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
          />
        </div>

        <div class="flex flex-col px-4 py-2">
          <label for="name"> Wallet Currency </label>
          <select
            type="text"
            name="currency"
            id="currency"
            placeholder="Wallet Currency"
            class="mt-2 appearance-none border-2 text-white border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-rp-dawn-overlay dark:focus:focus:bg-rp-moon-overlay focus:border-rp-dawn-iris dark:focus:border-rp-moon-iris bg-rp-dawn-surface dark:bg-rp-moon-surface"
          >
            <option>USD</option>
            <option>EURO</option>
            <option>IDR</option>
            <option>INR</option>
            <option>BSDD</option>
            <option>JKT</option>
            <option>HSD</option>
          </select>
        </div>

        <div class="float-right mt-4 p-4">
          <button
            @click="submit"
            class="px-8 py-2 border-2 border-rp-dawn-highlight-high dark:border-rp-moon-highlight-high rounded-md bg-rp-dawn-surface dark:bg-rp-moon-surface hover:bg-rp-dawn-overlay dark:hover:bg-rp-moon-overlay"
          >
            [Action]
          </button>
        </div>
      </form>
    </div>
  </NuxtLayout>
</template>
