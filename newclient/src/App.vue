<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import RegisterModal from './components/modal/RegisterModal.vue'
import LoginModal from './components/modal/LoginModal.vue'

</script>

<template>
  <Navbar />
  <RegisterModal />
  <LoginModal />
  <RouterView />
</template>
