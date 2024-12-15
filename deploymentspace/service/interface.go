package service

import (
	"gofr.dev/pkg/gofr"
)

type DeploymentSpaceService interface {
	AddDeploymentSpace(ctx *gofr.Context, deploymentSpace *DeploymentSpace, environmentID int) (*DeploymentSpace, error)
	FetchDeploymentSpace(ctx *gofr.Context, environmentID int) (*DeploymentSpaceResp, error)
}
