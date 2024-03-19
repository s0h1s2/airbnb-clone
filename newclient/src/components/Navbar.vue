<script setup lang="ts">
import Container from "./Container.vue"
import Logo from "./navbar/Logo.vue"
import Search from "./navbar/Search.vue"
import UserMenu from "./navbar/UserMenu.vue"
</script>
<template>
  <div class="fixed w-full bg-white shadow-sm">
    <div class="py-4 border-b-[1px]">
      <Container>
        <div class="flex flex-row items-center justify-between gap-3 md:gap-0">
          <Logo />
          <Search />
          <UserMenu />
        </div>
      </Container>
    </div>
  </div>
</template>

<script lang="ts">

</script>

<style scoped></style>
