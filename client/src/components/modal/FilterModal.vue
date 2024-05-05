<template>
  <Modal title="Filter" :isOpen="filterModal.isOpen" :disabled="false" :actionLabel="filterLabel" @onSubmit="filterSearch" @onClose="filterModal.onClose()">
    <template v-slot:body>
      <div v-show="currentStep==Steps.COUNTRY">
        Country
      </div>
      <div v-show="currentStep==Steps.DATE">
        Date
      </div >
      <div v-show="currentStep==Steps.ROOMS">
        Rooms
      </div>
    </template>
  </Modal>
</template>

<script setup lang="ts">
import Modal from "@/components/modal/Modal.vue"
import { useFilterModalStore } from "@/stores/filterModalStore";
import { computed, ref, type Ref } from "vue";
const filterModal=useFilterModalStore()
enum Steps{
  COUNTRY,
  DATE,
  ROOMS,
}
const currentStep:Ref<Steps>=ref(Steps.COUNTRY)
const filterLabel=computed(()=>{
  return currentStep.value!=Steps.ROOMS?"Next":"Search"
})
function filterSearch(){
  if(currentStep.value!=Steps.ROOMS){
    currentStep.value++
  }
}
</script>

<style scoped></style>
