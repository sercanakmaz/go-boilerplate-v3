package ddd

type UseCaseResult[T any] struct {
	Content T
	Meta    map[string]string
}

func NewUseCaseResultWithContent[T any](content T) *UseCaseResult[T] {
	return NewUseCaseResultAll[T](content, nil)
}

func NewUseCaseResultAll[T any](content T, meta map[string]string) *UseCaseResult[T] {
	return &UseCaseResult[T]{Content: content, Meta: meta}
}
