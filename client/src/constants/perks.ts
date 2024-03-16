import { IconType } from "react-icons"
import { FaDog } from "react-icons/fa"
import { FaRegSnowflake, FaSquareParking } from "react-icons/fa6"
import { IoWifi } from "react-icons/io5"

type Perks = { name: string, icon: IconType }
export const PERKS: Perks[] = [
  {
    name: "Wifi",
    icon: IoWifi
  },
  {
    name: "Pet Allowed",
    icon: FaDog
  },
  {
    name: "Parking",
    icon: FaSquareParking
  },
  {
    name: "AC",
    icon: FaRegSnowflake
  },
] 
