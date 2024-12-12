package service

import (
	"github.com/zopdev/zop-api/environments/store"
	"gofr.dev/pkg/gofr"
)

type EnvironmentService interface {
	FetchAllEnvironments(ctx *gofr.Context, applicationID int) ([]store.Environment, error)
	AddEnvironment(ctx *gofr.Context, environemt *store.Environment) (*store.Environment, error)
	UpdateEnvironments(ctx *gofr.Context, environments []store.Environment) ([]store.Environment, error)
}
