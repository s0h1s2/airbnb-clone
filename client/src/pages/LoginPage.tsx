import { SubmitHandler, useForm } from "react-hook-form"
import Input from "../components/Input"
import { Link } from "react-router-dom"
type LoginForm = {
  email: string,
  password: string
}
const LoginPage = () => {
  const { register, handleSubmit } = useForm<LoginForm>()
  const onSubmit: SubmitHandler<LoginForm> = (data) => { console.log(data) }

  return (
    <div className="mt-4 grow flex items-center justify-around">
      <div className="mb-32">
        <h1 className="text-4xl text-center mb-4">
          Login
        </h1>
        <form className="max-w-md mx-auto" onSubmit={handleSubmit(onSubmit)}>
          <Input type="email" placeholder="Enter your email" field={register("email", { required: true })} />
          <Input type="password" placeholder="Enter your password" field={register("password", { required: true })} />
          <div className="text-center py-2 text-gray-500">
            Don't you have an account yet? <Link className="underline text-black" to="/register">Register Now</Link>

          </div>
          <button type="submit" className="w-full p-2 text-white bg-primary rounded-2xl">Login</button>
        </form>
      </div>
    </div>
  )
}

export default LoginPage
