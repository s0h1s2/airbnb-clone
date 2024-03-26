<script setup lang="ts">

import MenuItem from './MenuItem.vue'
import { useRegisterModalStore } from "@/stores/registerModal"
import { useLoginModalStore } from "@/stores/loginModal"
import { useRentModalStore } from "@/stores/rentModalStore"
import { useUserStore } from "@/stores/userStore"
import { ref, type Ref } from "vue"
import { useRouter } from 'vue-router'
const registerModal = useRegisterModalStore()
const loginModal = useLoginModalStore()
const rentModal= useRentModalStore()

const userStore = useUserStore()
const router=useRouter()
const isOpen: Ref<boolean> = ref(false)
function toggleMenu() {
  isOpen.value = !isOpen.value
}
function openRentModal(){
  if (!userStore.isAuth) {
      loginModal.onOpen()
      return
  } 
  rentModal.onOpen()
}
</script>


<template>
  <div class="relative">

    <div class="flex flex-row items-center gap-3">
      <div @click="openRentModal" class="hidden 
        md:block 
        text-sm 
        font-semibold 
        py-3 
        px-4 
        rounded-full 
        hover:bg-neutral-100 
        transition 
        cursor-pointer">
        Airbnb your home
      </div>
      <div @click="toggleMenu" class="p-4 
        md:py-1 
        md:px-2 
        border-[1px] 
        border-neutral-200 
        flex 
        flex-row 
        items-center 
        gap-3 
        rounded-full 
        cursor-pointer 
        hover:shadow-md 
        transition">
        <v-icon name="pr-bars" />
        <div class="hidden md:block">
          <v-icon name="fa-user-alt" />
        </div>
      </div>
      <div v-if="isOpen"
        class="absolute rounded-xl shadow-md w-[40vw] md:w-3/4 bg-white overflow-hidden right-0 top-12 text-sm">
        <div class="flex flex-col cursor-pointer">
          <template v-if="!userStore.user">
            <MenuItem label="Sign in" @onClick="registerModal.onClose(); loginModal.onOpen()" />
            <MenuItem label="Sign up" @onClick="loginModal.onClose(); registerModal.onOpen()" />
          </template>
          <template v-if="userStore?.user">
            <MenuItem label="My Trips" @onClick="router.push({name:'trips'})" />
            <MenuItem label="My Favorites" @onClick="router.push({name:'favorites'})" />
            <MenuItem label="My Reservations" @onClick="router.push({name:'reservations'})" />
            <MenuItem label="My Properties" @onClick="router.push({name:'properties'})" />
            <hr/>
            <MenuItem label="Signout" @onClick="console.log('Signout')" />
          </template>
        </div>
      </div>
    </div>
  </div>
</template>


<style scoped></style>
