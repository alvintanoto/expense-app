<script setup>
if (process.client) {
  const classList = document.documentElement.classList;
  if (
    localStorage.theme === "dark" ||
    (!("theme" in localStorage) &&
      window.matchMedia("(prefers-color-scheme: dark)").matches)
  ) {
    classList.add("dark");
  } else {
    classList.remove("dark");
  }
}
</script>

<template>
  <div>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>
