import type { PropertyLocation } from "./location"
export type ListingUserFavorites = {
  userId: number
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
  username: string
  email: string
}
