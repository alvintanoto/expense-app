
export const useCurrencyStore = defineStore("currency", {
    state: () => ({
        filter: "",
        currencies: [],
    }),
    actions: {
        
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
                  item.currency_code.toLowerCase().includes(state.filter) ||
                  item.currency_name.toLowerCase().includes(state.filter)
                ) {
                  return item;
                }
            
                return;
            })
        }
    },
    persist: true
  });
  