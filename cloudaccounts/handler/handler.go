package handler

import (
	"strconv"
	"strings"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
)

const (
	nameLength = 255
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
		ctx.Logger.Error(err.Error())
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
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

func (h *Handler) ListDeploymentSpace(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	id = strings.TrimSpace(id)

	cloudAccountId, err := strconv.Atoi(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	res, err := h.service.FetchDeploymentSpace(ctx, cloudAccountId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ListNamespaces(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	id = strings.TrimSpace(id)

	cloudAccountId, err := strconv.Atoi(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	clusterName := strings.TrimSpace(ctx.Param("name"))
	clusterRegion := strings.TrimSpace(ctx.Param("region"))

	if clusterName == "" || clusterRegion == "" {
		return nil, http.ErrorInvalidParam{Params: []string{"cluster"}}
	}

	res, err := h.service.ListNamespaces(ctx, cloudAccountId, clusterName, clusterRegion)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ListDeploymentSpaceOptions(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	id = strings.TrimSpace(id)

	cloudAccountId, err := strconv.Atoi(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	res, err := h.service.FetchDeploymentSpaceOptions(ctx, cloudAccountId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// validateCloudAccount checks the required fields and values in a CloudAccount.
func validateCloudAccount(cloudAccount *store.CloudAccount) error {
	params := []string{}

	cloudAccount.Name = strings.TrimSpace(cloudAccount.Name)

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

	if len(cloudAccount.Name) > nameLength {
		return http.ErrorInvalidParam{Params: []string{"name"}}
	}

	return nil
}
