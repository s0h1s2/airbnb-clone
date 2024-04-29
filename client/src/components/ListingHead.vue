<template>
  <Heading :title="title" :subtitle="`${country?.region}, ${country?.label}`" />
  <div class="w-full h-[60vh] overflow-hidden rounded-xl relative">
    <img :src="imageSrc" alt="Image" class="object-cover w-full">
    <div class="absolute top-5 right-5">
      <HeartButton :listingId="id" :favorites="favorites" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Heading from "./Heading.vue"
import HeartButton from "./HeartButton.vue"
import type { PropertyLocation } from '@/types/location';
import { useCountriesStore } from "@/stores/countriesStore"
import type { ListingUserFavorites } from "@/types/listing";
const props = defineProps<{
  title: string
  imageSrc: string
  id: string | number
  country: string
  location: PropertyLocation
  favorites: ListingUserFavorites[]
}>()
const countryStore = useCountriesStore()
const country = countryStore.getByValue(props.country)
</script>

<style scoped></style>
