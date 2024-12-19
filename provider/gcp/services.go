package gcp

import (
	"encoding/json"
	"fmt"
	"github.com/zopdev/zop-api/provider"
	"gofr.dev/pkg/gofr"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

func (g *GCP) ListServices(ctx *gofr.Context, cluster *provider.Cluster,
	cloudAccount *provider.CloudAccount, credentials interface{}, namespace string) (interface{}, error) {
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

	// Step 4: Fetch services from the Kubernetes API
	apiEndpoint := fmt.Sprintf("https://%s/api/v1/namespaces/%s/services", gkeCluster.Endpoint, namespace)

	services, err := g.fetchServices(ctx, client, credBody, apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch services: %w", err)
	}

	return services, nil
}

// fetchServices fetches Kubernetes services from the specified namespace using the provided HTTP client.
func (*GCP) fetchServices(ctx *gofr.Context, client *http.Client, credBody []byte,
	apiEndpoint string) (*provider.ServiceResponse, error) {
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
		ctx.Logger.Errorf("failed to get token: %v", err)
		return nil, err
	}

	// Make a request to the Kubernetes API to list services
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiEndpoint, http.NoBody)
	if err != nil {
		ctx.Logger.Errorf("failed to create request: %w", err)
		return nil, err
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
		body, _ := io.ReadAll(resp.Body)
		ctx.Logger.Errorf("API call failed with status code %d: %s", resp.StatusCode, body)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var serviceResponse struct {
		Items []provider.Service `json:"items"`
	}

	if err := json.Unmarshal(body, &serviceResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Extract service details

	return &provider.ServiceResponse{
		Services: serviceResponse.Items,
		Metadata: provider.Metadata{
			Name: "services",
			Type: "kubernetes-cluster",
		},
	}, nil
}
