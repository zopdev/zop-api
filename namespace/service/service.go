package service

import (
	"context"
	"fmt"

	container "google.golang.org/api/container/v1"
	"google.golang.org/api/option"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetNamespaceList() error {
	ctx := context.Background()

	// Initialize the Container Service client
	containerService, err := container.NewService(ctx, option.WithScopes(container.CloudPlatformScope))
	if err != nil {
		return err
	}

	// Replace these with your actual GCP project and location details
	projectID := "your-project-id"
	location := "your-cluster-location" // e.g., "us-central1-a" for zonal or "us-central1" for regional
	clusterName := "your-cluster-name"

	// Get cluster information
	clusterPath := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectID, location, clusterName)

	cluster, err := containerService.Projects.Locations.Clusters.Get(clusterPath).Do()
	if err != nil {
		return err
	}

	// Create the kubeconfig
	config := &api.Config{
		APIVersion: "v1",
		Kind:       "Config",
		Clusters: map[string]*api.Cluster{
			cluster.Name: {
				Server:                   "https://" + cluster.Endpoint,
				CertificateAuthorityData: []byte(cluster.MasterAuth.ClusterCaCertificate),
			},
		},
		Contexts: map[string]*api.Context{
			cluster.Name: {
				Cluster:  cluster.Name,
				AuthInfo: cluster.Name,
			},
		},
		CurrentContext: cluster.Name,
		AuthInfos: map[string]*api.AuthInfo{
			cluster.Name: {
				Token: cluster.MasterAuth.ClusterCaCertificate,
			},
		},
	}

	// Create the clientset
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{})

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	// List all namespaces
	_, err = clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	// Print all namespace names
	fmt.Println("Namespaces in cluster:", cluster.Name)

	return nil
}
