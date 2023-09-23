// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    baseUrl: process.env.BASE_URL,
    currencyEndpoint: process.env.CURRENCY_ENDPOINT,
  },
  css: ["~/assets/base.css"],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  devtools: { enabled: true },
  modules: [["@pinia/nuxt", {
    autoImports: ['defineStore', 'definePiniaStore']
  }]],
  imports: {
    dirs: ["stores"]
  }
});
