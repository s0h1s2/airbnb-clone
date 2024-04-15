import { client } from "@/lib/client";
import type { Listing } from "@/types/listing";
import type { OkResponseResult } from "@/types/response";
import { defineStore } from "pinia";

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
    }
  }
})
