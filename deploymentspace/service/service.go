package service

import (
	"encoding/json"
	"github.com/zopdev/zop-api/deploymentspace"
	"github.com/zopdev/zop-api/deploymentspace/store"

	clusterStore "github.com/zopdev/zop-api/deploymentspace/cluster/store"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
)

type Service struct {
	store          store.DeploymentSpaceStore
	clusterService deploymentspace.DeploymentSpace
}

func New(store store.DeploymentSpaceStore, clusterSvc deploymentspace.DeploymentSpace) DeploymentSpaceService {
	return &Service{store: store, clusterService: clusterSvc}
}

func (s *Service) AddDeploymentSpace(ctx *gofr.Context, deploymentSpace *DeploymentSpace, environmentID int) (*DeploymentSpace, error) {
	if deploymentSpace.DeploymentSpace == nil {
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
	}

	dpSpace := store.DeploymentSpace{
		CloudAccountID: deploymentSpace.CloudAccount.ID,
		EnvironmentID:  int64(environmentID),
		Type:           deploymentSpace.Type.Name,
	}

	ds, err := s.store.InsertDeploymentSpace(ctx, &dpSpace)
	if err != nil {
		return nil, err
	}

	cl := clusterStore.Cluster{}
	bytes, err := json.Marshal(deploymentSpace.DeploymentSpace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &cl)
	if err != nil {
		return nil, err
	}

	cl.Provider = deploymentSpace.CloudAccount.Provider
	cl.ProviderID = deploymentSpace.CloudAccount.ProviderID
	cl.DeploymentSpaceID = ds.ID

	clResp, err := s.clusterService.Add(ctx, cl)
	if err != nil {
		return nil, err
	}

	deploymentSpace.DeploymentSpace = ds
	deploymentSpace.DeploymentSpace = clResp

	return deploymentSpace, nil
}

func (s *Service) FetchDeploymentSpace(ctx *gofr.Context, environmentID int) (*DeploymentSpaceResp, error) {
	deploymentSpace, err := s.store.GetDeploymentSpaceByEnvID(ctx, environmentID)
	if err != nil {
		return nil, err
	}

	resp, err := s.clusterService.FetchByDeploymentSpaceID(ctx, int(deploymentSpace.ID))
	if err != nil {
		return nil, err
	}

	return &DeploymentSpaceResp{
		DeploymentSpace: deploymentSpace,
		Cluster:         resp,
	}, nil
}
