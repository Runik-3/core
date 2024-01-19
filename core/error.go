package core

type Response[T any] struct {
	Data  T
	Error error
}
