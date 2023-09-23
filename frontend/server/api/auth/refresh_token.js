export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  const refreshToken = getCookie(event, "refresh_token");

  let { data } = await $fetch(config.baseUrl + config.refreshTokenEndpoint, {
    method: "POST",
    headers: event.headers,
    body: {
      refresh_token: refreshToken,
    },
  }).catch((err) => {
    return err;
  });

  return data;
});
