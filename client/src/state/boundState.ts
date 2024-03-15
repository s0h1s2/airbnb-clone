import { create } from "zustand";
import { IUserSlice, createUserSlice } from "./user";

export const useBoundStore = create<IUserSlice>()((...a) => ({
  ...createUserSlice(...a)
}))
