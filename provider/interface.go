package provider

import (
	"gofr.dev/pkg/gofr"
)

// Provider defines the interface for interacting with a cloud provider's resources.
// It includes methods for listing all clusters and retrieving namespaces for a given cluster.
//
// This interface can be implemented for various cloud providers such as AWS, GCP, or Azure.
// It allows users to interact with cloud infrastructure, retrieve clusters, and list namespaces.

type Provider interface {
	// ListAllClusters lists all clusters available for a given cloud account.
	//
	// ctx: The context for the request.
	// cloudAccount: The cloud account associated with the provider (e.g., AWS, GCP, Azure).
	// credentials: The authentication credentials used to access the provider's resources.
	//
	// Returns a ClusterResponse containing details of the available clusters, or an error if the request fails.
	ListAllClusters(ctx *gofr.Context, cloudAccount *CloudAccount, credentials interface{}) (*ClusterResponse, error)

	// ListNamespace retrieves namespaces for a given cluster within a cloud account.
	//
	// ctx: The context for the request.
	// cluster: The cluster for which to list namespaces.
	// cloudAccount: The cloud account associated with the provider.
	// credentials: The authentication credentials used to access the provider's resources.
	//
	// Returns the namespaces for the specified cluster, or an error if the request fails.
	ListNamespace(ctx *gofr.Context, cluster *Cluster, cloudAccount *CloudAccount, credentials interface{}) (interface{}, error)
	ListServices(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace string) (interface{}, error)
	ListDeployments(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace string) (interface{}, error)
	ListPods(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace string) (interface{}, error)
	ListCronJobs(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace string) (interface{}, error)

	GetService(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace, name string) (interface{}, error)
	GetDeployment(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace, name string) (interface{}, error)
	GetPod(ctx *gofr.Context, cluster *Cluster,
		cloudAccount *CloudAccount, credentials interface{}, namespace, name string) (interface{}, error)
	GetCronJob(ctx *gofr.Context, cluster *Cluster,
		cloudAcc *CloudAccount, creds any, namespace, name string) (any, error)
}
