<template>
  <LoadingSpinner v-if="isLoading" />
  <EmptyState v-else-if="trips.length == 0" title="No trips found" subtitle="You haven't any trips."
    :showReset="false" />
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
        :reservation="{ startDate: trip.startDate, endDate: trip.endDate }" actionLabel="Cancel Trip"
        @onClick="cancelTrip(trip.id)" />
    </div>
  </Container>
</template>

<script setup lang="ts">
import { ref, type Ref } from "vue"
import Container from "@/components/Container.vue"
import Heading from "@/components/Heading.vue"
import ListingCard from "@/components/ListingCard.vue"
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import EmptyState from "@/components/EmptyState.vue"
import type { Trips } from "@/types/trips"
import { client } from "@/lib/client"
import type { OkResponseResult } from "@/types/response";
import { useToast } from "vue-toastification"
const toast = useToast()
const trips: Ref<Trips[]> = ref([])
const isLoading = ref(true)
function cancelTrip(id: string | number) {
  client.delete(`/listing/${id}/cancel_reserve`).then(() => {
    trips.value = trips.value.filter((trip) => trip.id != id)
    toast.success("Reserve cancelled successfuly")
  }).catch(() => {
    toast.error("Unable to cancel reserve")
  })
}
client.get<OkResponseResult<Trips[]>>("/listing/trips").then((r) => {
  trips.value = r.data.data
}).catch(() => {
  toast.warning("Unable to load trips.")
}).finally(() => {
  isLoading.value = false
})
</script>

<style scoped></style>
