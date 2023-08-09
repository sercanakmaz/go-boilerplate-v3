package use_case

import (
	"context"
)

type UseCaseContext struct {
	ctx     context.Context
	objects map[string]interface{}
}
