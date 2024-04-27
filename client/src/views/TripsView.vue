<template>
  <LoadingSpinner v-if="isLoading" />
  <Container v-else>
    <Heading title="Trips" subtitle="Where you've been and where you'are going" />
    <div class="
      mt-10 
      grid 
      grid-cols-1 
      sm:grid-cols-2 
      md:grid-cols-3 
      lg:grid-cols-4 
      xl:grid-cols-5 
      2xl:grid-cols-6 
      gap-8">
      <ListingCard v-for="trip in trips" :listing="trip"
        :reservation="{ startDate: trip.startDate, endDate: trip.endDate }" />
    </div>
  </Container>
</template>

<script setup lang="ts">
import { ref, type Ref } from "vue"
import Container from "@/components/Container.vue"
import Heading from "@/components/Heading.vue"
import ListingCard from "@/components/ListingCard.vue"
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import type { Trips } from "@/types/trips"
import { client } from "@/lib/client"
import type { OkResponseResult } from "@/types/response";
import { useToast } from "vue-toastification"
const toast = useToast()
const trips: Ref<Trips[]> = ref([])
const isLoading = ref(true)
client.get<OkResponseResult<Trips[]>>("/listing/trips").then((r) => {
  trips.value = r.data.data
}).catch(() => {
  toast.warning("Unable to load trips.")
}).finally(() => {
  isLoading.value = false
})
</script>

<style scoped></style>
