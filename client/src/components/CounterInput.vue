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
        {{ counter }}
      </div>
      <div @click="increaseCounter()"
        class="w-10 h-10 rounded-full border-[1px] border-neutral-400 flex items-center justify-center text-neutral-600 cursor-pointer hover:opacity-80 transition">
        <v-icon name="hi-solid-plus-sm"></v-icon>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import { useField, useSetFieldValue } from 'vee-validate';
import { ref } from 'vue';
interface Props {
  title: string
  subtitle: string
  name: string
}
const props = defineProps<Props>()

const { setValue } = useField(() => props.name)
const counter = ref(1)

function increaseCounter() {
  counter.value++
  setValue(counter.value)

}
function decreaseCounter() {
  if (counter.value > 1) {
    counter.value--
    setValue(counter.value)
  }
}
</script>

<style scoped></style>
