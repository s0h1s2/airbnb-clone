<template>
  <LoadingSpinner v-if="!isLoaded" />
  <EmptyState v-if="favorites == null" title="No favorites found" subtitle="Looks like you don't have any favorites"
    :showReset="false" />
  <Container v-else>
    <Heading title="Favorites" subtitle="List of places you favorited!" />
    <div class="mt-10 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6">
      <ListingCard v-for="favorite in favorites" :listing="favorite" />
    </div>
  </Container>
</template>

<script setup lang="ts">
import EmptyState from "@/components/EmptyState.vue"
import ListingCard from "@/components/ListingCard.vue"
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import Container from "@/components/Container.vue"
import Heading from "@/components/Heading.vue"
import { client } from "@/lib/client"
import { type Listing } from "@/types/listing"
import { ref, type Ref } from "vue"
import { type OkResponseResult } from "@/types/response"
const isLoaded = ref(false)
const favorites: Ref<Listing[] | null> = ref(null)

client.get<OkResponseResult<Listing[]>>("/listing/favorites").then((r) => {
  favorites.value = r.data.data
}).catch(() => {
}).finally(() => {
  isLoaded.value = true
})

</script>

<style scoped></style>
