package store

import "gofr.dev/pkg/gofr"

type DeploymentSpaceStore interface {
	InsertDeploymentSpace(ctx *gofr.Context, deploymentSpace *DeploymentSpace) (*DeploymentSpace, error)
	GetDeploymentSpaceByEnvID(ctx *gofr.Context, environmentID int) (*DeploymentSpace, error)
}
