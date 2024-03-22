<template>
  <div :class="[props.center?'text-center':'text-start']">
    <div class="text-2xl font-bold">
      {{ title }}
    </div>
    <div class="font-light text-neutral-500 mt-2">
      {{ subtitle }}
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string
  subtitle?: string
  center?:boolean
}
const props=defineProps<Props>()

</script>

<style scoped></style>
