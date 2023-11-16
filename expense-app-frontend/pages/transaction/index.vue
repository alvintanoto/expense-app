<script setup>
definePageMeta({
  middleware: ["authenticated"],
});

const layout = "client";
</script>

<template>
  <NuxtLayout :name="layout">
    <div class="text-xl p-4">Transaction</div>
  </NuxtLayout>
</template>
