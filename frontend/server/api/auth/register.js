const config = useRuntimeConfig();

export default defineEventHandler(async (event) => {
  if (event.node.req.method == "POST") {
    const data = await $fetch(config.baseUrl + config.registerEndpoint, {
      method: "POST",
      body: await readBody(event),
      onResponse({ request, response, options }) {
        setResponseStatus(event, response.status);
        return response._data.data;
      },
      onResponseError({ request, response, options }) {
        setResponseStatus(event, response.status);
        return response._data.data;
      },
    });

    return data
  }

  setResponseStatus(event, 405);
  return {
    code: "405",
    message: "method not allowed",
  };
});
