import axios from "axios";

const apiBase = import.meta.env.VITE_API_BASE_URL

export const client = axios.create({
  baseURL: apiBase,
  withCredentials: true
})
