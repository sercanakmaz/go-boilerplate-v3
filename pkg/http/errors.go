package http

import "errors"

var ErrCommandBindFailed = errors.New("ErrCommandBindFailed")
var ErrUseCaseHandleFailed = errors.New("ErrUseCaseHandleFailed")
var ErrRabbitMQPublishFailed = errors.New("ErrRabbitMQPublishFailed")
