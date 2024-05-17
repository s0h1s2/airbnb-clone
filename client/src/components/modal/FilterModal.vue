<template>
  <Modal title="Filter" :isOpen="filterModal.isOpen" :disabled="false" :actionLabel="filterLabel"
    @onSubmit="filterSearch" @onClose="filterModal.onClose()" :secondaryActionLabel="backLabel"
    @onSecondaryAction="onBack">
    <template v-slot:body>
      <div v-show="currentStep == Steps.COUNTRY">
        <div class="flex flex-col gap-8">
          <Heading title="Where do you wanna go?" subtitle="Find a perfect location!" />
          <CountrySelect v-model="selectedCountry" />
          <div class="h-[40vh] w-full">
            <WorldMap :lat="selectedCountry?.latlang[0] || 0" :lan="selectedCountry?.latlang[1] || 0" />
          </div>
        </div>
      </div>
      <div v-show="currentStep == Steps.DATE">
        <div class="flex flex-col gap-8">
          <Heading title="When do you plan to go?" subtitle="Make sure everyone is free!" />
          <DatePicker :minDate="new Date()" class="w-full" v-model.range="selectedDates" />
        </div>
      </div>
      <div v-show="currentStep == Steps.ROOMS">
        <div class="flex flex-col gap-8">
          <Heading title="More information" subtitle="Find your perfect place!" />
          <CounterInput @onValueChange="(val) => guests = val" name="guests" title="Guests"
            subtitle="How many guests are coming?" />
          <CounterInput @onValueChange="(val) => rooms = val" name="rooms" title="Rooms"
            subtitle="How many rooms do you need?" />
          <CounterInput @onValueChange="(val) => bathrooms = val" name="bathrooms" title="Bathrooms"
            subtitle="How many bathrooms do you need?" />
        </div>
      </div>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import Modal from "@/components/modal/Modal.vue"
import Heading from "@/components/Heading.vue"
import CountrySelect from "@/components/inputs/CountrySelect.vue"
import WorldMap from "@/components/WorldMap.vue"
import { useFilterModalStore } from "@/stores/filterModalStore";
import { useRouter } from "vue-router"
import { computed, ref, type Ref } from "vue";
import type { Country } from "@/types/country";
import { DatePicker } from "v-calendar"
import { formatISO } from "date-fns"
import CounterInput from "../CounterInput.vue";
const filterModal = useFilterModalStore()
enum Steps {
  COUNTRY,
  DATE,
  ROOMS,
}
const currentStep: Ref<Steps> = ref(Steps.COUNTRY)
const selectedCountry: Ref<Country | undefined> = ref(undefined)
const router = useRouter()
function onBack() {
  if (currentStep.value - 1 < Steps.COUNTRY) {
    return
  }
  currentStep.value--
}
const filterLabel = computed(() => {
  return currentStep.value != Steps.ROOMS ? "Next" : "Search"
})
const backLabel = computed(() => {
  return currentStep.value > Steps.COUNTRY ? "Back" : undefined
})
const selectedDates = ref({
  start: new Date(),
  end: new Date()
})
const rooms = ref(1)
const bathrooms = ref(1)
const guests = ref(1)

interface QueryParams {
  country?: string
  startDate?: string
  endDate?: string
  guests?: number
  bathrooms?: number
  rooms?: number
}
function filterSearch() {
  if (currentStep.value != Steps.ROOMS) {
    currentStep.value++
    return
  }
  const query: QueryParams = {}
  if (selectedCountry?.value != undefined) {
    query.country = selectedCountry.value.value
  }
  query.startDate = formatISO(selectedDates.value.start, { representation: "complete" })
  if (selectedDates.value.end != null) {
    query.endDate = formatISO(selectedDates.value.end, { representation: "complete" })
  }
  query.rooms = rooms.value
  query.bathrooms = bathrooms.value
  query.guests = guests.value
  router.push({ name: "home", query: { ...query } })
  filterModal.onClose()
}
</script>

<style scoped></style>
