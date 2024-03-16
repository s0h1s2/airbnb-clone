import { useForm } from "react-hook-form"
import Input, { InputProps } from "../components/Input"
import React from "react"
import { PERKS } from "../constants/perks"

type NewPlaceForm = {
  title: string,
  address: string,
  description: string,
  extraInfo: string,
  checkIn: string,
  checkOut: string,
  maxGuests: number,
  perks: string[]
}
const PERKS_NAME_FIELD = "perks[]"
type NewPlaceHeadProps = { title: string, description?: string }
type Props = InputProps & NewPlaceHeadProps
const NewPlaceInputHead = ({ title, description }: NewPlaceHeadProps) => {
  return (
    <>
      <h2 className="text-2xl mt-4">{title}</h2>
      {description != null && <p className="text-gray-500 text-sm">{description}</p>}
    </>
  )

}
const CheckBox = ({ children, name }: { children: React.ReactNode, name: string }) => {
  return (
    <label className="border p-4 flex flex-row justify-center items-center rounded-xl gap-2 ">
      <input name={PERKS_NAME_FIELD} value={name} type="checkbox" />
      <span>
        {children}
      </span>
      <span>
        {name}
      </span>
    </label>

  )
}
const NewPlaceInput = ({ title, description, field, placeholder, type }: Props) => {
  return (
    <>
      <NewPlaceInputHead title={title} description={description} />
      <Input type={type} placeholder={placeholder} field={field} />
    </>
  )

}
const NewPlace = () => {
  const onSubmit = async (data: NewPlaceForm) => {
    console.log(data)
  }
  const { register, handleSubmit } = useForm<NewPlaceForm>()

  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <NewPlaceInput type={"text"} placeholder={"e.g. 'My Lovely Apt.'"} field={register("title", { required: true })} title={"Title"} description={"Your title for catch"} />
        <NewPlaceInput type={"text"} placeholder={"e.g. 'Greenland'"} field={register("address", { required: true })} title={"Address"} description={"Your Address to this place"} />
        <h2 className="text-2xl mt-4">Photos</h2>
        <p className="text-gray-500 text-sm">More=Better</p>
        <div className="grid grid-cols-3 lg:grid-cols-6 md:grid-cols-4 mt-2 ">
          <button className="border bg-transparent rounded-2xl p-8 text-2xl text-gray-500">+</button>
        </div>
        <NewPlaceInput type={"text"} placeholder={"e.g. 'Near Beach'"} field={register("description", { required: true })} title={"Description"} description={"Description of your place"} />
        <NewPlaceInputHead title="Perks" description="Select the following perks" />
        <div className="grid gap-2 grid-cols-2 md:grid-cols-4 lg:grid-cols-6">
          {PERKS.map(({ icon: Icon, name }) => (
            <CheckBox children={<Icon className="w-4 h-4" />} name={name} />
          ))}
        </div>
        <NewPlaceInput type="text" title="Extra info" description="More info like house rules,...etc" field={register("extraInfo")} placeholder="e.g. 'No smoking'" />
        <NewPlaceInputHead description="Add your checkout and checkin times" title="Checkout & Checkin times" />
        <div className="grid gap-2 sm:grid-cols-3">
          <div>
            <h3 className="mt-2 -mb1">Check in time</h3>
            <Input type={"text"} placeholder={"14:00"} field={register("checkIn")} />
          </div>
          <div>
            <h3 className="mt-2 -mb1">Check out time</h3>
            <Input type={"text"} placeholder={"14:00"} field={register("checkIn")} />
          </div>
          <div>
            <h3 className="mt-2 -mb1">Max number of guests</h3>
            <Input type={"number"} placeholder={"5"} field={register("checkIn")} />
          </div>

        </div>
        <button className="mt-4 primary">Create</button>
      </form>
    </div>
  )
}

export default NewPlace
