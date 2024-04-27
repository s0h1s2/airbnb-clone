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
      <DatePicker class="rounded-none" :min-date="new Date()" v-model.range="date" :attributes="attrs"
        :disabled-dates="disabledDates" expanded />
      <hr />
      <div class="p-4">
        <Button :disabled="isReserving" @onClick="reserve" label="Reserve"></Button>
      </div>
      <div class="p-4 flex flex-row items-center justify-between font-semibold text-lg">
        <div>
          Total
        </div>
        <div>
          $ {{ totalPrice }}
        </div>
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">

import Button from "@/components/Button.vue"
import 'v-calendar/style.css';
import { differenceInCalendarDays } from "date-fns"
import { DatePicker } from "v-calendar"

import { ref, watch } from "vue"
import type { Reservation } from "@/types/reservation"
import { useListingStore } from "@/stores/listingStore";
import { useUserStore } from "@/stores/userStore";
import { useLoginModalStore } from "@/stores/loginModal";
import { useToast } from "vue-toastification";

const props = defineProps<{ id: string | number, price: number, reservations: Reservation[] }>()
const disabledDates = ref([...props.reservations])
const date = ref({
  start: new Date(),
  end: new Date()
})
watch(date, () => {
  const days = differenceInCalendarDays(date.value.end, date.value.start) + 1
  totalPrice.value = days * props.price
})
const listingStore = useListingStore()
const userStore = useUserStore()
const loginModalStore = useLoginModalStore()
const isReserving = ref(false)
const toast = useToast()
function reserve() {
  if (!userStore.isAuth) {
    loginModalStore.onOpen()
    return
  }
  isReserving.value = true
  const reserveDate = { start: date.value.start, end: date.value.end }
  listingStore.reserveListing(props.id, reserveDate).then(() => {
    disabledDates.value.push(reserveDate)
    toast.success("Reserve listing successful.")
  }).catch(() => {
    toast.error("Unable to reserve listing")
  })
  isReserving.value = false
}
const totalPrice = ref(props.price)
const attrs = ref({
  highlight: true,
  popover: {
    visiability: "hover"
  }
})
</script>

<style lang="css" scoped></style>
