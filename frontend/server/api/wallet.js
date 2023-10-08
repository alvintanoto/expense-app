export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const accessToken = getCookie(event, "access_token");
  const api = config.baseUrl + config.createWalletEndpoint;

  if (event.method == "GET") {
    const data = await $fetch(api, {
      method: event.method,
      headers: {
        authorization: "Bearer " + accessToken,
      },
    })
      .catch((err) => {
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
  }

  if (event.method == "POST") {
    const data = await $fetch(api, {
      method: event.method,
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
  }

  setResponseStatus(event, 405);
  return {
    code: "40500",
    message: "Method not allowed",
    data: null,
  };
});
