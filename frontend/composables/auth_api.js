
export async function login (username, password) {
    const config = useRuntimeConfig();

    const {data, pending, error, refresh} = await useFetch(config.public.loginEndpoint, {
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
            return error.value.data.client_message
        }

        return  "Could not connect to server, please try again later"
    }

    const authStore = useAuthStore();
    authStore.accessToken = data.value.data.access_token;
    authStore.refreshToken = data.value.data.refresh_token;
    return null
}

export async function register (username, email, password) {
    const config = useRuntimeConfig();

    const {data, pending, error, refresh} = await useFetch(config.public.registerEndpoint, {
        onRequest({request, options}) {
            options.method = "POST"
            options.body = {
                "username": username,
                "password": password,
                "email": email, 
            }
        }, 
    })

    if (error?.value) {
        if (error?.value?.data?.client_message) {
            return error.value.data.client_message
        }

        return "Could not connect to server, please try again later"
    }

    const authStore = useAuthStore();
    authStore.accessToken = data.value.data.access_token;
    authStore.refreshToken = data.value.data.refresh_token;
    return null
}