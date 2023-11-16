export async function fetchCurrencies() {
    console.log("fetchCurrencies")
    const authStore = useAuthStore();
    const globalStore = useGlobalStore();

    const {data, pending, error, refresh} = await useFetch("/api/v1/currencies", {
        onRequest({request, options}) {
            options.method = "GET"
            options.headers.authorization = "Bearer " + authStore.accessToken
        }
    })

    if (error?.value) {
        if (error?.value?.data?.code === '40101') {
            globalStore.error = error.value.data.client_message
            return [null, error.value.data]
        }

        if (error?.value?.data?.client_message) {
            return [null, error.value.data]
        }

        return [null, "Could not connect to server, please try again later"]
    }

    return [data.value.data, null]
}