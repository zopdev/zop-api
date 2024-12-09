package service

import (
	"gofr.dev/pkg/gofr"

	"github.com/zopdev/zop-api/cloudaccounts/store"
)

type CloudAccountService interface {
	AddCloudAccount(ctx *gofr.Context, accounts *store.CloudAccount) (*store.CloudAccount, error)
	FetchAllCloudAccounts(ctx *gofr.Context) ([]store.CloudAccount, error)
}
