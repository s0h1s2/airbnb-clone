<template>
        <l-map ref="map" v-model:zoom="zoom" :center="[lat,lan]">
            <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
                name="OpenStreetMap"></l-tile-layer>
          <l-marker :lat-lng="coords" draggable >
            <l-popup>Property located here</l-popup>
          </l-marker>
        </l-map>
</template>

<script setup lang="ts">
import { ref, watch } from "vue"
import "leaflet/dist/leaflet.css";

const zoom = ref(2)
import { LMap, LTileLayer,LMarker,LPopup } from "@vue-leaflet/vue-leaflet";
const props=defineProps<{
    lat: number,
    lan: number 
}>()

const coords=ref([props.lat,props.lan] as L.LatLngExpression)
watch(()=>[props.lan,props.lat],()=>{
  coords.value=[props.lat,props.lan] as L.LatLngExpression
})

</script>

<style scoped></style>
