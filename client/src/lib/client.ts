import axios from "axios";

const base = import.meta.env.BASE_URL

export const client = axios.create({
  baseURL: base,
  withCredentials: true
})
