<template>
  <div>
    <Modal :isOpen="registerModalStore.isOpen" actionLabel="Register" :disabled="false" @onClose="registerModalStore.onClose()">
      <template v-slot:body>
          <h1>Hello</h1>
      </template>
    </Modal>
    
  </div>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue"
import { useRegisterModal } from "@/stores/registerModal";
const registerModalStore=useRegisterModal()

</script>

<style scoped></style>
