export const useWalletStore = defineStore("wallets", {
  state: () => ({
    wallets: [],
    filter: "",
  }),
  actions: {
    async createWallet(walletName, currencyID, initialBalance) {
      let result = {
        data: null,
        error: null,
      }

      const api = "/api/wallet";

      const data = await $fetch(api, {
        method: "POST",
        body: {
          wallet_name: walletName,
          currency_id: currencyID,
          initial_balance: initialBalance,
        },
      }).catch((err) => {
        result.error = err.data
        return result
      });

      result.data = data
      return result
    },
  },
  getters: {},
});
