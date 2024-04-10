<template>
  <div @click="() => { }" class="relative hover:opacity-80 transition cursor-pointer">
    <v-icon scale="1.4" class="fill-white absolute -top-[2px] -right-[2px]" name="ri-heart-3-line"></v-icon>
    <v-icon scale="1.2" :class="[isFavorited ? 'fill-rose-500' : 'fill-neutral-500/70']"
      name="ri-heart-3-fill"></v-icon>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{ listingId: string | number }>()
const isFavorited = ref(false)
</script>

<style lang="scss" scoped></style>
