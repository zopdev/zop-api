package handler

import (
	"github.com/zopdev/zop-api/deploymentspace/service"
	"gofr.dev/pkg/gofr"
	"strconv"
	"strings"
)

type Handler struct {
	service service.DeploymentSpaceService
}

func New(svc service.DeploymentSpaceService) Handler {
	return Handler{
		service: svc,
	}
}

func (h *Handler) Add(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	id = strings.TrimSpace(id)

	environmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	deploymentSpace := service.DeploymentSpace{}

	err = ctx.Bind(&deploymentSpace)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.AddDeploymentSpace(ctx, &deploymentSpace, environmentID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
