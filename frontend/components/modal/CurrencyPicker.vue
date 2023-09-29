<script setup>
import { ref } from "vue";

const emit = defineEmits(["onCurrencySelected"]);
const searchCurrencies = ref("");

const currencyStore = useCurrenciesStore();
const { data } = await useAsyncData('currency', () => currencyStore.getCurrencies())

const onCurrencySelected = (currency) => {
  emit("onCurrencySelected", currency);
};

const onSearchChanged = (event) => {
  currencyStore.setCurrencyFilter(event.target.value)
};

</script>

<template>
  <div
    @click.stop=""
    class="absolute top-[128px] left-1/2 transform -translate-x-1/2 bg-rp-dawn-surface dark:bg-rp-moon-surface rounded-md max-w-[640px] min-w-[640px]"
  >
    <div class="border-b-[1px] font-bold p-4">Select Currency</div>
    <div class="p-2 border-b-[1px]">
      <input
        type="text"
        @input="onSearchChanged($event)"
        placeholder="Search currency"
        class="w-full p-2 outline-none bg-rp-dawn-surface dark:bg-rp-moon-surface"
      />
    </div>
    <div
      v-if="currencyStore.displayedCurrencyList.length > 0"
      class="max-h-[480px] mt-2 overflow-y-auto grid grid-cols-3"
    >
      <div
        v-for="(currency, index) in currencyStore.displayedCurrencyList"
        :key="currency.currency_id"
        @click="onCurrencySelected(currency)"
        class="m-2 p-2 rounded-md cursor-pointer border-[1px] border-rp-dawn-text/20 dark:border-rp-moon-text/50 hover:border-rp-dawn-text dark:hover:border-rp-moon-text"
      >
        <div>{{ currency.currency_code }}</div>
        <div class="text-sm opacity-50">{{ currency.currency_name }}</div>
      </div>
    </div>
    <div
      v-else
      class="max-h-[480px flex p-4 my-8 justify-center"
    >
      Couldn't find currency 
    </div>
  </div>
</template>
