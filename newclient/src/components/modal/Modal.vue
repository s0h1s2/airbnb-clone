<template>
    <div v-if="showModal" class="flex 
    items-center 
    justify-center 
    overflow-x-hidden 
    overflow-y-auto 
    inset-0 
    fixed
    z-50 
    outline-none 
    focus:outline-none
    bg-neutral-800/70
    ">
    <div class="
    relative 
    w-full 
    md:w-4/6 
    lg:w-3/6 
    xl:w-2/5 
    my-6 
    mx-auto 
    h-full 
    lg:h-auto 
    md:h-auto">
    <!-- Content here -->
    <div class="translate duration-300 h-full" :class="[showModal?'translate-y-0 opacity-100':'translate-y-full opacity-0']">
        <div class="translate h-full lg:h-auto md:h-auto border-0 rounded-lg shadow-lg relative flex flex-col w-full outline-none focus:outline-none  bg-white">
            <!-- Header -->
            <div class="flex items-center p-6 rounded-t justify-center relative border-b-[1px] ">
                <button @click="handleClose" class="p-1 border-0 hover:opacity-75 transition absolute left-9">
                    <v-icon name="md-close" scale="1.2"></v-icon>
                </button>
                <div class="text-lg font-semibold">
                    {{title}}
                </div>
            </div>
            <!-- Body -->
            <div class="relative p-6 flex-auto">
                <slot name="body"></slot>
            </div>
            <!-- Footer -->
            <div class="flex flex-col gap-2 p-6">
                <div class="flex flex-row items-center gap-4 w-full">
                    <Button :disabled="disabled" :on-click="handleSubmit" :label="actionLabel"/>
                    <Button v-if="secondaryAction && secondaryActionLabel" :disabled="disabled" :on-click="()=>secondaryAction" :label="secondaryActionLabel"/>
                    <slot name="footer"></slot>
                </div>
            </div>
        </div>
    </div>
    </div>
    </div>
</template>

<script setup lang="ts">
import { watch,watchEffect, ref, type Ref } from 'vue';
import Button from "../Button.vue"

interface Props {
    isOpen?: boolean
    title?: string
    actionLabel: string
    secondaryActionLabel?: string
    secondaryAction?: () => void
    disabled: boolean
    onClose: () => void
    onSubmit: () => void
}
const props = defineProps<Props>()
const showModal: Ref<boolean> = ref(props.isOpen)
watchEffect(() => {
    showModal.value = props.isOpen
});
watch(()=>[props.disabled, props.onClose], () => {
    if (!props.disabled) {
        return
    }
    showModal.value = false;
    setTimeout(() => {
        props.onClose()
    }, 300);
});
watch([props.onSubmit, props.disabled],function handleSubmit () {
    if (props.disabled) {
        return
    }
    props.onSubmit()
});

watch([props.secondaryAction, props.disabled], () => {
    if (props.disabled || !props.secondaryAction) {
        return
    }
    props.secondaryAction()
});
function handleClose(){
    showModal.value=false
}
</script>

<style scoped></style>
