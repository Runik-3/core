package core

type Response[T any] struct {
	Data T
	// error objects don't send well to the frontend
	// it's best to pull out the message and send that.
	Error string
}
