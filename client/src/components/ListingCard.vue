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
      <div class="font-semibold text-lg"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { format } from "date-fns"
import { useRouter } from 'vue-router';
import HeartButton from "./HeartButton.vue"
import type { Listing } from '@/types/listing';
const props = defineProps<{
  listing: Listing
  reservation?: {
    startDate: string
    endDate: string
  }
}>()
const router = useRouter()
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
