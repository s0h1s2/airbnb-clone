import { client } from "@/lib/client";
import type { Listing } from "@/types/listing";
import type { Reservation } from "@/types/reservation";
import type { OkResponseResult } from "@/types/response";
import type { Axios } from "axios";
import { defineStore } from "pinia";
interface ApiParams {
  category?: string
}
export const useListingStore = defineStore("listings", {
  state: () => ({
    listings: [] as Listing[],
    isLoaded: false
  }),
  actions: {
    getListings(params?: ApiParams) {
      this.isLoaded = false
      client.get<OkResponseResult<Listing[]>>("/listing", { params: params }).then((r) => {
        this.listings = r.data.data
      }).catch(() => {
      }).finally(() => {
        this.isLoaded = true
      })
    },
    reserveListing(id: string | number, reservation: Reservation): Promise<Axios> {
      return client.post(`/listing/${id}/reserve`, reservation)
    },
  }
})
