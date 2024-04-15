<template>
  <div>

  </div>
</template>

<script setup lang="ts">
import type { PropertyLocation } from '@/types/location';
import { useCountriesStore } from "@/stores/countriesStore"
const props = defineProps<{
  title: string
  imageSrc: string
  id: string | number
  country: string
  location: PropertyLocation
}>()
const countryStore = useCountriesStore()
countryStore.getByValue(props.country)
</script>

<style scoped></style>
