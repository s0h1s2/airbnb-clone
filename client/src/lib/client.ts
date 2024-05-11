import axios from "axios";

export const client = axios.create({
  baseURL: "https://airbnb-clone-wt5p.onrender.com/api/v1/",
  withCredentials: true
})
