package service_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zop-api/deploymentspace"
	clusterStore "github.com/zopdev/zop-api/deploymentspace/cluster/store"
	"github.com/zopdev/zop-api/deploymentspace/service"
	"github.com/zopdev/zop-api/deploymentspace/store"
)

var errTest = errors.New("service error")

func TestService_AddDeploymentSpace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockDeploymentSpaceStore(ctrl)
	mockClusterService := deploymentspace.NewMockDeploymentSpace(ctrl)

	ctx := &gofr.Context{}

	deploymentSpace := &service.DeploymentSpace{
		CloudAccount: service.CloudAccount{
			ID:         1,
			Provider:   "aws",
			ProviderID: "provider-123",
		},
		Type: service.Type{Name: "test-type"},
		DeploymentSpace: map[string]interface{}{
			"key": "value",
		},
	}

	mockCluster := clusterStore.Cluster{
		DeploymentSpaceID: 1,
		Provider:          "aws",
		ProviderID:        "provider-123",
	}

	testCases := []struct {
		name          string
		mockBehavior  func()
		input         *service.DeploymentSpace
		envID         int
		expectedError error
	}{
		{
			name: "success",
			mockBehavior: func() {
				mockStore.EXPECT().
					Insert(ctx, gomock.Any()).
					Return(&store.DeploymentSpace{ID: 1}, nil)
				mockClusterService.EXPECT().
					Add(ctx, gomock.Any()).
					Return(mockCluster, nil)
			},
			input:         deploymentSpace,
			envID:         1,
			expectedError: nil,
		},
		{
			name: "store layer error",
			mockBehavior: func() {
				mockStore.EXPECT().
					Insert(ctx, gomock.Any()).
					Return(nil, errTest)
			},
			input:         deploymentSpace,
			envID:         1,
			expectedError: errTest,
		},
		{
			name: "cluster service error",
			mockBehavior: func() {
				mockStore.EXPECT().
					Insert(ctx, gomock.Any()).
					Return(&store.DeploymentSpace{ID: 1}, nil)
				mockClusterService.EXPECT().
					Add(ctx, gomock.Any()).
					Return(nil, errTest)
			},
			input:         deploymentSpace,
			envID:         1,
			expectedError: errTest,
		},
		{
			name:         "invalid request body",
			mockBehavior: func() {},
			input: &service.DeploymentSpace{
				CloudAccount:    service.CloudAccount{},
				Type:            service.Type{},
				DeploymentSpace: nil, // Invalid DeploymentSpace
			},
			envID:         1,
			expectedError: http.ErrorInvalidParam{Params: []string{"body"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			svc := service.New(mockStore, mockClusterService)
			_, err := svc.Add(ctx, tc.input, tc.envID)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_FetchDeploymentSpace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockDeploymentSpaceStore(ctrl)
	mockClusterService := deploymentspace.NewMockDeploymentSpace(ctrl)

	ctx := &gofr.Context{}

	mockDeploymentSpace := &store.DeploymentSpace{
		ID:             1,
		CloudAccountID: 1,
		EnvironmentID:  1,
		Type:           "test-type",
	}

	mockCluster := clusterStore.Cluster{
		ID:                1,
		DeploymentSpaceID: 1,
		Name:              "test-cluster",
	}

	testCases := []struct {
		name          string
		mockBehavior  func()
		envID         int
		expectedError error
	}{
		{
			name: "success",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetByEnvironmentID(ctx, 1).
					Return(mockDeploymentSpace, nil)
				mockClusterService.EXPECT().
					FetchByDeploymentSpaceID(ctx, 1).
					Return(mockCluster, nil)
			},
			envID:         1,
			expectedError: nil,
		},
		{
			name: "store layer error",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetByEnvironmentID(ctx, 1).
					Return(nil, errTest)
			},
			envID:         1,
			expectedError: errTest,
		},
		{
			name: "no cluster found",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetByEnvironmentID(ctx, 1).
					Return(mockDeploymentSpace, nil)
				mockClusterService.EXPECT().
					FetchByDeploymentSpaceID(ctx, 1).
					Return(nil, sql.ErrNoRows)
			},
			envID:         1,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			svc := service.New(mockStore, mockClusterService)
			resp, err := svc.Fetch(ctx, tc.envID)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}
