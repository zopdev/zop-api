package store

import "gofr.dev/pkg/gofr"

type EnvironmentStore interface {
	InsertEnvironment(ctx *gofr.Context, environment *Environment) (*Environment, error)
	GetALLEnvironments(ctx *gofr.Context, applicationID int) ([]Environment, error)
	GetEnvironmentByName(ctx *gofr.Context, applicationID int, name string) (*Environment, error)
	UpdateEnvironment(ctx *gofr.Context, environment *Environment) (*Environment, error)
}
