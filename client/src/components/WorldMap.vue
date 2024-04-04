<template>
  <l-map :use-global-leaflet="false" ref="map" v-model:zoom="zoom" :center="[lat, lan]">
    <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
      name="OpenStreetMap"></l-tile-layer>
    <l-marker :lat-lng="coords" draggable>
      <l-popup>Property located here</l-popup>
    </l-marker>
  </l-map>
</template>

<script setup lang="ts">

import "leaflet/dist/leaflet.css";
import { ref, watch } from "vue"
import { LMap, LTileLayer, LMarker, LPopup } from "@vue-leaflet/vue-leaflet";

const zoom = ref(2)

const props = defineProps<{
  lat: number,
  lan: number
}>()
const emit = defineEmits<{ onCoordChange: [L.LatLngExpression] }>()
const coords = ref([props.lat, props.lan] as L.LatLngExpression)
watch(() => [props.lan, props.lat], () => {
  coords.value = [props.lat, props.lan] as L.LatLngExpression
})
watch(coords, () => {
  emit('onCoordChange', coords.value)
})


</script>

<style scoped></style>
