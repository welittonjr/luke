package provider

import (
	"context"
	"errors"

	"github.com/welittonjr/luke/internal/api"
)

var ErrNoCredentials = errors.New("nenhuma credential configurada")

type Provider interface {
	Name() string
	Complete(context context.Context, req *api.Request) (*api.Response, error)
}
