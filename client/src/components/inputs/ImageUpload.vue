<template>
  <div
    class="relative cursor-pointer hover:opacity-70 transition p-20 border-2 border-neutral-300 flex flex-col justify-center items-center gap-4 text-neutral-600"
    @click="openUploadWidget" :style="{ 'pointer-events': isDisabled ? 'none' : 'auto' }">
    <v-icon :scale="2" name="md-addphotoalternate-outlined"></v-icon>
    <div class="font-semibold text-lg">
      Click to upload
    </div>
    <div v-if="image != null" class="absolute overflow-hidden inset-0 w-full h-auto">
      <img :src="image" alt="Upload" class="w-full h-full object-cover object-center " />
    </div>

  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"

const image = defineModel<string | null>()

const scriptLoaded = ref(false)
const isDisabled = ref(false)
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

})
function processResults(error: any, result: any) {
  if (error || result.event == 'close') {
    isDisabled.value = false
    return
  }
  if (result && result.event == 'success') {
    image.value = result.info.secure_url
  }
}
function openUploadWidget() {
  if (!scriptLoaded) {
    return
  }
  isDisabled.value = true
  //@ts-ignore
  window?.cloudinary.openUploadWidget({
    cloudName: "dyek1w4dj",
    uploadPreset: "oassehvr",
    sources: ['local', 'url'],
    clientAllowedFormats: ['image'],
    resourceType: 'image',
    maxFiles: 1

  }, processResults)
}
</script>

<style scoped></style>
