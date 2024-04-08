import { client } from "@/lib/client";
import type { PropertyLocation } from "@/types/location";
import type { OkResponseResult } from "@/types/response";
import { defineStore } from "pinia";

interface Listing {
  title: string
  description: string
  category: string
  location: PropertyLocation
}
export const useListingStore = defineStore("listings", {
  state: () => ({
    listings: [] as Listing[],
    isLoaded: false
  }),
  actions: {
    async getListings() {
      await client.get<OkResponseResult<Listing[]>>("/listing").then((r) => {
        this.listings = r.data.data
      }).catch(() => {
      }).finally(() => {
        this.isLoaded = true
      })
    }
  }
})
