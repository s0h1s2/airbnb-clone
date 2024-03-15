import { SubmitHandler, useForm } from "react-hook-form"
import Input from "../components/Input"
import { Link, useNavigate } from "react-router-dom"
import client, { ResponseResult } from "../lib/client"
import { toast } from "react-toastify"
import { useBoundStore } from "../state/boundState"
import clsx from "clsx"
import { FIELD_REQUIRED } from "../constants/messages"
type LoginForm = {
  email: string,
  password: string
}
type UserResponseSuccess = ResponseResult & {
  data: { email: string, name: string }
}
const LoginPage = () => {
  const { register, handleSubmit, formState: { isSubmitting, errors } } = useForm<LoginForm>()
  const navigate = useNavigate()
  const setUser = useBoundStore((state) => state.setUser)
  const onSubmit: SubmitHandler<LoginForm> = async (data) => {
    return client.post<UserResponseSuccess>("/users/auth", data, { withCredentials: true })
      .then((r) => {
        toast.success("Login was successful")
        setUser({ name: r.data.data.name, email: r.data.data.email })
        navigate("/")
      })
      .catch(() => {
        toast.error("Invalid credentials")
      })

  }

  return (
    <div className="mt-4 grow flex items-center justify-around">
      <div className="mb-32">
        <h1 className="text-4xl text-center mb-4">
          Login
        </h1>
        <form className="max-w-md mx-auto" onSubmit={handleSubmit(onSubmit)}>
          <Input type="email" placeholder="Enter your email" field={register("email", { required: true })} />
          {errors.email && <span className="text-red-700">{FIELD_REQUIRED}</span>}
          <Input type="password" placeholder="Enter your password" field={register("password", { required: true })} />
          {errors.password && <span className="text-red-700">{FIELD_REQUIRED}</span>}
          <div className="text-center py-2 text-gray-500">
            Don't you have an account yet? <Link className="underline text-black" to="/register">Register Now</Link>
          </div>
          <button disabled={isSubmitting} type="submit" className={clsx("w-full p-2 text-white bg-primary rounded-2xl", isSubmitting && "opacity-75")}>Login</button>
        </form>
      </div>
    </div>
  )
}

export default LoginPage
