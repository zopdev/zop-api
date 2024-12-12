package handler

import (
	"github.com/zopdev/zop-api/environments/service"
	"github.com/zopdev/zop-api/environments/store"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"strconv"
	"strings"
)

type Handler struct {
	service service.EnvironmentService
}

func New(svc service.EnvironmentService) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) AddEnvironment(ctx *gofr.Context) (interface{}, error) {
	environment := store.Environment{}

	err := ctx.Bind(&environment)
	if err != nil {
		return nil, err
	}

	err = validateEnvironment(&environment)
	if err != nil {
		return nil, err
	}

	res, err := h.service.AddEnvironment(ctx, &environment)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) ListEnvironments(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	id = strings.TrimSpace(id)

	applicationID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	res, err := h.service.FetchAllEnvironments(ctx, applicationID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) UpdateEnvironments(ctx *gofr.Context) (interface{}, error) {
	environments := []store.Environment{}
	err := ctx.Bind(&environments)
	if err != nil {
		return nil, err
	}

	for i := range environments {
		err = validateEnvironment(&environments[i])
		if err != nil {
			return nil, err
		}
	}

	res, err := h.service.UpdateEnvironments(ctx, environments)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func validateEnvironment(environment *store.Environment) error {
	environment.Name = strings.TrimSpace(environment.Name)
	params := []string{}

	if environment.Name == "" {
		params = append(params, "name")
	}

	if environment.ID == 0 {
		params = append(params, "id")
	}

	if environment.ApplicationID == 0 {
		params = append(params, "application_id")
	}

	if len(params) > 0 {
		return http.ErrorInvalidParam{Params: params}
	}

	return nil
}
