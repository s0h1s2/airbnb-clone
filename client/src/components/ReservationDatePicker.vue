<template>
  <div class="order-first mb-10 md:order-last md:col-span-3">
    <div class="bg-white rounded-xl border-[1px] border-neutral-200 overflow-hidden">
      <div class="flex flex-row items-center gap-1 p-4">
        <div class="text-2xl font-semibold">
          $ {{ price }}
        </div>
        <div class="font-light text-neutral-600">
          Night
        </div>
      </div>
      <hr />
      <DatePicker class="rounded-none" v-model.range="date" :attributes="attrs" :disabled-dates="disabledDates"
        expanded />

      <Button label="Reverse" outline></Button>
    </div>
  </div>

</template>

<script setup lang="ts">

import Button from "@/components/Button.vue"
import 'v-calendar/style.css';
import { DatePicker } from "v-calendar"

import { ref, type Ref } from "vue"
import type { Reservation } from "@/types/reservation"
const props = defineProps<{ price: number, reservations: Reservation[] }>()
const date = ref({
  start: new Date(),
  end: new Date()
})
const disabledDates = ref([...props.reservations])
const attrs = ref({
  highlight: true
})
</script>

<style lang="css" scoped></style>
