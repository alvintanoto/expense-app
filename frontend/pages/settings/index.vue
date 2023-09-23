<script setup>
definePageMeta({
  middleware: ["authenticated"],
});
const layout = "client";
</script>

<template>
  <NuxtLayout :name="layout">
  </NuxtLayout>
</template>
