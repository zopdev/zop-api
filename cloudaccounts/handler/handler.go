package handler

import (
	"strings"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
)

type Handler struct {
	service service.CloudAccountService
}

// New creates a new Handler with the provided CloudAccountService.
func New(clService service.CloudAccountService) Handler {
	return Handler{service: clService}
}

// AddCloudAccount handles the addition of a new CloudAccount, binding input and validating it.
func (h *Handler) AddCloudAccount(ctx *gofr.Context) (interface{}, error) {
	cloudAccount := store.CloudAccount{}

	err := ctx.Bind(&cloudAccount)
	if err != nil {
		return nil, err
	}

	err = validateCloudAccount(&cloudAccount)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.AddCloudAccount(ctx, &cloudAccount)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListCloudAccounts retrieves all existing CloudAccounts using the service layer.
func (h *Handler) ListCloudAccounts(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.FetchAllCloudAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// validateCloudAccount checks the required fields and values in a CloudAccount.
func validateCloudAccount(cloudAccount *store.CloudAccount) error {
	params := []string{}

	if cloudAccount.Name == "" {
		params = append(params, "name")
	}

	if cloudAccount.Provider == "" {
		params = append(params, "provider")
	}

	if len(params) != 0 {
		return http.ErrorMissingParam{Params: params}
	}

	if !strings.EqualFold(cloudAccount.Provider, "gcp") {
		return http.ErrorInvalidParam{Params: []string{"provider"}}
	}

	return nil
}
