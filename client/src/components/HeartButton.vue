<template>
  <div @click.stop="makeItFavorite" class="relative hover:opacity-80 transition cursor-pointer">
    <v-icon scale="1.4" class="fill-white absolute -top-[2px] -right-[2px]" name="ri-heart-3-line"></v-icon>
    <v-icon scale="1.2" :class="[isFavorited ? 'fill-rose-500' : 'fill-neutral-500/70']"
      name="ri-heart-3-fill"></v-icon>
  </div>
</template>

<script setup lang="ts">
import type { ListingUserFavorites } from '@/types/listing';
import { client } from "@/lib/client"
import { ref, watch } from 'vue';
import { useToast } from 'vue-toastification';
import { useUserStore } from "@/stores/userStore"
import { storeToRefs } from 'pinia';
const userStore = useUserStore()
const props = defineProps<{ listingId: string | number, favorites: ListingUserFavorites[] }>()
const toast = useToast()
async function makeItFavorite() {
  await client.post(`/listing/${props.listingId}/favorite`).then(() => {
    isFavorited.value = true
    toast.success("Favorited")
  }).catch(() => {
    toast.error("Unable to make it favorite")
  })
}
const { user } = storeToRefs(userStore)
const isFavorited = ref(Boolean(props.favorites.some((fav) => fav.userId == user.value?.id)))
watch(() => [user.value?.id], () => {
  isFavorited.value = Boolean(props.favorites.some((fav) => fav.userId == user.value?.id))
})
</script>

<style lang="scss" scoped></style>
