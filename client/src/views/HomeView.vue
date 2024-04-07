<script setup lang="ts">
import Container from "@/components/Container.vue"
import EmptyState from "@/components/EmptyState.vue"
import { ref, type Ref } from "vue";
const isEmpty: Ref<boolean> = ref(true)


</script>

<template>
  <EmptyState v-if="isEmpty" />
  <Container v-else>
    <div
      class="pt-24 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6 gap-8">
      <div>
        My Future listing
      </div>
    </div>
  </Container>
</template>
