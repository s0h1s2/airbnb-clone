<template>
  <Container>
    <div class="pt-4 flex flex-row items-center justify-between overflow-x-auto">
      <CategoryBox v-for="category in categories" :label="category.label" :iconName="category.iconName" :selected="false" />
    </div>
  </Container>
</template>

<script setup lang="ts">
import { ref } from "vue"
import Container from "../Container.vue"
import CategoryBox from "../CategoryBox.vue"
import { CATEGORIES } from "@/constants/categories"

const categories = ref(CATEGORIES)

</script>

<style scoped></style>
