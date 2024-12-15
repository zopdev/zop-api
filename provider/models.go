package provider

type ClusterResponse struct {
	Clusters []Cluster `json:"options"`
	NextPage NextPage  `json:"nextPage"`
}

type NextPage struct {
	Name   string            `json:"Name"`
	Path   string            `json:"Path"`
	Params map[string]string `json:"Params"`
}
type Cluster struct {
	Name       string     `json:"name"`
	Identifier string     `json:"identifier"`
	Locations  []string   `json:"locations"`
	Region     string     `json:"region"`
	NodePools  []NodePool `json:"node_pools"`
	Type       string     `json:"type"`
}

type NodePool struct {
	MachineType       string   `json:"machine_type"`
	AvailabilityZones []string `json:"availability_zones"`
	NodeVersion       string   `json:"nodeVersion,omitempty"`
	CurrentNode       int32    `json:"currentNode"`
	NodeName          string   `json:"nodeName"`
}

type CloudAccount struct {
	ID int64 `json:"id"`
	// Name is the name of the cloud account.
	Name string `json:"name"`

	// Provider is the name of the cloud service provider.
	Provider string `json:"provider"`

	// ProviderID is the identifier for the provider account.
	ProviderID string `json:"providerId"`

	// ProviderDetails contains additional details specific to the provider.
	ProviderDetails interface{} `json:"providerDetails"`

	// Credentials hold authentication information for access to the provider.
	Credentials interface{} `json:"credentials,omitempty"`
}

type NamespaceResponse struct {
	Options []Namespace `json:"options"`
}

type Namespace struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
