import type { PropertyLocation } from "./location"
export type ListingUserFavorites = {
  userId: number
}
interface User {
  name: string
  email: string
}

export interface Listing {
  id: number | string
  title: string
  description: string
  imageSrc: string
  country: string
  category: string
  price: number
  location: PropertyLocation
  favorites: ListingUserFavorites[]
  guestCount: number
  roomCount: number
  bathroomCount: number
  user: User
}
