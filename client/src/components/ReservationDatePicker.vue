<template>
  <div>
    <VueDatePicker v-model="date" inline range></VueDatePicker>
  </div>
</template>

<script setup lang="ts">
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import { ref, type Ref } from "vue"
const date: Ref<[Date, Date]> = ref([new Date(), new Date()])
</script>

<style scoped></style>
