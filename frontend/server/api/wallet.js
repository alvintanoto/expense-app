export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const accessToken = getCookie(event, "access_token");
  const api = config.baseUrl + config.createWalletEndpoint;

  const data = await $fetch(api, {
    method: "POST",
    headers: {
      authorization: "Bearer " + accessToken,
    },
    body: await readBody(event),
  }).catch((err) => {
    if (err.statusCode) {
      setResponseStatus(event, err.statusCode);
      return err.data;
    }

    setResponseStatus(event, 503);
    return {
      code: "50300",
      message: "Server unavailable, please try again later",
      data: null,
    };
  });

  return data;
});
