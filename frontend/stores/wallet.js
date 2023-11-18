
export const useWalletStore = defineStore("wallet", {
    state: () => ({
        wallets: []
    }),
    actions: {
        async fetchWallets() {
            // fetch from composable
            const [data, error] = await fetchWallets()
            if (error) {
               return error
            } else {
                this.wallets = data
            }

            return null
        }
    },
    getters: {
        walletList(state) {
            return state.wallets;
        },
    },
    persist: true
  });
  