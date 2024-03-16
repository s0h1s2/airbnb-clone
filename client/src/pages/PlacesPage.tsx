import { FaPlus } from "react-icons/fa"
import { Link } from "react-router-dom"

const PlacesPage = () => {
  return (
    <div className="mt-4">
      <div className="text-center">
        <Link className="inline-flex gap-1 bg-primary text-white py-2 px-6 rounded-full " to="/account/places/new">

          <FaPlus className="my-[3px]" />
          Add New Place
        </Link>
      </div>
      PlacesPage
    </div>
  )
}

export default PlacesPage
