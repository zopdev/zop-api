package store

type DeploymentSpace struct {
	ID               int64  `json:"id"`
	CloudAccountID   int64  `json:"cloud_account_id"`
	EnvironmentID    int64  `json:"environment_id"`
	CloudAccountName string `json:"cloud_account_name"`
	Type             string `json:"type"`
	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`
}

type Cluster struct {
	DeploymentSpaceID int64     `json:"deployment_space_id"`
	ID                int64     `json:"id"`
	Identifier        string    `json:"identifier"`
	Name              string    `json:"name"`
	Region            string    `json:"region"`
	Namespace         Namespace `json:"namespace"`
	Provider          string    `json:"provider"`
	ProviderID        string    `json:"provider_id"`

	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`
}

type Namespace struct {
	Name string `json:"Name"`
}
