export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  let { data } = await $fetch(config.baseUrl + config.checkTokenEndpoint, {
    headers: event.headers
  }).catch((err) => {
    return err;
  });

  return data;
});
