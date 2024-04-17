<template>
  <template v-if="!isLoaded">
    <LoadingSpinner />
  </template>
  <template v-else>
    <Container>
      <div class="max-w-screen-lg mx-auto">
        <div class="flex flex-col gap-6">
          <ListingHead :id="listing?.id!" :title="listing?.title!" :imageSrc="listing?.imageSrc!"
            :location="listing?.location!" :country="listing?.country!" />
          <div class="grid grid-cols-1 md:grid-cols-7 md:gap-10 mt-6">
            <div class="col-span-4 flex flex-col gap-8">
              <div class="flex flex-col gap-2">
                <div class="text-xl font-semibold flex flex-row items-center gap-2">
                  <div>
                    Hosted by {{ listing?.username }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </Container>
  </template>
</template>

<script setup lang="ts">

import { ref, watchEffect, computed, type Ref } from "vue"
import { type Listing } from "@/types/listing"
import { client } from "@/lib/client"
import { useRoute, useRouter } from "vue-router"
import type { OkResponseResult } from "@/types/response";
import LoadingSpinner from "@/components/LoadingSpinner.vue"
import Container from "@/components/Container.vue"
import ListingHead from "@/components/ListingHead.vue"
import { CATEGORIES } from "@/constants/categories";


const listing: Ref<Listing | null> = ref(null)
const isLoaded: Ref<boolean> = ref(false)
const route = useRoute()
const router = useRouter()
const category = computed(() => {
  return CATEGORIES.filter((cat) => cat.label == listing.value?.category)
})
watchEffect(async () => {
  await client.get<OkResponseResult<Listing>>(`listing/${route.params.id}`).then((r) => {
    listing.value = r.data.data
  }).catch(() => {
    router.replace("/")
  }).finally(() => {
    isLoaded.value = true
  })
})
</script>
