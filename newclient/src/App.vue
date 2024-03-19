<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Modal from './components/modal/Modal.vue'

</script>

<template>
  <Navbar />
  <Modal actionLabel="Submit" title="Login" :isOpen="true" />
  <RouterView />
</template>
