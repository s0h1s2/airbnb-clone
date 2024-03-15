import { DNA } from "react-loader-spinner"
import { useBoundStore } from "../state/boundState"
import { Navigate } from 'react-router-dom'
const ProtectedPage = () => {
  const user = useBoundStore((state) => state.user)
  if (user == undefined) {
    return (
      <DNA
        visible={true}
        height="80"
        width="80"
        ariaLabel="dna-loading"
        wrapperClass="text-center"
      />
    )
  }
  return user === null && <Navigate to={"/login"} replace />
}

export default ProtectedPage
