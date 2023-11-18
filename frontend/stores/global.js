export const useGlobalStore = defineStore("global", {
    state: () => ({
        forceLogout: false,
        errorMessage: ""
    }),
    actions: {
    },
    persist: true
});
