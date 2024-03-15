import { useBoundStore } from "../state/boundState"

const IndexPage = () => {
  const user = useBoundStore((state) => state.user)
  return (
    <div>
      Hello {user?.name}
    </div >
  )
}

export default IndexPage
