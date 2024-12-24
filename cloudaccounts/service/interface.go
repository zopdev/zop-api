package service

import (
	"gofr.dev/pkg/gofr"

	"github.com/zopdev/zop-api/cloudaccounts/store"
)

type CloudAccountService interface {
	AddCloudAccount(ctx *gofr.Context, accounts *store.CloudAccount) (*store.CloudAccount, error)
	FetchAllCloudAccounts(ctx *gofr.Context) ([]store.CloudAccount, error)
	FetchDeploymentSpace(ctx *gofr.Context, cloudAccountID int) (interface{}, error)
	ListNamespaces(ctx *gofr.Context, id int, clusterName, clusterRegion string) (interface{}, error)
	FetchDeploymentSpaceOptions(ctx *gofr.Context, id int) ([]DeploymentSpaceOptions, error)
	FetchCredentials(ctx *gofr.Context, cloudAccountID int64) (interface{}, error)
}
