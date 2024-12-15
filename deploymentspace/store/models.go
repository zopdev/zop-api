package store

type DeploymentSpace struct {
	ID             int64  `json:"id"`
	CloudAccountID int64  `json:"cloud_account_id"`
	EnvironmentID  int64  `json:"environment_id"`
	Type           string `json:"type"`
	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`
}
