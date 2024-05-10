import { client } from "@/lib/client";
import type { Listing } from "@/types/listing";
import type { Reservation } from "@/types/reservation";
import type { OkResponseResult } from "@/types/response";
import type { Axios } from "axios";
import { defineStore } from "pinia";
import _ from "lodash"
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
    page: 0,
    totalPage: 100,
  }),
  actions: {
    getListings(params?: ApiParams) {
      this.page++
      this.isLoaded = false
      client.get<OkResponseResult<Listing[]>>("/listing", { params: { ...params, page: this.page } }).then((r) => {
        this.listings = _.unionBy(r.data.data, this.listings, "id")
        this.totalPage = r.data.totalPages
        if (this.page + 1 > this.totalPage) {
          return
        }
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
