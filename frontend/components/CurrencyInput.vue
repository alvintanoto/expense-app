<script setup>
import {defineProps} from "vue"
import { useCurrencyInput } from "vue-currency-input";

const props = defineProps(["modelValue", "options"])
const {data} = toRefs(props);


const { inputRef } = useCurrencyInput(props.options)
</script>

<template>
  <input ref="inputRef" type="text" placeholder="Type your starting balance" class="w-full outline-none bg-rp-dawn-surface dark:bg-rp-moon-surface"/>
</template>
