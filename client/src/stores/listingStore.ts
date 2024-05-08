import { client } from "@/lib/client";
import type { Listing } from "@/types/listing";
import type { Reservation } from "@/types/reservation";
import type { OkResponseResult } from "@/types/response";
import type { Axios } from "axios";
import { defineStore } from "pinia";
interface ApiParams {
  category?: string
  country?: string
  bathrooms?: number
  guests?: number
  rooms?: number
  startDate?: string
  endDate?: string
}
export const useListingStore = defineStore("listings", {
  state: () => ({
    listings: [] as Listing[],
    isLoaded: false,
    page: 1,
    totalPage: 100,
  }),
  actions: {
    getListings(params?: ApiParams) {
      this.$reset()
      this.isLoaded = false
      client.get<OkResponseResult<Listing[]>>("/listing", { params: { ...params, page: this.page } }).then((r) => {
        this.listings = [...this.listings, ...r.data.data]
        this.totalPage = r.data.totalPages
        this.page++
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
