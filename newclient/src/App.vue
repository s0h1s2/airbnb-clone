<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import RegisterModal from './components/modal/RegisterModal.vue'
import LoginModal from './components/modal/LoginModal.vue'
</script>

<template>
  <header>
    <Navbar />
  </header>
  <RegisterModal />
  <LoginModal />
  <main class="pt-24 pb-20">
    <RouterView />
  </main>
</template>
