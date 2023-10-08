export const useCurrenciesStore = defineStore("currencies", {
  state: () => ({
    currencies: [],
    filter: "",
  }),
  actions: {
    async getCurrencies() {
      const token = useCookie("access_token");
      const config = useRuntimeConfig();
      if (this.currencies.length > 0) {
        return;
      }

      const api = config.baseUrl + config.currencyEndpoint;

      await fetch(api, {
        headers: {
          authorization: "bearer " + token.value,
        },
      })
        .then((response) => response.json())
        .then(({ data }) => {
          return (this.currencies = data);
        })
        .catch((error) => console.log(error));

      return;
    },
    setCurrencyFilter(filter) {
      this.filter = filter;
    },
  },
  getters: {
    currencyList(state) {
      return state.currencies;
    },
    displayedCurrencyList(state) {
      if (state.filter === "") {
        return this.currencies;
      }

      return state.currencies.filter((item) => {
        if (
          item.currency_code
            .toLowerCase()
            .includes(state.filter.toLowerCase()) ||
          item.currency_name.toLowerCase().includes(state.filter.toLowerCase())
        ) {
          return item;
        }

        return;
      });
    },
  },
  persist: true
});
