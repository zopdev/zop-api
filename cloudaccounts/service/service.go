package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/cloudaccounts/store"
)

type Service struct {
	store store.CloudAccountStore
}

// New creates a new CloudAccountService with the provided CloudAccountStore.
func New(clStore store.CloudAccountStore) CloudAccountService {
	return &Service{store: clStore}
}

// AddCloudAccount adds a new cloud account to the store if it doesn't already exist.
func (s *Service) AddCloudAccount(ctx *gofr.Context, cloudAccount *store.CloudAccount) (*store.CloudAccount, error) {
	//nolint:gocritic //addition of more providers
	switch strings.ToUpper(cloudAccount.Provider) {
	case gcp:
		err := fetchGCPProviderDetails(cloudAccount)
		if err != nil {
			return nil, err
		}
	}

	tempCloudAccount, err := s.store.GetCloudAccountByProvider(ctx, cloudAccount.Provider, cloudAccount.ProviderID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if tempCloudAccount != nil {
		return nil, http.ErrorEntityAlreadyExist{}
	}

	cloudAccount.CreatedAt = time.UTC.String()

	return s.store.InsertCloudAccount(ctx, cloudAccount)
}

// FetchAllCloudAccounts retrieves all cloud accounts from the store.
func (s *Service) FetchAllCloudAccounts(ctx *gofr.Context) ([]store.CloudAccount, error) {
	return s.store.GetALLCloudAccounts(ctx)
}

// fetchGCPProviderDetails retrieves and assigns GCP details for a cloud account.
func fetchGCPProviderDetails(cloudAccount *store.CloudAccount) error {
	var gcpCred gcpCredentials

	body, err := json.Marshal(cloudAccount.Credentials)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &gcpCred)
	if err != nil {
		return err
	}

	if gcpCred.ProjectID == "" {
		return http.ErrorInvalidParam{Params: []string{"credentials"}}
	}

	cloudAccount.ProviderID = gcpCred.ProjectID

	return nil
}
