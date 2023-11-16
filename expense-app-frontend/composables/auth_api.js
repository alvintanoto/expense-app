
export async function login (username, password) {
    const {data, pending, error, refresh} = await useFetch("/api/v1/authentication/login", {
        onRequest({request, options}) {
            options.method = "POST"
            options.body = {
                "username": username,
                "password": password
            }
        }, 
    })

    if (error?.value) {
        if (error?.value?.data?.client_message) {
            return [null, error.value.data.client_message]
        }

        return "Could not connect to server, please try again later"
    }
    
    const store = useAuthStore();
    store.accessToken = data.value.data.access_token;
    store.refreshToken = data.value.data.refresh_token;
    return null
}