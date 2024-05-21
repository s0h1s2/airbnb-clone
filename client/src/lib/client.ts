import axios from "axios";

const apiBase = import.meta.env.API_BASE_URL

export const client = axios.create({
  baseURL: apiBase,
  withCredentials: true
})
