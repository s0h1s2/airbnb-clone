<template>
    <div @click="emit('onClick',label)" class="rounded-xl border-2 p-4 flex flex-col gap-3 hover:border-black transition cursor-pointer" :class="[selected?'border-black':'border-neutral-200']">
        <v-icon :name="iconName"/>
        <div class="font-semibold">
            {{ label }}
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps<{label:string,iconName:string,selected?:boolean}>()
const emit=defineEmits<{onClick:[string]}>()
</script>

<style scoped>

</style>