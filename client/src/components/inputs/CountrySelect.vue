<template>
  <v-select placeholder="Anywhere" v-model="selectedCountry" title="Countries" :options="countries.getAll" :clearable="true">
    <template #option="{ label,region,flag }">
      <div class="flex flex-row items-center gap-3">
        <div>
          {{ flag }}
        </div>
        <div>
          {{ label }}
          <span class="text-neutral-500 ml-1">
            {{ region }}
          </span>
        </div>
      </div>
    </template>
  </v-select>
</template>

<script setup lang="ts">
import { useCountriesStore } from "@/stores/countriesStore"
import type { Country } from "@/types/country";
const countries = useCountriesStore()
const selectedCountry=defineModel<Country>()
</script>

<style scoped>
:deep(){
  --vs-search-input-placeholder-color:#A3A3A3;
}
</style>
