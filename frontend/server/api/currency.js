const config = useRuntimeConfig();

export default defineEventHandler(async (event) => {
  const response = await $fetch(
    config.baseUrl + config.currencyEndpoint,
    "get"
  )
  return response
});
