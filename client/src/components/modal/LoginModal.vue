<template>
  <div>
    <Modal :isOpen="loginModal.isOpen" actionLabel="Continue" :disabled="isSubmitting" @onClose="loginModal.onClose()"
      @onSubmit="onSubmit">
      <template v-slot:body>
        <div class="flex flex-col gap-4">
          <Heading title="Welcome back!" subtitle="Login to your account!" />
          <Input label="Email" name="email" placeholder="Enter your email" />
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
import { z } from "zod"
import { client } from "@/lib/client"
import { useLoginModalStore } from "@/stores/loginModal";
import { toTypedSchema } from "@vee-validate/zod"
import { useForm } from "vee-validate"
import { useToast } from "vue-toastification"
import { useUserStore } from "@/stores/userStore"

const schema = z.object({
  email: z.string().email().min(1),
  password: z.string().min(1)
})


const validationSchema = toTypedSchema(schema)
const { handleSubmit, isSubmitting, setErrors } = useForm({ validationSchema: validationSchema })
const loginModal = useLoginModalStore()
const toast = useToast()
const userStore = useUserStore()

const onSubmit = handleSubmit(async (values) => {
  await userStore.login(values)
  if (userStore.isAuth) {
    toast.success("You are loggedin successfuly")
    loginModal.onClose()
  } else {
    toast.error("Invalid credentials")
  }
})
</script>

<style scoped></style>
