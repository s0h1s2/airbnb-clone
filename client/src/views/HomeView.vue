<script setup lang="ts">
import Container from "@/components/Container.vue"
import EmptyState from "@/components/EmptyState.vue"
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import ListingCard from "@/components/ListingCard.vue"

import { useListingStore } from "@/stores/listingStore"
import { ref, type Ref } from "vue";

const isEmpty: Ref<boolean> = ref(true)
const listingStore = useListingStore()
listingStore.getListings()

</script>

<template>
  <LoadingSpinner v-if="!listingStore.isLoaded" />
  <template v-else>
    <EmptyState v-if="listingStore.listings.length == 0" />
    <Container v-else>
      <div
        class="pt-24 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6 gap-8">
        <ListingCard :listing="listing" v-for="listing in listingStore.listings"></ListingCard>
      </div>
    </Container>
  </template>
</template>
