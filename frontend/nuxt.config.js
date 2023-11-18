// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    baseUrl: process.env.BASE_URL,
    
    public: {
      loginEndpoint: process.env.LOGIN_ENDPOINT,
      logoutEndpoint: process.env.LOGOUT_ENDPOINT,
      registerEndpoint: process.env.REGISTER_ENDPOINT,
      checkTokenEndpoint: process.env.CHECK_TOKEN_ENDPOINT,
      refreshTokenEndpoint: process.env.REFRESH_TOKEN_ENDPOINT,
      currencyEndpoint: process.env.CURRENCY_ENDPOINT,
      createWalletEndpoint: process.env.CREATE_WALLET_ENDPOINT,
      getWalletEndpoint: process.env.GET_WALLET_ENDPOINT,
    }
  },
  devtools: { enabled: true },
  modules: [
    "@nuxtjs/tailwindcss", 
    "@pinia/nuxt",
    "@pinia-plugin-persistedstate/nuxt"
  ],
  pinia: {
    storesDirs: ['./stores/**'],
  },
  imports: {
    dirs: [
      'composables',
      'composables/**',
    ]
  },
  nitro: {
    routeRules: {
      "/api/**": {
        proxy: process.env.BASE_URL + "/api/**",
        changeOrigin:true,
      }
    }
  }
})