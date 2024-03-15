import { useNavigate } from "react-router-dom"
import client from "../lib/client"
import { useBoundStore } from "../state/boundState"
import { toast } from "react-toastify"

const ProfilePage = () => {
  const user = useBoundStore((state) => state.user)
  const removeUser = useBoundStore((state) => state.removeUser)
  const navigate = useNavigate()
  return (
    <div className="text-center max-w-lg mx-auto mt-8">
      Logged in as {user?.name} ({user?.email})
      <button onClick={async () => {
        client.post("/users/logout").then(() => {
          removeUser()
          navigate("/login")
          toast.error("Logout was successful")

        }).catch((e) => {
          console.log(e)
        })

      }} className="primary max-w-sm mt-4">Logout</button>
    </div>
  )
}

export default ProfilePage
