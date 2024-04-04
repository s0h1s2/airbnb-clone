<template>
  <div class="w-full relative">
    <div>
      <v-icon scale="1.4" class="text-neutral-700 absolute top-5 left-2" v-if="formatPrice" name="bi-currency-dollar" />
    </div>
    <input v-model="value" :name="props.name" :type="props.type" :disabled="props.disabled" placeholder=" "
      class="peer w-full p-4 pt-6 font-light bg-white border-2 rounded-md outline-none transition disabled:opacity-70 disabled:cursor-not-allowed"
      :class="[errorMessage ? 'border-rose-500 border-[1px]' : 'outline-none', formatPrice ? 'pl-9' : 'pl-4']" />
    <label :for="props.name"
      class="absolute text-md duration-150 transform -translate-y-3 top-5 z-10 origin-[0] peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-4"
      :class="[errorMessage ? 'text-rose-500' : 'text-zing-400', formatPrice ? 'left-9' : 'left-4']">
      {{ label }}
    </label>
    <p class="mt-2 text-rose-500">
      {{ errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue"
import { useField } from "vee-validate"

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
