import type { Country } from "@/types/country";
import { defineStore } from "pinia";
import countries from "world-countries";

const formattedCountries = countries.map<Country>((country) => ({
  value: country.cca2,
  label: country.name.common,
  flag: country.flag,
  region: country.region,
  latlang: country.latlng

}))
export const useCountriesStore = defineStore("countries", {
  state: () => ({
    countries: formattedCountries
  }),
  getters: {
    getAll: (state) => state.countries,
    getByValue: (state) => {
      return (value: string) => state.countries.find((country) => country.value == value)
    }
  }

})
