export interface OkResponseResult<T> {
    statusCode: number,
    data: T
}
export interface ErrResponseResult<T> {
    statusCode: number,
    errors: T
}
  