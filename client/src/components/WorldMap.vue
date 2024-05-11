  <template>
    <l-map :use-global-leaflet="false" style="z-index: 0;" ref="map" v-model="zoom" v-model:zoom="zoom"
      :center="[lat, lan]">
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
import { ref, watch, type Ref } from "vue"
import { LMap, LTileLayer, LMarker, LPopup } from "@vue-leaflet/vue-leaflet";
import L from 'leaflet';
globalThis.L = L

const map: Ref<L.Map | null> = ref(null)

const zoom = ref(4)

defineExpose({ map })

const emit = defineEmits<{ onCoordChange: [L.LatLngExpression] }>()
const coords = ref([props.lat, props.lan] as L.LatLngExpression)
watch(() => [props.lan, props.lat], () => {
  coords.value = [props.lat, props.lan] as L.LatLngExpression
})
watch(coords, (newCoords) => {
  emit('onCoordChange', newCoords)
})
</script>

<style scoped></style>
