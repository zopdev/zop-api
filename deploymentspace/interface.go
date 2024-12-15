package deploymentspace

import (
	"gofr.dev/pkg/gofr"
)

type DeploymentSpace interface {
	FetchByDeploymentSpaceID(ctx *gofr.Context, id int) (interface{}, error)
	Add(ctx *gofr.Context, resource any) (interface{}, error)
}
