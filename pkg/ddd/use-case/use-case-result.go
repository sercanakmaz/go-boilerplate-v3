package use_case

type UseCaseResult[T any] struct {
	Content        T
	Headers        map[string]string
	HttpStatusCode int
}

func NewUseCaseResult[T any](httpStatusCode int) *UseCaseResult[T] {
	return &UseCaseResult[T]{HttpStatusCode: httpStatusCode}
}
func NewUseCaseResultWithContent[T any](httpStatusCode int, content T) *UseCaseResult[T] {
	return NewUseCaseResultAll[T](httpStatusCode, content, nil)
}

func NewUseCaseResultAll[T any](httpStatusCode int, content T, headers map[string]string) *UseCaseResult[T] {
	return &UseCaseResult[T]{HttpStatusCode: httpStatusCode, Content: content, Headers: headers}
}
