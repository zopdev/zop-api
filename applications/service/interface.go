package service

import (
	"github.com/zopdev/zop-api/applications/store"
	"gofr.dev/pkg/gofr"
)

type ApplicationService interface {
	AddApplication(ctx *gofr.Context, application *store.Application) (*store.Application, error)
	FetchAllApplications(ctx *gofr.Context) ([]store.Application, error)
}
