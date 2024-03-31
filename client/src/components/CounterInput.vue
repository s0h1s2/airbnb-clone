<template>
  <div class="flex flex-row items-center justify-between">
    <div class="flex flex-col">
      <div class="font-medium">
        {{ title }}
      </div>
      <div class="font-light text-gray-600">
        {{subtitle}}
      </div>
    </div>
    <div class="flex flex-row items-center gap-4">
        
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
interface Props{
  title:string
  subtitle:string
}
defineProps<Props>()

const emit=defineEmits<{
  onChange:[number]
}>()


const counter=ref(1)
function increaseCounter(){
  counter.value++
  emit("onChange",counter.value)
}
function decreaseCounter(){
  if (counter.value>1) {
    counter.value--
    emit("onChange",counter.value)
  }
}
</script>

<style scoped>

</style>
