<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Modal from './components/modal/Modal.vue'
import RegisterModal from './components/modal/RegisterModal.vue'

</script>

<template>
  <Navbar />

  <RegisterModal />
  <RouterView />
</template>
