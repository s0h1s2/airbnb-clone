<template>
  <div class="flex flex-col items-center justify-center gap-2 p-3 border-b-2 hover:text-neutral-800 transition cursor-pointer"
    :class="[selected ? 'border-b-neutral-800 text-neutral-800' : 'border-transparent text-neutral-500']">
    <v-icon :scale="1" :name="iconName"></v-icon>
    <div class="font-medium text-sm">
        {{ label }}
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  label: string,
  iconName: string,
  selected?: boolean
}>()

</script>

<style scoped></style>
