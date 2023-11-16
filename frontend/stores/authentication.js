export const useAuthStore = defineStore("auth", {
  state: () => ({
    isLoggedIn: false,
    accessToken: "",
    refreshToken: "",
  }),
  actions: {
  },
  persist: true
});
