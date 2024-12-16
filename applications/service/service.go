package service

import (
	"errors"

	"database/sql"

	"github.com/zopdev/zop-api/applications/store"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
)

type Service struct {
	store store.ApplicationStore
}

func New(str store.ApplicationStore) ApplicationService {
	return &Service{store: str}
}

func (s *Service) AddApplication(ctx *gofr.Context, application *store.Application) (*store.Application, error) {
	tempApplication, err := s.store.GetApplicationByName(ctx, application.Name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if tempApplication != nil {
		return nil, http.ErrorEntityAlreadyExist{}
	}

	application, err = s.store.InsertApplication(ctx, application)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (s *Service) FetchAllApplications(ctx *gofr.Context) ([]store.Application, error) {
	return s.store.GetALLApplications(ctx)
}
