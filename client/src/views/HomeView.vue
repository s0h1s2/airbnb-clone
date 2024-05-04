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

const route = useRoute()

const scrollContainer = ref(null)

listingStore.getListings()
useInfiniteScroll(scrollContainer, () => {
  listingStore.getListings()
}, { distance: 10 })
watch(() => route.query.category, () => {
  listingStore.getListings({ category: route.query.category?.toString() })
})

</script>

<template>
  <LoadingSpinner v-if="!listingStore.isLoaded" />
  <template v-else>
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
</template>
<style scoped>
body {
  margin: 0;
}
</style>
