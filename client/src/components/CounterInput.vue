<template>
  <div class="flex flex-row items-center justify-between">
    <div class="flex flex-col">
      <div class="font-medium">
        {{ title }}
      </div>
      <div class="font-light text-gray-600">
        {{ subtitle }}
      </div>
    </div>
    <div class="flex flex-row items-center gap-4">
      <div @click="decreaseCounter()"
        class="w-10 h-10 rounded-full border-[1px] border-neutral-400 flex items-center justify-center text-neutral-600 cursor-pointer hover:opacity-80 transition">
        <v-icon name="hi-solid-minus-sm"></v-icon>
      </div>
      <div class="font-light text-xl text-neutral-600 ">
        {{ count }}
      </div>
      <div @click="increaseCounter()"
        class="w-10 h-10 rounded-full border-[1px] border-neutral-400 flex items-center justify-center text-neutral-600 cursor-pointer hover:opacity-80 transition">
        <v-icon name="hi-solid-plus-sm"></v-icon>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import { useField } from 'vee-validate';
import { onMounted, ref, watch } from 'vue';
interface Props {
  title: string
  subtitle: string
  name: string
}
const props = defineProps<Props>()
const emit=defineEmits<{onValueChange:[number,void]}>()
const { setValue } = useField(() => props.name)
onMounted(() => {
  setValue(1)
})

const count=ref(1)
watch(count,()=>{
  emit("onValueChange",count.value)
})
function increaseCounter() {
  count.value++
  setValue(count.value)
}
function decreaseCounter() {
  if (count.value > 1) {
    count.value--
    setValue(count.value)
  }
}
</script>

<style scoped></style>
