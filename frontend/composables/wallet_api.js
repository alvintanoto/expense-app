export async function fetchWallets() {
    const config = useRuntimeConfig();
    const authStore = useAuthStore();
    const globalStore = useGlobalStore();

    const { data, pending, error, refresh } = await useFetch(config.public.getWalletEndpoint,
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

export async function createWallet(walletName, currencyID, initialBalance) {
    const config = useRuntimeConfig();
    const authStore = useAuthStore();
    const globalStore = useGlobalStore();

    const { data, pending, error, refresh } = await useFetch(config.public.createWalletEndpoint,
        {
            method: "POST",
            headers: {
                Authorization: `Bearer ${authStore.accessToken}`
            },
            body: {
                "wallet_name": walletName,
                "currency_id": currencyID,
                "initial_balance": initialBalance
            }
        })

    if (error?.value) {
        if (error?.value?.data?.code == '40104' || error?.value?.data?.code == '40101') {
            globalStore.forceLogout = true;
            globalStore.errorMessage = "Session expired, please login again";
            return "Session expired, please login again"
        }

        if (error?.value?.data?.client_message) {
            globalStore.errorMessage = error.value.data.client_message;
            return error.value.data.client_message
        }

        return "Could not connect to server, please try again later"
    }

    return null
}