export async function fetchCurrencies() {
    const config = useRuntimeConfig();
    const authStore = useAuthStore();
    const globalStore = useGlobalStore();

    const { data, pending, error, refresh } = await useFetch(config.public.currencyEndpoint,
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${authStore.accessToken}`
            }
        })

    if (error?.value) {
        if (error?.value?.data?.code == '40104' || error?.value?.data?.code == '40101') {
            globalStore.forceLogout = true;
            globalStore.errorMessage = "Session expired, please login again";
            return [null, error.value.data]
        }

        if (error?.value?.data?.client_message) {
            globalStore.errorMessage = error.value.data.client_message;
            return [null, error.value.data]
        }

        return [null, "Could not connect to server, please try again later"]
    }

    return [data.value.data, null]
}