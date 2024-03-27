<template>
  <div>
    <Modal title="Aribnb Your Home" :isOpen="rentModal.isOpen" :actionLabel="isLastStep() ? 'Create' : 'Next'"
      :disabled="isSubmitting" @onClose="rentModal.onClose()" @onSubmit="isLastStep() ? onSubmit() : nextStep()" secondaryActionLabel="Back" @onSecondaryAction="backStep()">
      <template v-slot:body>
        <div class="flex flex-col gap-4">
          <Heading title="Rent Modal" subtitle="Login to your account!" />

          <Input label="Email" name="email" placeholder="Enter your email" />
          <Input type="password" label="Password" name="password" placeholder="Enter your password" />

        </div>
      </template>
      <template v-slot:footer>
        <p class="mt-4 text-center font-light">Don't have an account? <span
            @click="rentModal.onClose(); registerModal.onOpen()"
            class="text-sky-500 cursor-pointer hover:text-sky-700">Register</span></p>
      </template>
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
import { type Ref, ref } from "vue"
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
const currentStep: Ref<Steps> = ref(Steps.CATEGORY)
const isLastStep = (): bool => {
  return currentStep.value+1 == Steps.END - Steps.START 
}
const schema = z.object({
  email: z.string().email().min(1),
  password: z.string().min(1)
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
const { handleSubmit, isSubmitting } = useForm({ validationSchema: validationSchema })
const rentModal = useRentModalStore()

const onSubmit = handleSubmit(async (values) => {
  console.log("Submit Handle")
})
</script>

<style scoped></style>
