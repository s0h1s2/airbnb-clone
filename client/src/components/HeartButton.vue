<template>
  <div @click.stop="" class="relative hover:opacity-80 transition cursor-pointer">
    <v-icon scale="1.4" class="fill-white absolute -top-[2px] -right-[2px]" name="ri-heart-3-line"></v-icon>
    <v-icon scale="1.2" :class="[isFavorited ? 'fill-rose-500' : 'fill-neutral-500/70']"
      name="ri-heart-3-fill"></v-icon>
  </div>
</template>

<script setup lang="ts">
import type { ListingUserFavorites } from '@/types/listing';
import { ref, watch } from 'vue';

const props = defineProps<{ listingId: string | number, favorites: ListingUserFavorites[], userId?: number }>()
const isFavorited = ref(false)
watch(() => props.userId, (newUserId) => {
  isFavorited.value = Boolean(props.favorites.some((fav) => fav.userId == newUserId))

})
</script>

<style lang="scss" scoped></style>
