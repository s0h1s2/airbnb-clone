<template>
  <div>

  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
const scriptLoaded = ref(false)

onMounted(() => {
  const script = document.createElement('script')
  script.setAttribute('async', '')
  script.setAttribute('id', 'uw')
  script.src = 'https://upload-widget.cloudinary.com/global/all.js'
  script.addEventListener('load', () => (scriptLoaded.value = true))

  script.onerror = () => {
    console.error('Failed to load third-party script.')
  }
  document.head.appendChild(script)
  // window!.cloudinary!.openUploadWidget()
})
</script>

<style scoped></style>
