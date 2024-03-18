<script setup>
const props = defineProps < {
  label: string,
  onClick: () => void
} > ()
</script>


<template>
  <div click="onClick" class="px-4 py-3 hover:bg-neutral-100 transition font-semibold">
    {{ label }}
  </div>
</template>


<style lang="scss" scoped></style>
