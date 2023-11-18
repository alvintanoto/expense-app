export async function fetchWallets() {
    const authStore = useAuthStore();
    const walletStore = useWalletStore();

    const {data, pending, error, refresh} = await useFetch("/api/v1/user/wallet", {
        onRequest({request, options}) {
            options.method = "GET"
            
            if (authStore.accessToken) {
                options.headers.authorization = "Bearer " + authStore.accessToken
            }
        }
    })

    if (error?.value) {
        if (error?.value?.data?.code === '40101') {
            return [null, error.value.data]
        }

        if (error?.value?.data?.client_message) {
            return [null, error.value.data]
        }

        return [null, "Could not connect to server, please try again later"]
    }

    return [data.value.data, null]
}