import type { PropertyLocation } from "./location"

export interface Listing {
  id: number | string
  title: string
  description: string
  imageSrc: string
  location: PropertyLocation
}
