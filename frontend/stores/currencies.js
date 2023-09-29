import { defineStore } from "pinia";

export const useCurrenciesStore = defineStore("currencies", {
  state: () => ({
    currencies: [],
    filter: "", 
  }),
  actions: {
    async getCurrencies() {
      const config = useRuntimeConfig();
      if (this.currencies.length > 0) {
        return;
      }

      const api = config.baseUrl + config.currencyEndpoint;

      await fetch(api)
        .then((response) => response.json())
        .then(({ data }) => {
          return (this.currencies = data);
        })
        .catch((error) => console.log(error));
    
        return
    },
    setCurrencyFilter(filter) {
        this.filter = filter
    }
  },
  getters: {
    currencyList(state) {
      return state.currencies;
    },
    displayedCurrencyList(state) {
        if (state.filter === "") {
            return this.currencies
        }

        return state.currencies.filter((item) => {
            if (
              item.currency_code.includes(state.filter) ||
              item.currency_name.includes(state.filter)
            ) {
              return item;
            }
        
            return;
        })
    }
  },
});
