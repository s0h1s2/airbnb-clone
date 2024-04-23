import { client } from "@/lib/client";
import type { Listing } from "@/types/listing";
import type { Reservation } from "@/types/reservation";
import type { OkResponseResult } from "@/types/response";
import { defineStore } from "pinia";
import { useToast } from "vue-toastification";
const toast = useToast()
export const useListingStore = defineStore("listings", {
  state: () => ({
    listings: [] as Listing[],
    isLoaded: false
  }),
  actions: {
    async getListings() {
      this.isLoaded = false
      await client.get<OkResponseResult<Listing[]>>("/listing").then((r) => {
        this.listings = r.data.data
      }).catch(() => {
      }).finally(() => {
        this.isLoaded = true
      })
    },
    async reserveListing(id: string | number, reservation: Reservation) {
      await client.post(`/listing/${id}/reserve`, reservation).then(() => {
        toast.success("Reserve listing successfully")
      }).catch(() => {
        toast.error("Unable to reserve listing.Try Again")
      })
    },
  }
})
