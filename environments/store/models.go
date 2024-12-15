package store

type Environment struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Level         int    `json:"level"`
	ApplicationID int64  `json:"applicationId"`

	DeploymentSpace any `json:"deploymentSpace"`

	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`
}
