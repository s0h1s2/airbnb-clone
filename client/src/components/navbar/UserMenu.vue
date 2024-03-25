<script setup lang="ts">

import MenuItem from './MenuItem.vue'
import { useRegisterModalStore } from "@/stores/registerModal"
import { useLoginModalStore } from "@/stores/loginModal"
import { useUserStore } from "@/stores/userStore"
import { ref, type Ref } from "vue"
const registerModal = useRegisterModalStore()
const loginModal = useLoginModalStore()
const userStore = useUserStore()

const isOpen: Ref<boolean> = ref(false)
function toggleMenu() {
  isOpen.value = !isOpen.value
}

</script>


<template>
  <div class="relative">

    <div class="flex flex-row items-center gap-3">
      <div class="hidden 
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
          <span class="font-sembold text-sm pl-1">
            {{ userStore?.user?.name }}
          </span>
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
            <MenuItem label="My Profile" @onClick="console.log('Show profile')" />
          </template>
        </div>
      </div>
    </div>
  </div>
</template>


<style scoped></style>
