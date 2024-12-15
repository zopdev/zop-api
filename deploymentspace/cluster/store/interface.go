package store

import "gofr.dev/pkg/gofr"

type ClusterStore interface {
	Insert(ctx *gofr.Context, cluster *Cluster) (*Cluster, error)
	GetByDeploymentSpaceID(ctx *gofr.Context, deploymentSpaceID int) (*Cluster, error)
}
