package store

type DeploymentSpace struct {
	ID               int64  `json:"id"`
	CloudAccountID   int64  `json:"cloudAccountId"`
	EnvironmentID    int64  `json:"environmentId"`
	CloudAccountName string `json:"cloudAccountName"`
	Type             string `json:"type"`
	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`
}

type Cluster struct {
	DeploymentSpaceID int64  `json:"deploymentSpaceId"`
	ID                int64  `json:"id"`
	Identifier        string `json:"identifier"`
	Name              string `json:"name"`
	Region            string `json:"region"`
	Provider          string `json:"provider"`
	ProviderID        string `json:"providerId"`
	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`

	Namespace Namespace `json:"namespace"`
}

type Namespace struct {
	Name string `json:"Name"`
}
