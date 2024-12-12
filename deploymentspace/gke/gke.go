package gke

import (
	"cloud.google.com/go/container/apiv1/containerpb"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zopdev/zop-api/deploymentspace"
	"gofr.dev/pkg/gofr"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"

	container "cloud.google.com/go/container/apiv1"

	container2 "google.golang.org/api/container/v1"
)

type GKE struct {
}

func New() deploymentspace.ClusterService {
	return &GKE{}
}

func (g *GKE) ListAllClusters(ctx *gofr.Context, cloudAccount *deploymentspace.CloudAccount,
	credentials interface{}) (*deploymentspace.ClusterResponse, error) {
	credBody, err := g.getCredGCP(credentials)
	if err != nil {
		return nil, err
	}

	client, err := g.getClusterManagerClientGCP(ctx, credBody)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	req := &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", cloudAccount.ProviderID),
	}

	resp, err := client.ListClusters(ctx, req)
	if err != nil {
		return nil, err
	}

	gkeClusters := make([]deploymentspace.Cluster, 0)

	for _, cl := range resp.Clusters {
		gkeCluster := deploymentspace.Cluster{
			Name:      cl.Name,
			ID:        cl.Id,
			Region:    cl.Location,
			Locations: cl.Locations,
		}

		for _, nps := range cl.NodePools {
			cfg := nps.GetConfig()

			nodepool := deploymentspace.NodePool{
				MachineType: cfg.MachineType,
				NodeVersion: nps.Version,
				CurrentNode: nps.InitialNodeCount,
				NodeName:    nps.Name,
			}

			gkeCluster.NodePools = append(gkeCluster.NodePools, nodepool)
		}

		gkeClusters = append(gkeClusters, gkeCluster)
	}

	response := &deploymentspace.ClusterResponse{
		Clusters: gkeClusters,
		NextPage: deploymentspace.NextPage{
			Name: "Namespace",
			Path: fmt.Sprintf("/cloud-accounts/%v/deployment-space/clusters", cloudAccount.ID),
			Params: map[string]string{
				"region": "region",
				"name":   "name",
			},
		},
	}
	return response, nil
}

func (g *GKE) ListNamespace(ctx *gofr.Context, cluster *deploymentspace.Cluster,
	cloudAccount *deploymentspace.CloudAccount, credentials interface{}) (interface{}, error) {

	// Step 1: Get GCP credentials
	credBody, err := g.getCredGCP(credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	// Step 2: Get cluster information
	gkeCluster, err := g.getClusterInfo(ctx, cluster, cloudAccount, credBody)
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster info: %w", err)
	}

	// Step 3: Create HTTP client with TLS configured
	client, err := g.createTLSConfiguredClient(gkeCluster.MasterAuth.ClusterCaCertificate)
	if err != nil {
		return nil, fmt.Errorf("failed to create TLS configured client: %w", err)
	}

	// Step 4: Fetch namespaces from the Kubernetes API
	apiEndpoint := fmt.Sprintf("https://%s/api/v1/namespaces", gkeCluster.Endpoint)
	namespaces, err := g.fetchNamespaces(ctx, client, credBody, apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch namespaces: %w", err)
	}

	return namespaces, nil
}

func (g *GKE) getClusterInfo(ctx *gofr.Context, cluster *deploymentspace.Cluster,
	cloudAccount *deploymentspace.CloudAccount, credBody []byte) (*container2.Cluster, error) {

	// Create the GCP Container service
	containerService, err := container2.NewService(ctx, option.WithCredentialsJSON(credBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create container service: %w", err)
	}

	// Construct the full cluster name
	clusterFullName := fmt.Sprintf("projects/%s/locations/%s/clusters/%s",
		cloudAccount.ProviderID, cluster.Region, cluster.Name)

	// Get the GKE cluster details
	gkeCluster, err := containerService.Projects.Locations.Clusters.Get(clusterFullName).
		Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get GKE cluster details: %w", err)
	}

	return gkeCluster, nil
}

func (g *GKE) createTLSConfiguredClient(caCertificate string) (*http.Client, error) {
	// Decode the Base64-encoded CA certificate
	caCertBytes, err := base64.StdEncoding.DecodeString(caCertificate)
	if err != nil {
		return nil, fmt.Errorf("failed to decode CA certificate: %w", err)
	}

	// Create a CA certificate pool
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCertBytes) {
		return nil, fmt.Errorf("failed to append CA certificate to pool")
	}

	// Create a custom HTTP client with the CA certificate
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return client, nil
}

func (g *GKE) fetchNamespaces(ctx *gofr.Context, client *http.Client, credBody []byte, apiEndpoint string) ([]string, error) {
	// Generate a JWT token from the credentials
	config, err := google.JWTConfigFromJSON(credBody, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT config: %w", err)
	}

	// Create a TokenSource
	tokenSource := config.TokenSource(ctx)

	// Get a token
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	// Make a request to the Kubernetes API to list namespaces
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle unexpected status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response: %s", body)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var namespaceResponse struct {
		Items []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &namespaceResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Extract namespace names
	var namespaces []string
	for _, item := range namespaceResponse.Items {
		namespaces = append(namespaces, item.Metadata.Name)
	}

	return namespaces, nil
}

func (g *GKE) getCredGCP(credentials any) ([]byte, error) {
	var cred gcpCredentials

	credBody, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(credBody, &cred)
	if err != nil {
		return nil, err
	}

	return json.Marshal(cred)
}

func (g *GKE) getClusterManagerClientGCP(ctx *gofr.Context, credentials []byte) (*container.ClusterManagerClient, error) {
	client, err := container.NewClusterManagerClient(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		return nil, err
	}

	return client, nil
}
