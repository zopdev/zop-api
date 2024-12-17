package service

import (
	"database/sql"
	"encoding/json"
	"errors"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/deploymentspace"
	"github.com/zopdev/zop-api/deploymentspace/cluster/store"
)

type Service struct {
	store store.ClusterStore
}

func New(str store.ClusterStore) deploymentspace.DeploymentEntity {
	return &Service{
		store: str,
	}
}

func (s *Service) FetchByDeploymentSpaceID(ctx *gofr.Context, id int) (interface{}, error) {
	cluster, err := s.store.GetByDeploymentSpaceID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (s *Service) Add(ctx *gofr.Context, data any) (interface{}, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	cluster := store.Cluster{}

	err = json.Unmarshal(bytes, &cluster)
	if err != nil {
		return nil, err
	}

	resp, err := s.store.Insert(ctx, &cluster)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) DuplicateCheck(ctx *gofr.Context, data any) (interface{}, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	cluster := store.Cluster{}

	err = json.Unmarshal(bytes, &cluster)
	if err != nil {
		return nil, err
	}

	resp, err := s.store.GetByCluster(ctx, &cluster)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if resp != nil {
		return nil, http.ErrorEntityAlreadyExist{}
	}

	return nil, nil
}