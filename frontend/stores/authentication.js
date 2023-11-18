export const useAuthStore = defineStore("auth", {
    state: () => ({
      accessToken: "",
      refreshToken: "",
    }),
    actions: {
    },
    persist: true
  });
  