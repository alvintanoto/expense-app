export default defineNuxtRouteMiddleware(async (to, from) => {
  const accessTokenCookie = useCookie("access_token");
  const refreshTokenCookie = useCookie("refresh_token");

  if (!accessTokenCookie.value) {
    if (to.path == "/login" || to.path == "/register" || to.path == "/") {
      return;
    } else {
      return navigateTo("/login");
    }
  }

  const ctData = await useFetch("/api/auth/check_token", {
    headers: {
      authorization: "Bearer " + accessTokenCookie.value,
    },
  });

  if (ctData.data?.value.code === "40101") {
    const rtData = await useFetch("/api/auth/refresh_token", {
      headers: {
        authorization: "Bearer " + accessTokenCookie.value,
      },
    });

    if (rtData?.data?.value?.access_token) {
      accessTokenCookie.value = rtData.data.value.access_token;
      refreshTokenCookie.value = rtData.data.value.refresh_token;
    } else {
      accessTokenCookie.value = null;
      refreshTokenCookie.value = null;
    }

  }

  if (ctData.data?.value.code === "40103") {
    // fail refreshing data
    accessTokenCookie.value = null;
    refreshTokenCookie.value = null;
  }

  if (ctData.data?.value.code === "40104") {
    // fail refreshing data
    accessTokenCookie.value = null;
    refreshTokenCookie.value = null;
  }

  if (!accessTokenCookie.value) {
    if (to.path == "/login" || to.path == "/register" || to.path == "/") {
      return;
    }
    return navigateTo("/login");
  }

  if (accessTokenCookie.value) {
    if (to.path == "/login" || to.path == "/register" || to.path =="/") {
      return navigateTo("/transaction");
    }
  }
});
