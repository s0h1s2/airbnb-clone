<template>
  <div class="w-full relative">
    <input v-model="value" :name="props.name" :type="props.type" :disabled="props.disabled" placeholder=" "
      class="peer w-full p-4 pt-6 font-light bg-white border-2 rounded-md outline-none transition disabled:opacity-70 disabled:cursor-not-allowed"
      :class="[errorMessage ? 'border-rose-500 border-[1px]' : 'outline-none']" />
    <label :for="props.name"
      class="absolute text-md duraiton-150 transform -translate-y-3 top-5 z-10 origin-[0] left-4 peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4 "
      :class="[errorMessage ? 'text-rose-500' : 'text-zing-400']">
      {{ label }}
    </label>
    <p class="mt-2 text-rose-500">
      {{ errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue"
import { Field, useField } from "vee-validate"

interface Props {
  name: string
  label: string
  placeholder?: string
  type?: string
  formatPrice?: boolean
  disabled?: boolean
  required?: boolean
}
const props = withDefaults(defineProps<Props>(), {
  type: "text"
})

const { value, errorMessage } = useField(() => props.name)

</script>

<style scoped></style>
