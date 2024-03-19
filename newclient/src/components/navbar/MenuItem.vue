<script setup lang="ts">
interface Props{
  label:string
  onClick: () => void
}
const props = defineProps<Props>()
</script>

<template>
  <div click="onClick" class="px-4 py-3 hover:bg-neutral-100 transition font-semibold">
    {{ label }}
  </div>
</template>


