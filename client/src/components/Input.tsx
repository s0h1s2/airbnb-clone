interface Props {
  type: string
  placeholder: string
  field: any

}
const Input = ({ type, placeholder, field }: Props) => {
  return (
    <>
      <input placeholder={placeholder} type={type} className="w-full border my-2 py-2 px-3 rounded-xl" {...field} />
    </>
  )
}

export default Input
