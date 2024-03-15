import { StateCreator } from "zustand";
export interface IUser {
  email: string,
  name: string
}
export interface IUserSlice {
  user: IUser | null
  setUser: (user: IUser) => void
}
export const createUserSlice: StateCreator<IUserSlice> = (set) => ({
  setUser(user) {
    set(() => ({ user: user }))
  },
  user: null
})
