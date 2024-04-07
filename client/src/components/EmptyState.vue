<template>
  <div class="h-[80vh] flex flex-col gap-2 justify-center items-center ">
    <Heading :title="title" :subtitle="subtitle" :center="true" />
    <div v-if="showReset" class="w-48 mt-4">
      <Button label="Remove all filters" outline @onClick="router.push('/')" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Heading from "@/components/Heading.vue"
import Button from "@/components/Button.vue"
import { useRouter } from 'vue-router';

withDefaults(defineProps<{
  title?: string
  subtitle?: string
  showReset?: boolean
}>(), {
  title: "Not exact match",
  subtitle: "Try remove or reset some of your filters",
  showReset: true
})
const router = useRouter()
</script>
