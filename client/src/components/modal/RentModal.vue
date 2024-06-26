<template>
  <div>
    <Modal title="Aribnb Your Home" :isOpen="rentModal.isOpen" :actionLabel="isLastStep() ? 'Create' : 'Next'"
      :disabled="isSubmitting" @onClose="rentModal.onClose()" @onSubmit="isLastStep
        () ? onSubmit() : nextStep()" secondaryActionLabel="Back" @onSecondaryAction="backStep()">
      <template v-slot:body>
        <div class="flex flex-col gap-4">
          <div v-show="currentStep == Steps.CATEGORY">
            <Heading title="Which of these best describe your place" subtitle="Pick a category" />
            <div class="grid grid-cols-1 md:grid-cols-2 gap-3 max-h-[50vh] overflow-y-auto">
              <div class="col-span-1" v-for="category in categories">
                <CategoryInput :label="category.label" :iconName="category.iconName"
                  :selected="selectedCategory == category.label"
                  @onClick="(value) => { selectedCategory = value; setFieldValue('category', value) }" />
              </div>
            </div>
          </div>
          <div v-if="currentStep == Steps.LOCATION">
            <div class="flex flex-col gap-8">
              <Heading title="Where is your place located?" subtitle="Help guests find you!" />
              <CountrySelect v-model="selectedCountry" />
              <div class="w-auto h-[40vh]">
                <WorldMap :lat="selectedCountry?.latlang[0] || 40.737" :lan="selectedCountry?.latlang[1] || -73.923"
                  @onCoordChange="(coord) => locationCoord = coord as L.LatLng" />
              </div>
            </div>
          </div>
          <div v-show="currentStep == Steps.INFO">
            <div class="flex flex-col gap-8">
              <Heading title="Share some basics about your place" subtitle="What amenities do you have" />
              <CounterInput title="Guests" subtitle="How many guests do you allow?" name="guestCount" />
              <CounterInput title="Rooms" subtitle="How many rooms do you have?" name="roomCount" />
              <CounterInput title="Bathrooms" subtitle="How many bathrooms do you have?" name="bathroomCount" />
            </div>
          </div>
          <div v-show="currentStep == Steps.IMAGES">
            <div class="flex flex-col gap-8">
              <Heading title="Add photo of your place" subtitle="Show your place what look like!" />
              <ImageUpload v-model="imageSrc" />
            </div>
          </div>
          <div v-show="currentStep == Steps.DESCRIPTION">
            <div class="flex flex-col gap-8">
              <Heading title="How you describe your place?" subtitle="Short and sweet works best!" />
              <Input name="title" label="Title" required />
              <hr />
              <Input name="description" label="Description" required />
            </div>
          </div>
          <div v-show="currentStep == Steps.PRICE">
            <div class="flex flex-col gap-8">
              <Heading title="Now,set your price" subtitle="How much you charge for a night?" />
              <Input name="price" label="Price" type="number" formatPrice required />
            </div>
          </div>

        </div>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue"
import Heading from "../Heading.vue"
import Input from "../inputs/Input.vue"
import { useForm } from "vee-validate"
import { useToast } from "vue-toastification"
import { useRentModalStore } from "@/stores/rentModalStore";
import { type Ref, ref, readonly, watch } from "vue"
import { CATEGORIES } from "@/constants/categories"
import CategoryInput from "@/components/inputs/CategoryInput.vue"
import CountrySelect from "@/components/inputs/CountrySelect.vue"
import CounterInput from "@/components/CounterInput.vue"
import ImageUpload from "@/components/inputs/ImageUpload.vue"
import type { Country } from "@/types/country"
import { client } from "@/lib/client"
import WorldMap from "@/components/WorldMap.vue"
enum Steps {
  START,
  CATEGORY,
  LOCATION,
  INFO,
  IMAGES,
  DESCRIPTION,
  PRICE,
  END
}

const toast = useToast()
const selectedCategory: Ref<string> = ref(CATEGORIES[0].label)
const currentStep: Ref<Steps> = ref(Steps.CATEGORY)
const imageSrc: Ref<string | null> = ref(null)
const rentModal = useRentModalStore()
const categories = readonly(CATEGORIES)
const selectedCountry = defineModel<Country>()
const locationCoord: Ref<{lat:number,lng:number}> = ref({lat:40.737, lng:-73.923})
watch(imageSrc, () => {
  setFieldValue('imageSrc', imageSrc?.value || "")
})

watch(selectedCountry, () => {
  setFieldValue('country', selectedCountry.value?.value || "")
})
watch(locationCoord, () => {
  setFieldValue("location", { lat: locationCoord.value.lat, lng: locationCoord.value.lng})
})

type FormData = {
  category: string
  title: string
  description: string
  guestCount: number
  roomCount: number
  bathroomCount: number
  imageSrc: string
  price: number
  location: { lat: number, lng: number }
  country: string
}
const { handleSubmit, isSubmitting, setFieldValue } = useForm<FormData>({
  initialValues: {
    category: CATEGORIES[0].label,
    location: { lat: 0, lng: 0 },
    country: ""
  }

})

const isLastStep = (): boolean => {
  return currentStep.value + 1 == Steps.END - Steps.START
}

const nextStep = () => {
  if (currentStep.value < Steps.END)
    currentStep.value++
}


const backStep = () => {
  if (currentStep.value > Steps.START + 1)
    currentStep.value--;

}

const onSubmit = handleSubmit(async (data) => {
  await client.post("/listing", data).then(() => {
    toast.success("Listing created successfully")
    rentModal.onClose()
  }).catch(() => {
    toast.error("Unable to create listing. Please try again")
  })
})
</script>
