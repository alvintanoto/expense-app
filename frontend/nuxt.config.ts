// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    baseUrl: process.env.BASE_URL,
    loginEndpoint: process.env.LOGIN_ENDPOINT,
    logoutEndpoint: process.env.LOGOUT_ENDPOINT,
    registerEndpoint: process.env.REGISTER_ENDPOINT,
    checkTokenEndpoint: process.env.CHECK_TOKEN_ENDPOINT,
    refreshTokenEndpoint: process.env.REFRESH_TOKEN_ENDPOINT,
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
