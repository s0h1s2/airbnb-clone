<template>
  <div>
    <Modal :isOpen="registerModalStore.isOpen" actionLabel="Continue" :disabled="isSubmitting"
      @onClose="registerModalStore.onClose()" @onSubmit="onSubmit" >
      <template v-slot:body>
        <div class="flex flex-col gap-4">
          <Heading title="Welcome to an airbnb" subtitle="Create an account!" />
          <Input label="Email" name="email" placeholder="Enter your email" />
          <Input label="Name" name="name" placeholder="Enter your name" />
          <Input type="password" label="Password" name="password" placeholder="Enter your password" />
        </div>

      </template>
      <template v-slot:footer>
        <p class="mt-4 font-light text-center">Already have an account? <span @click="registerModalStore.onClose();loginModal.onOpen()" class="text-sky-500 cursor-pointer hover:text-sky-700">Login</span></p>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue"
import Heading from "../Heading.vue"
import Input from "../inputs/Input.vue"
import { z } from "zod"
import { client } from "@/lib/client"
import { useRegisterModalStore } from "@/stores/registerModal";
import { toTypedSchema } from "@vee-validate/zod"
import { useForm } from "vee-validate"
import { useToast } from "vue-toastification"
import { useLoginModalStore } from "@/stores/loginModal"

const schema = z.object({
  email: z.string().email().min(1),
  name: z.string().min(1),
  password: z.string().min(1)
})


const validationSchema = toTypedSchema(schema)
const { handleSubmit,isSubmitting,setErrors } = useForm({ validationSchema: validationSchema })
const registerModalStore = useRegisterModalStore()
const loginModal=useLoginModalStore()
const toast = useToast()

const onSubmit = handleSubmit((values) => {
  client.post("/users", values).then((r) => {
    toast.success("User registered successfuly")
    registerModalStore.onClose()
  }).catch((e) => {
    setErrors(e.response.data.errors)
    toast.error("Something went wrong.")
  })
})
</script>

<style scoped></style>
