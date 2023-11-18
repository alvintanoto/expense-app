import { fetchCurrencies } from '~/composables/currency_api'

export const useCurrencyStore = defineStore("currency", {
    state: () => ({
        filter: "",
        currencies: [],
    }),
    actions: {
        async fetchCurrenciesData() {
            if (this.currencies.length !== 0) {
                return null
            }

            // fetch from composable
            const [data, error] = await fetchCurrencies()
            if (error) {
               return error
            } else {
                this.currencies = data
            }

            return null
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
