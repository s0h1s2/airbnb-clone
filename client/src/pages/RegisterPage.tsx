import { SubmitHandler, useForm } from "react-hook-form"
import Input from "../components/Input"
import { Link } from "react-router-dom"
import client from "../lib/client"
type LoginForm = {
  name: string,
  email: string,
  password: string
}
const RegisterPage = () => {
  const { register, handleSubmit } = useForm<LoginForm>()
  const onSubmit: SubmitHandler<LoginForm> = async (data) => {
    client.post("/users", data,).then((response) => console.log(response.data)).catch((e) => console.log("ERROR" + e))
  }

  return (
    <div className="mt-4 grow flex items-center justify-around">
      <div className="mb-32">
        <h1 className="text-4xl text-center mb-4">
          Login
        </h1>
        <form className="max-w-md mx-auto" onSubmit={handleSubmit(onSubmit)}>
          <Input type="text" placeholder="Enter your name" field={register("name", { required: true })} />
          <Input type="email" placeholder="Enter your email" field={register("email", { required: true })} />
          <Input type="password" placeholder="Enter your password" field={register("password", { required: true })} />
          <div className="text-center py-2 text-gray-500">
            Already a member? <Link className="underline text-black" to="/login">Login</Link>
          </div>
          <button type="submit" className="w-full p-2 text-white bg-primary rounded-2xl">Register</button>
        </form>
      </div>
    </div>
  )
}
export default RegisterPage 
