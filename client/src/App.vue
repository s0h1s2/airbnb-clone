<script setup lang="ts">
import { RouterView } from 'vue-router'
import Navbar from '@/components/Navbar.vue'
import RegisterModal from '@/components/modal/RegisterModal.vue'
import LoginModal from '@/components/modal/LoginModal.vue'
import RentModal from '@/components/modal/RentModal.vue'
import FilterModal from '@/components/modal/FilterModal.vue'

import { useRentModalStore } from './stores/rentModalStore'
import { useFilterModalStore } from './stores/filterModalStore'

const rentModal = useRentModalStore()
const filterModal= useFilterModalStore()
</script>

<template>
  <header>
    <Navbar />
  </header>
  <RegisterModal />
  <LoginModal />
  <!-- Make sure component re render after close.-->
  <RentModal v-if="rentModal.isOpen" />
  <FilterModal v-if="filterModal.isOpen" />

  <main class="pt-24 pb-20">
    <RouterView />
  </main>
</template>
