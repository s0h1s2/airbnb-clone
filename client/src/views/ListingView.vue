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
                    Hosted by {{ listing?.user.name }}
                  </div>
                  <div>
                    <v-icon scale="1.2" name="bi-person-circle"></v-icon>
                  </div>
                </div>
                <div class="flex flex-row items-center gap-4 font-light text-neutral-500">
                  <div>
                    {{ listing?.guestCount }} guests
                  </div>
                  <div>

                    {{ listing?.roomCount }} rooms
                  </div>
                  <div>
                    {{ listing?.roomCount }} bathrooms
                  </div>

                </div>
              </div>
              <hr />
              <div class="flex flex-col gap-6">
                <div class="flex flex-row items-center gap-4">
                  <v-icon class="text-neutral-600" :scale="2" :name="category?.iconName" />
                  <div class="flex flex-col">
                    <div class="text-lg font-semibold">
                      {{ category?.label }}
                    </div>
                    <div class="text-neutral-500 font-light">
                      {{ category?.description }}
                    </div>
                  </div>
                </div>

              </div>
              <hr />
              <div class="text-lg font-light text-neutral-500">
                {{ listing?.description }}
              </div>
              <hr />
              <div class="w-[400px] h-[360px]">
                <WorldMap :lat="Number(listing?.location.lat!)" :lan="Number(listing?.location.lng!)" />
              </div>
            </div>
            <div class="order-first mb-10 md:order-last md:col-span-3">
              <div class="bg-white rounded-xl border-[1px] border-neutral-200 overflow-hidden">
                <div class="flex flex-row items-center gap-1 p-4">
                  <div class="text-2xl font-semibold">
                    $ {{ listing?.price }}
                  </div>
                  <div class="font-light text-neutral-600">
                    Night
                  </div>
                </div>
                <hr />
              </div>
              <ReservationDatePicker />
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
import ReservationDatePicker from "@/components/ReservationDatePicker.vue"
import WorldMap from "@/components/WorldMap.vue"
import { CATEGORIES } from "@/constants/categories";


const listing: Ref<Listing | null> = ref(null)
const isLoaded: Ref<boolean> = ref(false)
const route = useRoute()
const router = useRouter()
const category = computed(() => {
  return CATEGORIES.find((cat) => cat.label == listing.value?.category)
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
