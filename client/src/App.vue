<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import RegisterModal from './components/modal/RegisterModal.vue'
import LoginModal from './components/modal/LoginModal.vue'
import RentModal from './components/modal/RentModal.vue'
import { useRentModalStore } from './stores/rentModalStore'

const rentModal = useRentModalStore()
</script>

<template>
  <header>
    <Navbar />
  </header>
  <RegisterModal />
  <LoginModal />
  <!-- Make sure component re render after close.-->
  <RentModal v-if="rentModal.isOpen" />
  <main class="pt-24 pb-20">
    <RouterView />
  </main>
</template>
