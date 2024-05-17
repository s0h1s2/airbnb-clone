<template>
  <div @click="filterModal.onOpen()"
    class="border-[1px] w-full md:w-auto py-2 rounded-full shadow-sm hover:shadow-md transition cursor-pointer">
    <div class="flex flex-row items-center justify-between">
      <div class="text-sm font-semibold px-6">
        {{ place }}
      </div>
      <div class="hidden sm:block text-sm font-semibold px-6 border-x-[1px] flex-1 text-center">
        {{ week }}
      </div>
      <div class="text-sm pl-6 pr-2 text-gray-600 flex flex-row items-center gap-3">
        <div class="hidden sm:block">{{ guests == 0 ? "Add Guests" : guests + " Guests" }}</div>
        <div class="p-2 bg-rose-500 rounded-full text-white">
          <v-icon name="oi-search" scale="1.1" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useFilterModalStore } from '@/stores/filterModalStore';
import { useCountriesStore } from '@/stores/countriesStore';
import { ref, watch } from "vue"
import { useRoute } from 'vue-router';
import { formatDistanceStrict, parseISO } from "date-fns"
const filterModal = useFilterModalStore()
const countriesStore = useCountriesStore()
const place = ref("Anywhere")
const week = ref("Anyweek")
const guests = ref(0)
const route = useRoute()
watch(() => [route.query, route.query.country, route.query.startDate, route.query.endDate, route.query.guests], () => {
  if (route.query.country != undefined) {
    place.value = countriesStore.getByValue(route.query.country?.toString() || "")?.label.toString() || "Anywhere"
  } else {
    place.value = "Anywhere"
  }
  if (route.query.startDate != undefined && route.query.endDate != undefined) {
    route.query.startDate = route.query.startDate.toString().replace(/['"]+/g, '')
    route.query.endDate = route.query.endDate.toString().replace(/['"]+/g, '')
    week.value = formatDistanceStrict(parseISO(route.query.startDate), parseISO(route.query.endDate))
  } else {
    week.value = "Anyweek"
  }
  if (route.query.guests != undefined) {
    guests.value = parseInt(route.query.guests.toString()) || 0
  } else {
    guests.value = 0
  }
})
</script>

<style scoped></style>
