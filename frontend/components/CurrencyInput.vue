<script setup>
import { toRefs } from "vue";
import { useCurrencyInput } from "vue-currency-input";

const props = defineProps(["modelValue", "options"]);
const data = toRefs(props);

const {  inputRef,
      formattedValue,
      setOptions,
      setValue } = useCurrencyInput(data.options.value);

watch(
  () => data.options.value,
  (options) => {
    setOptions(options);
  }
);
</script>

<template>
  <input
    ref="inputRef"
    :value="formattedValue"
    type="text"
    placeholder="Type your starting balance"
    class="w-full outline-none bg-rp-dawn-surface dark:bg-rp-moon-surface"
  />
</template>
