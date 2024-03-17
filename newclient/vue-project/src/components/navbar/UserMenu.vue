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
      <div class="p-4 
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
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

</script>

<style scoped></style>
