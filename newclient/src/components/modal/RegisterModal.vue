<template>
  <div>
      <Modal :isOpen="registerModalStore.isOpen" actionLabel="Continue" :disabled="false"
        @onClose="registerModalStore.onClose()" @onSubmit="onSubmit">
        <template v-slot:body>
          <div class="flex flex-col gap-4">
            <Heading title="Welcome to an airbnb" subtitle="Create an account!" />
            <Input label="Email" name="email" placeholder="Enter your email" />
            <Input label="Name" name="name" placeholder="Enter your name" />
            <Input type="password" label="Password" name="password" placeholder="Enter your password" />
          </div>

        </template>
      </Modal>
  </div>
</template>

<script setup lang="ts">
import Modal from "./Modal.vue"
import Heading from "../Heading.vue"
import Input from "../inputs/Input.vue"
import {z}from "zod"
import { useRegisterModal } from "@/stores/registerModal";
import { toTypedSchema } from "@vee-validate/zod"
import { useForm } from "vee-validate"
import { useToast } from "vue-toastification"
const schema=z.object({
    email:z.string().email().min(1),
    name:z.string().min(1),
    password:z.string().min(1)
})


const validationSchema = toTypedSchema(schema)
const { handleSubmit }=useForm({validationSchema:validationSchema})
const registerModalStore = useRegisterModal()
const toast=useToast()

const onSubmit=handleSubmit((values)=>{
  console.log(values)
  toast.success("In the first part of the journey")
})
</script>

<style scoped></style>
