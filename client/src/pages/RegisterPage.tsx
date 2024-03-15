import { SubmitHandler, useForm } from "react-hook-form"
import Input from "../components/Input"
import { Link, useNavigate } from "react-router-dom"
import client, { ResponseResult } from "../lib/client"
import { FIELD_REQUIRED } from "../constants/messages"
import clsx from "clsx"
import { toast } from "react-toastify"
import { useState } from "react"
import { AxiosError } from "axios"
type RegisterForm = {
  name: string,
  email: string,
  password: string
}
type RegisterFormError = ResponseResult & {
  errors: {
    name: string,
    email: string,
  }

}
const RegisterPage = () => {
  const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<RegisterForm>()
  const navigate = useNavigate()
  const [apiErrors, setApiErrors] = useState<RegisterFormError>()
  const onSubmit: SubmitHandler<RegisterForm> = async (data) => {
    return client.post("/users", data).then(() => {
      toast.success("User registered successfuly")
      navigate("/login")
    }).catch((e: AxiosError) => {
      console.log((e.response?.data as RegisterFormError).errors.email)
      setApiErrors((e.response?.data as RegisterFormError))
      toast.error("Something went wrong")
    })
  }

  return (
    <div className="mt-4 grow flex items-center justify-around">
      <div className="mb-32">
        <h1 className="text-4xl text-center mb-4">
          Register
        </h1>
        <form className="max-w-md mx-auto" onSubmit={handleSubmit(onSubmit)}>
          <Input type="text" placeholder="Enter your name" field={register("name", { required: true })} />
          {errors.name && <span className="text-red-800">{FIELD_REQUIRED}</span>}
          <Input type="email" placeholder="Enter your email" field={register("email", { required: true })} />
          {errors.email && <span className="text-red-800">{FIELD_REQUIRED}</span>}
          {apiErrors?.errors.email && <span className="text-red-800">{apiErrors.errors.email}</span>}
          <Input type="password" placeholder="Enter your password" field={register("password", { required: true })} />
          {errors.email && <span className="text-red-800">{FIELD_REQUIRED}</span>}
          <div className="text-center py-2 text-gray-500">
            Already a member? <Link className="underline text-black" to="/login">Login</Link>
          </div>
          <button disabled={isSubmitting} type="submit" className={clsx("w-full p-2 text-white bg-primary rounded-2xl", isSubmitting && "opacity-50")}>Register</button>
        </form>
      </div>
    </div>
  )
}
export default RegisterPage 
