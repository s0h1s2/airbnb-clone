<script setup lang="ts">
import Container from "@/components/Container.vue"
import EmptyState from "@/components/EmptyState.vue"
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import ListingCard from "@/components/ListingCard.vue"

import { useListingStore } from "@/stores/listingStore"
import { useRoute } from "vue-router"
import { watch, ref } from "vue"
import { useInfiniteScroll } from "@vueuse/core"
const listingStore = useListingStore()
const firstTimeLoad = ref(true)
const route = useRoute()

const scrollContainer = ref(null)

listingStore.getListings()
useInfiniteScroll(scrollContainer, () => {
  if (listingStore.page >= listingStore.totalPage) {
    return
  }
  listingStore.getListings()
}, { distance: 10 })
watch(() => [route.query, route.query.category, route.query.guests, route.query.bathrooms, route.query.rooms, route.query.endDate, route.query.startDate, route.query.country], () => {
  firstTimeLoad.value = true
  listingStore.$reset()
  listingStore.getListings({ ...route.query })
})
</script>

<template>
  <template v-if="listingStore.isLoaded">
    <EmptyState v-if="listingStore.listings == null || listingStore.listings.length === 0" />
    <div class="h-screen overflow-y-scroll" ref="scrollContainer" v-else>
      <Container>
        <div
          class="pt-24 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6 gap-8">
          <ListingCard :listing="listing" v-for="listing in listingStore.listings" :key="listing.id"></ListingCard>
        </div>
      </Container>
    </div>
  </template>
  <LoadingSpinner v-else />
</template>
<style scoped>
body {
  width: 100%;
  overflow-x: hidden;
  overflow-y: hidden;
}
</style>
