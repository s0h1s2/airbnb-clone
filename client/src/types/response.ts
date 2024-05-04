export interface OkResponseResult<T> {
  statusCode: number,
  data: T,
  currentPage: number,
  totalPages: number
}
export interface ErrResponseResult<T> {
  statusCode: number,
  errors: T
}

