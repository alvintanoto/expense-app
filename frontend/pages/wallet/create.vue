<script setup>
import { ref } from "vue";
const layout = "client";

const isCurrencyPickerShown = ref(false);
const selectedCurrency = ref({ currency: 'USD' })

const showCurrencyPicker = () => {
  isCurrencyPickerShown.value = true;
};

const closeAllModal = () => {
  isCurrencyPickerShown.value = false;
};

const onCurrencySelected = (currency) => {
    closeAllModal()

    selectedCurrency.value = {
        currency: currency.currency_code
    }
}

</script>

<template>
  <NuxtLayout :name="layout">
    <template #modal>
      <div
        v-if="isCurrencyPickerShown"
        @click.stop="closeAllModal()"
        class="absolute z-50 w-full h-full bg-rp-dawn-highlight-high/50 dark:bg-rp-dawn-highlight-high/50"
      >
        <ModalCurrencyPicker @onCurrencySelected="onCurrencySelected" /></div
    ></template>
    <template #header> </template>
    <template #container>
      <div class="p-4 font-bold border-b-[1px]">Create a new wallet!</div>
      <div class="p-4 mt-4">
        <div
          class="border-[1px] rounded-md border-rp-dawn-text/20 dark:border-rp-moon-text/50 hover:border-rp-dawn-text dark:hover:border-rp-moon-text"
        >
          <div
            class="text-[12px] mx-2 mt-2 text-rp-dawn-text/50 dark:text-rp-moon-text/75"
          >
            Wallet Name
          </div>
          <div class="mx-2 py-2">
            <input
              type="text"
              placeholder="Type your wallet name"
              class="w-full outline-none bg-rp-dawn-surface dark:bg-rp-moon-surface"
            />
          </div>
        </div>

        <div class="flex flex-row mt-2">
          <div
            class="border-[1px] rounded-md border-rp-dawn-text/20 dark:border-rp-moon-text/50 hover:border-rp-dawn-text dark:hover:border-rp-moon-text"
          >
            <div
              class="cursor-pointer text-[12px] mx-2 mt-2 text-rp-dawn-text/50 dark:text-rp-moon-text/75"
            >
              Currency
            </div>
            <div
              class="cursor-pointer flex flex-row mx-2 py-2"
              @click="showCurrencyPicker"
            >
              <div class="cursor-pointer"><IconCurrency /></div>
              <div
                class="cursor-pointer mx-2 text-rp-dawn-text dark:text-rp-moon-text"
              >
                {{selectedCurrency.currency}}
              </div>
              <div class="cursor-pointer max-h-[24px]">
                <IconChevronRight />
              </div>
            </div>
          </div>

          <div
            class="border-[1px] ml-2 flex-1 rounded-md border-rp-dawn-text/20 dark:border-rp-moon-text/50 hover:border-rp-dawn-text dark:hover:border-rp-moon-text"
          >
            <div
              class="text-[12px] mx-2 mt-2 text-rp-dawn-text/50 dark:text-rp-moon-text/75"
            >
              Starting Balance
            </div>
            <div class="mx-2 py-2">
              <CurrencyInput modelValue="0" :options="selectedCurrency" />
            </div>
          </div>
        </div>

        <div class="flex flex-row mt-4 justify-end">
          <button
            class="text-sm rounded-md px-8 py-2 text-rp-dawn-surface bg-rp-dawn-text dark:text-rp-moon-surface dark:bg-rp-moon-text"
          >
            Save
          </button>
        </div>
      </div>
    </template>
  </NuxtLayout>
</template>
