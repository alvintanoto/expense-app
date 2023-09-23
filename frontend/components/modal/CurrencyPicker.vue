<script setup>
import { ref } from "vue";

const emit = defineEmits(["onCurrencySelected"]);
const currencies = ref([]);

const fetchCurrencyData = async () => {
  const { data } = await useFetch("/api/currency");

  if (!data.value) {
    // TODO: handle cannot connect to server error
    console.log("handle cannot connect to server error");
    return;
  }

  if (data.value.code !== "20000") {
    // TODO: handle bad request error
    console.log("handle bad request error");
    return;
  }

  currencies.value = data.value.data;
};

const onCurrencySelected = (currency) => {
  emit("onCurrencySelected", currency);
};

fetchCurrencyData();
// TODO: search currency
</script>

<template>
  <div
    @click.stop=""
    class="absolute top-1/4 left-1/2 transform -translate-x-1/2 bg-rp-dawn-surface dark:bg-rp-moon-surface rounded-md max-w-[640px]"
  >
    <div class="border-b-[1px] font-bold p-4">Select Currency</div>
    <div class="m-2 p-2">
      <input
        type="text"
        placeholder="Search currency"
        class="w-full outline-none bg-rp-dawn-surface dark:bg-rp-moon-surface"
      />
    </div>
    <div class="max-h-[480px] overflow-y-auto grid grid-cols-3">
      <div
        v-for="(currency, index) in currencies"
        :key="currency.currency_id"
        @click="onCurrencySelected(currency)"
        class="m-2 p-2 rounded-md cursor-pointer border-[1px] border-rp-dawn-text/20 dark:border-rp-moon-text/50 hover:border-rp-dawn-text dark:hover:border-rp-moon-text"
      >
        <div>{{ currency.currency_code }}</div>
        <div class="text-sm opacity-50">{{ currency.currency_name }}</div>
      </div>
    </div>
  </div>
</template>
