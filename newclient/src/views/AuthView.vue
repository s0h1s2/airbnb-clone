<template>
  <h1 class="flex h-screen items-center justify-center text-center">
    Authentication Require to visit
  </h1>
</template>

<script setup lang="ts">
import RegisterModal from "@/components/modal/RegisterModal.vue"
import LoginModal from "@/components/modal/LoginModal.vue"
import { useLoginModalStore } from "@/stores/loginModal"
import { onMounted } from "vue"
const loginModal = useLoginModalStore()

onMounted(() => {
  loginModal.onOpen()
})

</script>

<style scoped></style>
