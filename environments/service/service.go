package service

import (
	"database/sql"
	"errors"
	"time"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/environments/store"
)

type Service struct {
	store store.EnvironmentStore
}

func New(enStore store.EnvironmentStore) EnvironmentService {
	return &Service{store: enStore}
}

func (s *Service) AddEnvironment(ctx *gofr.Context, environemt *store.Environment) (*store.Environment, error) {
	tempEnvironment, err := s.store.GetEnvironmentByName(ctx, int(environemt.ApplicationID), environemt.Name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if tempEnvironment != nil {
		return nil, http.ErrorEntityAlreadyExist{}
	}

	environemt.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	return s.store.InsertEnvironment(ctx, environemt)
}

func (s *Service) FetchAllEnvironments(ctx *gofr.Context, applicationID int) ([]store.Environment, error) {
	return s.store.GetALLEnvironments(ctx, applicationID)
}

func (s *Service) UpdateEnvironments(ctx *gofr.Context, environments []store.Environment) ([]store.Environment, error) {
	for i := range environments {
		env, err := s.store.UpdateEnvironment(ctx, &environments[i])
		if err != nil {
			return nil, err
		}

		environments[i] = *env
	}

	return environments, nil
}
