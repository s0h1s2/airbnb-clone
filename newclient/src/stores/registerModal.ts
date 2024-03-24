import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

export const useRegisterModalStore = defineStore("registerModal", () => {
  const isOpen: Ref<boolean> = ref(false)
  function onOpen() {
    isOpen.value = true;
  }
  function onClose() {
    isOpen.value = false;
  }
  return { isOpen, onOpen, onClose }
})
