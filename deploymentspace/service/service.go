/*
Package service provides the implementation of the DeploymentSpaceService interface.
It manages the addition and retrieval of deployment spaces, including their associated clusters and cloud account details.
The service interacts with underlying data stores and cluster management components to fulfill requests.
*/
package service

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/zopdev/zop-api/deploymentspace"
	"github.com/zopdev/zop-api/deploymentspace/store"

	clusterStore "github.com/zopdev/zop-api/deploymentspace/cluster/store"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
)

// Service implements the DeploymentSpaceService interface.
// It uses a combination of deployment space and cluster stores to manage deployment space operations.
type Service struct {
	store          store.DeploymentSpaceStore
	clusterService deploymentspace.DeploymentEntity
}

// New initializes a new instance of Service with the provided deployment space store and cluster service.
//
// Parameters:
//
//	str - The deployment space store used for data persistence.
//	clusterSvc - The cluster service used for managing clusters.
//
// Returns:
//
//	DeploymentSpaceService - An instance of the DeploymentSpaceService interface.
func New(str store.DeploymentSpaceStore, clusterSvc deploymentspace.DeploymentEntity) DeploymentSpaceService {
	return &Service{store: str, clusterService: clusterSvc}
}

// Add adds a new deployment space along with its associated cluster to the system.
//
// Parameters:
//
//	ctx - The GoFR context that carries request-specific data.
//	deploymentSpace - The DeploymentEntity object containing cloud account, type, and deployment details.
//	environmentID - The ID of the environment where the deployment space is being created.
//
// Returns:
//
//	*DeploymentEntity - The newly created deployment space with updated details (including ID and cluster response).
//	error - Any error encountered during the add operation.
func (s *Service) Add(ctx *gofr.Context, deploymentSpace *DeploymentSpace, environmentID int) (*DeploymentSpace, error) {
	if deploymentSpace.DeploymentSpace == nil {
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
	}

	dpSpace := store.DeploymentSpace{
		CloudAccountID: deploymentSpace.CloudAccount.ID,
		EnvironmentID:  int64(environmentID),
		Type:           deploymentSpace.Type.Name,
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

	_, err = s.clusterService.DuplicateCheck(ctx, &cl)
	if err != nil {
		return nil, err
	}

	ds, err := s.store.Insert(ctx, &dpSpace)
	if err != nil {
		return nil, err
	}

	cl.DeploymentSpaceID = ds.ID

	clResp, err := s.clusterService.Add(ctx, cl)
	if err != nil {
		return nil, err
	}

	deploymentSpace.DeploymentSpace = ds
	deploymentSpace.DeploymentSpace = clResp

	return deploymentSpace, nil
}

// Fetch retrieves a deployment space and its associated cluster details by environment ID.
//
// Parameters:
//
//	ctx - The GoFR context that carries request-specific data.
//	environmentID - The ID of the environment for which the deployment space is being fetched.
//
// Returns:
//
//	*DeploymentSpaceResp - The deployment space response containing the deployment space and cluster details.
//	error - Any error encountered during the fetch operation.
func (s *Service) Fetch(ctx *gofr.Context, environmentID int) (*DeploymentSpaceResp, error) {
	deploymentSpace, err := s.store.GetByEnvironmentID(ctx, environmentID)
	if err != nil {
		return nil, err
	}

	resp, err := s.clusterService.FetchByDeploymentSpaceID(ctx, int(deploymentSpace.ID))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	cluster := store.Cluster{}

	err = json.Unmarshal(bytes, &cluster)
	if err != nil {
		return nil, err
	}

	return &DeploymentSpaceResp{
		DeploymentSpace: deploymentSpace,
		Cluster:         &cluster,
	}, nil
}
