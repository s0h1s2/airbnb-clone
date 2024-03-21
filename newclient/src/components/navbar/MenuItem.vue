<script setup lang="ts">
interface Props{
  label:string
}
const props = defineProps<Props>()
const emit=defineEmits<{onClick:[]}>()
</script>

<template>
  <div @click="emit('onClick')" class="px-4 py-3 hover:bg-neutral-100 transition font-semibold">
    {{ props.label }}
  </div>
</template>


