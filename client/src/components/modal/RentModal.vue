<template>
  <div>
    <Modal title="Aribnb Your Home" :isOpen="rentModal.isOpen" :actionLabel="isLastStep() ? 'Create' : 'Next'"
      :disabled="isSubmitting" @onClose="rentModal.onClose()" @onSubmit="isLastStep() ? onSubmit(): nextStep()"
      secondaryActionLabel="Back" @onSecondaryAction="backStep()">
      <template v-slot:body>
        <div class="flex flex-col gap-4">
          <template v-if="currentStep==Steps.CATEGORY">
          <Heading title="Which of these best describe your place" subtitle="Pick a category" />
          <div  class="grid grid-cols-1 md:grid-cols-2 gap-3 max-h-[50vh] overflow-y-auto">
              <div  class="col-span-1" v-for="category in categories">
                <CategoryInput :label="category.label" :iconName="category.iconName" :selected="selectedCategory==category.label" @onClick="(value)=>{selectedCategory=value;setFieldValue('category',value)}"/>
              </div>
          </div>
          </template>
          <template v-if="currentStep==Steps.LOCATION">
          <div class="flex flex-col gap-8">
            <Heading title="Where is your place located?" subtitle="Help guests find you!" />
            <CountrySelect @onChange="(country)=>console.log(country)"/>
          </div>
          </template>
        </div>
      </template>
      <template v-slot:footer></template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue"
import Heading from "../Heading.vue"
import Input from "../inputs/Input.vue"
import { z } from "zod"
import { toTypedSchema } from "@vee-validate/zod"
import { useForm } from "vee-validate"
import { useToast } from "vue-toastification"
import { useUserStore } from "@/stores/userStore"
import { useRentModalStore } from "@/stores/rentModalStore";
import { type Ref, ref, readonly } from "vue"
import { CATEGORIES } from "@/constants/categories"
import CategoryInput from "@/components/inputs/categoryinput.vue"
import CountrySelect from "@/components/inputs/countryselect.vue"
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
const selectedCategory:Ref<string>=ref(CATEGORIES[0].label)
const currentStep: Ref<Steps> = ref(Steps.CATEGORY)
const isLastStep = (): boolean => {
  return currentStep.value + 1 == Steps.END - Steps.START
}
const schema = z.object({
  category:z.string().min(1),
  title:z.string().min(1),
  description:z.string().min(1),
  guestCount:z.number().min(1),
  roomCount:z.number(),
  bathroomCount:z.number(),
  imageSrc:z.string(),
  price:z.number(),
  location:z.string().min(1)

})

function nextStep() {
  if (currentStep.value < Steps.END) {
    currentStep.value++
  }
}

function backStep() {
  if (currentStep.value > Steps.START) {
    currentStep.value--;
  }
}

const validationSchema = toTypedSchema(schema)
const { handleSubmit, isSubmitting,setFieldValue} = useForm({validationSchema})

const rentModal = useRentModalStore()
const categories = readonly(CATEGORIES)
const onSubmit = handleSubmit(async (values) => {
  console.log(values)
  console.log("Submit Handle")
})
</script>

<style scoped></style>
