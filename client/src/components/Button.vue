<template>
  <button :disabled="disabled" @click.stop="emit('onClick')"
    class="relative disabled:opacity-75 disabled:cursor-not-allowed rounded-lg hover:opacity-80 transition w-full"
    :class="[outline ? 'bg-white border-black text-black' : 'bg-rose-500 border-rose-500 text-white', small ? 'py-1 text-sm font-light border-[1px]' : 'py-3 text-md font-semibold border-2']">
    <slot></slot>
    {{ label }}
  </button>
</template>

<script setup lang="ts">

interface Props {
  label: string
  disabled?: boolean
  outline?: boolean
  small?: boolean
}
defineProps<Props>()
const emit = defineEmits<{ onClick: [] }>()

</script>

<style scoped></style>
