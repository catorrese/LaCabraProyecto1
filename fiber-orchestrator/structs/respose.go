package structs

type Response struct {
	Data       []byte
	StatusCode int
	Err        error
}
