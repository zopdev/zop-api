package store

import "gofr.dev/pkg/gofr"

type Store struct{}

func New() DeploymentSpaceStore {
	return &Store{}
}

func (s *Store) InsertDeploymentSpace(ctx *gofr.Context, deploymentSpace *DeploymentSpace) (*DeploymentSpace, error) {
	res, err := ctx.SQL.ExecContext(ctx, INSERTQUERY, deploymentSpace.CloudAccountID, deploymentSpace.EnvironmentID, deploymentSpace.Type)
	if err != nil {
		return nil, err
	}

	deploymentSpace.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return deploymentSpace, nil
}

func (s *Store) GetDeploymentSpaceByEnvID(ctx *gofr.Context, environmentID int) (*DeploymentSpace, error) {
	deploymentSpace := DeploymentSpace{}

	err := ctx.SQL.QueryRowContext(ctx, GETQUERYBYENVID, environmentID).Scan(&deploymentSpace.ID, &deploymentSpace.CloudAccountID, &deploymentSpace.EnvironmentID,
		&deploymentSpace.Type, &deploymentSpace.CreatedAt, &deploymentSpace.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &deploymentSpace, nil
}
