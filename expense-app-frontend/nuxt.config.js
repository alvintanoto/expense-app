// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
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