import { defineStore } from "pinia";
import { ref, type Ref } from "vue";

export function generateModalStore(name:string){
    return defineStore(name, () => {
        const isOpen: Ref<boolean> = ref(false)
        function onOpen() {
          isOpen.value = true;
        }
        function onClose() {
          isOpen.value = false;
        }
        return { isOpen, onOpen, onClose }
      })
}
