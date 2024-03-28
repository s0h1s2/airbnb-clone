<template>
  <v-select title="Countries" :options="countries.getAll"></v-select>
</template>

<script setup lang="ts">
import { useCountriesStore } from "@/stores/countriesStore.ts"
const countries = useCountriesStore()

</script>

<style scoped></style>
