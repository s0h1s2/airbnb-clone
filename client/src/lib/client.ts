import axios from "axios"

const client = axios.create({
  baseURL: "http://localhost:8080/api/v1/",
})

export interface OkResponseResult<T> {
  statusCode: number,
  data: T
}
export interface ErrResponseResult<T> {
  statusCode: number,
  errors: T
}

export default client
