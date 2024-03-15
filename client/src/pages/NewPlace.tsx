import { useForm } from "react-hook-form"
import Input, { InputProps } from "../components/Input"
import React from "react"

type NewPlaceForm = {
  title: string,
  address: string,
  description: string
}
type NewPlaceHeadProps = { title: string, description: string }
type Props = InputProps & NewPlaceHeadProps
const NewPlaceInputHead = ({ title, description }: NewPlaceHeadProps) => {
  return (
    <>
      <h2 className="text-2xl mt-4">{title}</h2>
      <p className="text-gray-500 text-sm">{description}</p>
    </>
  )

}
const CheckBox = ({ children, name }: { children: React.ReactNode, name: string }) => {
  return (
    <label className="border p-4 flex rounded-xl gap-2 items-center">
      <input type="checkbox" />
      <span>
        {children}
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
  const onSubmit = async () => { }
  const { register, handleSubmit } = useForm<NewPlaceForm>()
  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <NewPlaceInput type={"text"} placeholder={"e.g. 'My Lovely Apt.'"} field={register("title", { required: true })} title={"Title"} description={"Your title for catch"} />
        <NewPlaceInput type={"text"} placeholder={"e.g. 'Greenland'"} field={register("address", { required: true })} title={"Description"} description={"Your Address to this place"} />
        <h2 className="text-2xl mt-4">Photos</h2>
        <p className="text-gray-500 text-sm">More=Better</p>
        <div className="grid grid-cols-3 lg:grid-cols-6 md:grid-cols-4 mt-2 ">
          <button className="border bg-transparent rounded-2xl p-8 text-2xl text-gray-500">+</button>
        </div>
        <NewPlaceInput type={"text"} placeholder={"e.g. 'Near Beach'"} field={register("description", { required: true })} title={"Description"} description={"Description of your place"} />
        <NewPlaceInputHead title="Perks" description="Select the following perks" />
        <div className="grid gap-2 grid-cols-2 md:grid-cols-4 lg:grid-cols-6">
          <CheckBox name="Wifi">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
              <path strokeLinecap="round" strokeLinejoin="round" d="M8.288 15.038a5.25 5.25 0 0 1 7.424 0M5.106 11.856c3.807-3.808 9.98-3.808 13.788 0M1.924 8.674c5.565-5.565 14.587-5.565 20.152 0M12.53 18.22l-.53.53-.53-.53a.75.75 0 0 1 1.06 0Z" />
            </svg>
          </CheckBox>
          <CheckBox name="Pet Allowed">
            <svg
              viewBox="0 0 512 512"
              fill="currentColor"
              height="1em"
              width="1em"
            >
              <path d="M288 192h17.1c22.1 38.3 63.5 64 110.9 64 11 0 21.8-1.4 32-4v228c0 17.7-14.3 32-32 32s-32-14.3-32-32V339.2L248 448h56c17.7 0 32 14.3 32 32s-14.3 32-32 32H160c-53 0-96-43-96-96V192.5c0-16.1-12-29.8-28-31.8l-7.9-1C10.5 157.6-1.9 141.6.2 124s18.2-30 35.7-27.8l7.9 1c48 6 84.1 46.8 84.1 95.3v85.3c34.4-51.7 93.2-85.8 160-85.8zm160 26.5c-10 3.5-20.8 5.5-32 5.5-28.4 0-54-12.4-71.6-32-3.7-4.1-7-8.5-9.9-13.2C325.3 164 320 146.6 320 128V10.7C320 4.8 324.7.1 330.6 0h.2c3.3 0 6.4 1.6 8.4 4.2v.1l12.8 17 27.2 36.3L384 64h64l4.8-6.4L480 21.3l12.8-17v-.1c2-2.6 5.1-4.2 8.4-4.2h.2c5.9.1 10.6 4.8 10.6 10.7V128c0 17.3-4.6 33.6-12.6 47.6-11.3 19.8-29.6 35.2-51.4 42.9zM400 128c0-8.8-7.2-16-16-16s-16 7.2-16 16 7.2 16 16 16 16-7.2 16-16zm48 16c8.8 0 16-7.2 16-16s-7.2-16-16-16-16 7.2-16 16 7.2 16 16 16z" />
            </svg>
          </CheckBox>
          <CheckBox name="Car Parking">
            <svg
              viewBox="0 0 448 512"
              fill="currentColor"
              height="1em"
              width="1em"
            >
              <path d="M64 32C28.7 32 0 60.7 0 96v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H64zm128 224h48c17.7 0 32-14.3 32-32s-14.3-32-32-32h-48v64zm48 64h-48v32c0 17.7-14.3 32-32 32s-32-14.3-32-32V168c0-22.1 17.9-40 40-40h72c53 0 96 43 96 96s-43 96-96 96z" />
            </svg>
          </CheckBox>



        </div>

      </form>
    </div>
  )
}

export default NewPlace
