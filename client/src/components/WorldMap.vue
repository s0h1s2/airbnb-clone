<template>
  <l-map ref="map" noWrap v-model="zoom" v-model:zoom="zoom" style="z-index: 0;" :center="[lat, lan]"
    @ready="onReady()">
    <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
      name="OpenStreetMap"></l-tile-layer>
    <l-marker :lat-lng="coords" @moveend="(v) => { emit('onCoordChange', v.target._latlng as L.LatLngExpression) }"
      draggable>
      <l-popup>Property located here</l-popup>
    </l-marker>
  </l-map>
</template>

<script setup lang="ts">
const props = defineProps<{
  lat: number
  lan: number
}>()

import { onMounted, ref, watch, nextTick } from "vue"
import { LMap, LTileLayer, LMarker, LPopup } from "@vue-leaflet/vue-leaflet";
import L from 'leaflet';
globalThis.L = L;
import "leaflet/dist/leaflet.css";
const map = ref(null)
const zoom = ref(4)
const emit = defineEmits<{ onCoordChange: [L.LatLngExpression] }>()
const coords = ref([props.lat, props.lan] as L.LatLngExpression)
watch(() => [props.lan, props.lat], () => {
  coords.value = [props.lat, props.lan] as L.LatLngExpression
})
watch(() => coords, (newCoords) => {
  emit('onCoordChange', newCoords.value)
})
function onReady() {
  console.log(map.value.leafletObject)
  map.value.leafletObject.invalidateSize()
}
</script>

<style scoped></style>
