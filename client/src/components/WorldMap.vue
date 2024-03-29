<template>
    <div style="height:400px; width:400px">
        <l-map ref="map" v-model:zoom="zoom" :center="[lat || 47.41322, lan || -1.219482]">
            <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base"
                name="OpenStreetMap"></l-tile-layer>
        </l-map>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import "leaflet/dist/leaflet.css";
import markerIcon2x from "leaflet/dist/images/marker-icon-2x.png"
import markerIcon from "leaflet/dist/images/marker-icon.png"
import markerShadow from "leaflet/dist/images/marker-shadow.png"
const zoom = ref(2)
import { LMap, LTileLayer } from "@vue-leaflet/vue-leaflet";
defineProps<{
    lat: number | null,
    lan: number | null
}>()
</script>

<style scoped></style>
