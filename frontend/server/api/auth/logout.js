export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  if (event.node.req.method == "GET") {
    const data = await $fetch(config.baseUrl + config.logoutEndpoint, {
      headers: event.headers,
      method: "GET",
    });

    return data;
  }

  setResponseStatus(event, 405);
  return {
    code: "405",
    message: "method not allowed",
  };
});
