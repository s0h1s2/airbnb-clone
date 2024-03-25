<template>
  <template v-if="route.name == 'home'">
    <Container>
      <div class="pt-4 flex flex-row items-center justify-between overflow-x-auto">
        <CategoryBox v-for="category in categories" :label="category.label" :iconName="category.iconName"
          :selected="false" />
      </div>
    </Container>

  </template>
</template>

<script setup lang="ts">
import { ref } from "vue"
import Container from "../Container.vue"
import CategoryBox from "../CategoryBox.vue"
import { CATEGORIES } from "@/constants/categories"
import { useRoute } from "vue-router"
const categories = ref(CATEGORIES)
const route = useRoute()

</script>

<style scoped></style>
