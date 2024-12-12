package deploymentspace

import "gofr.dev/pkg/gofr"

type ClusterService interface {
	ListAllClusters(ctx *gofr.Context, cloudAccount *CloudAccount, credenitals interface{}) (*ClusterResponse, error)
	ListNamespace(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}) (interface{}, error)
}
