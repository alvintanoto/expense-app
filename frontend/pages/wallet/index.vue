<script setup>
definePageMeta({
  middleware: ["authenticated"],
});

const layout = "client";
</script>

<template>
  <NuxtLayout :name="layout">
    <template #header> </template>
    <template #container> </template>
  </NuxtLayout>
</template>
