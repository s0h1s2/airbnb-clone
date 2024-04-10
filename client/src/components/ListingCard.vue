<template>
  <div @click="router.push(`/listing/${listing.id}`)" class="col-span-1 cursor-pointer group">
    <div class="flex flex-col gap-2 w-full">
      <div class="aspect-square w-full relative overflow-hidden rounded-xl">
        <img :src="listing.imageSrc" class="object-cover w-full h-full group-hover:scale-110 transition"
          alt="Property photo">
        <div class="absolute top-3 right-3">
          <HeartButton :listing-id="1" />
        </div>
      </div>
      <div class="font-semibold text-lg">
        {{ countryAndRegion }}
      </div>
      <div class="font-light text-neutral-500">
        {{ reservationDate || listing.category }}
      </div>
      <div class="flex flex-row items-center gap-1">
        <div class="font-semibold">
          {{ totalPrice }}
        </div>
        <div v-if="!reservation" class="font-light">Night</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { format } from "date-fns"
import { useRouter } from 'vue-router';
import HeartButton from "./HeartButton.vue"
import type { Listing } from '@/types/listing';
import { useCountriesStore } from "@/stores/countriesStore"
const countryStore = useCountriesStore()
const props = defineProps<{
  listing: Listing
  reservation?: {
    startDate: string
    endDate: string
  }
}>()


const country = countryStore.getByValue(props.listing.country)

const router = useRouter()
const countryAndRegion = computed(() => {
  return `${country?.region},${country?.label}`
})
const totalPrice = computed(() => {
  return `${props.listing.price}$`
})
const reservationDate = computed(() => {
  if (!props.reservation) {
    return
  }
  const start = new Date(props.reservation.startDate)
  const end = new Date(props.reservation.endDate)
  return `${format(start, 'pp')} - ${format(end, 'pp')}`
})

</script>

<style scoped></style>
