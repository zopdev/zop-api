package service

import (
	"encoding/json"
	"github.com/zopdev/zop-api/deploymentspace"
	"github.com/zopdev/zop-api/deploymentspace/cluster/store"
	"gofr.dev/pkg/gofr"
)

type Service struct {
	store store.ClusterStore
}

func New(store store.ClusterStore) deploymentspace.DeploymentSpace {
	return &Service{
		store: store,
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
