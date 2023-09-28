<script setup>
definePageMeta({
  middleware: ["authenticated"],
});
const layout = "client";
</script>

<template>
  <NuxtLayout :name="layout">
    <template #header>
      <div class="grid grid-cols-2">
        <div></div>
        <div class="flex h-[64px] justify-end items-center">
          <ButtonLogout />
        </div>
      </div>
    </template>
  </NuxtLayout>
</template>
