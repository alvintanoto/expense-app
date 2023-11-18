import { createWallet } from "~/composables/wallet_api";

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
        },
        async createWallet(walletName, currencyID, initialBalance) {
            return await createWallet(walletName, currencyID, initialBalance)
        }
    },
    getters: {
        walletList(state) {
            return state.wallets;
        },
    },
    persist: true
  });
  