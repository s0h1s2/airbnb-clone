import { client } from "@/lib/client";
import type { OkResponseResult } from "@/types/response";
import { defineStore, } from "pinia";
interface User {
  name: string
  email: string
}
export const useUserStore = defineStore("user", {
  state: () => {
    return {
      user: null as User | null,
      isAuth: false as boolean
    }
  },
  actions: {
    async login({ email, password }: { email: string, password: string }) {
      this.$reset()
      await client.post<OkResponseResult<User>>("/users/auth", { email, password }).then((r) => {
        this.user = r.data.data
        this.isAuth = true
      }).catch(
        (e) => { console.log(e) }
      )
    }

  }
})
