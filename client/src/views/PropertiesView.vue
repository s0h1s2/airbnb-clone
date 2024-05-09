<template>
  <LoadingSpinner v-if="isLoading" />
  <EmptyState v-else-if="properties === null || properties?.length == 0" title="No property found"
    subtitle="Looks like you don't have any property" :showReset="false" />
  <Container v-else>
    <Heading title="Properties" subtitle="Manage your properties" />
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
      <ListingCard v-for="property in properties" :listing="property" actionLabel="Delete property"
        @onClick="deleteProperty(property.id)" />
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
import type { Listing } from "@/types/listing"
import { client } from "@/lib/client"
import type { OkResponseResult } from "@/types/response";
import { useToast } from "vue-toastification"
const toast = useToast()
const properties: Ref<Listing[]> = ref([])
const isLoading = ref(true)
function deleteProperty(id: string | number) {
  client.delete(`/listing/${id}`).then(() => {
    properties.value = properties.value.filter((property) => property.id != id)
    toast.success("Property deleted successfuly")
  }).catch(() => {
    toast.error("Unable to delete property")
  })
}
client.get<OkResponseResult<Listing[]>>("/listing/properties").then((r) => {
  properties.value = r.data.data
}).catch(() => {
  toast.warning("Unable to load properties.")
}).finally(() => {
  isLoading.value = false
})
</script>

<style scoped></style>
