package service

import "github.com/zopdev/zop-api/applications/store"

type Application struct {
	ID   int64  `json:"Identifier"`
	Name string `json:"name"`

	// CreatedAt is the timestamp of when the cloud account was created.
	CreatedAt string `json:"createdAt"`

	// UpdatedAt is the timestamp of the last update to the cloud account.
	UpdatedAt string `json:"updatedAt"`

	// DeletedAt is the timestamp of when the cloud account was deleted, if applicable.
	DeletedAt string `json:"deletedAt,omitempty"`

	Environments []store.Environment `json:"environments"`
}
