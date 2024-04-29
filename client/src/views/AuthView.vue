<template>
  <h1 class="flex h-screen items-center justify-center text-center">
    Authentication Require to visit {{ route.query?.redirect }}
  </h1>
</template>

<script setup lang="ts">
import { useLoginModalStore } from "@/stores/loginModal"
import { onMounted } from "vue"
import { storeToRefs } from "pinia"
import { useUserStore } from "@/stores/userStore";
import { useRouter, useRoute } from "vue-router";
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const { isAuth } = storeToRefs(userStore)
const loginModal = useLoginModalStore()
if (isAuth.value == true) {
  router.replace(route.query.redirect?.toString() || "/")
}

onMounted(() => {
  loginModal.onOpen()
})
</script>

<style scoped></style>
