<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
import Modal from './components/modal/Modal.vue'

</script>

<template>
  <Navbar />
  <Modal :disabled="false" actionLabel="Submit" secondaryActionLabel="Login" title="Login" :isOpen="true" @on-submit="()=>console.log('Submit')"  @on-close="()=>console.log('Close') " @on-secondary-action="()=>console.log('Whatever?')"/>

  <RouterView />
</template>
