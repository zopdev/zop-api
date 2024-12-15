package service

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"

	"github.com/zopdev/zop-api/applications/store"
	"github.com/zopdev/zop-api/environments/service"
	"gofr.dev/pkg/gofr/http"
)

var (
	errTest = errors.New("service error")
)

func TestService_AddApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockApplicationStore(ctrl)
	mockEvironmentService := service.NewMockEnvironmentService(ctrl)
	ctx := &gofr.Context{}

	application := &store.Application{
		Name: "Test Application",
	}

	testCases := []struct {
		name          string
		mockBehavior  func()
		input         *store.Application
		expectedError error
	}{
		{
			name: "success",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetApplicationByName(ctx, "Test Application").
					Return(nil, sql.ErrNoRows)
				mockStore.EXPECT().
					InsertApplication(ctx, application).
					Return(application, nil)
			},
			input:         application,
			expectedError: nil,
		},
		{
			name: "application already exists",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetApplicationByName(ctx, "Test Application").
					Return(application, nil)
			},
			input:         application,
			expectedError: http.ErrorEntityAlreadyExist{},
		},
		{
			name: "error fetching application by name",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetApplicationByName(ctx, "Test Application").
					Return(nil, errTest)
			},
			input:         application,
			expectedError: errTest,
		},
		{
			name: "error inserting application",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetApplicationByName(ctx, "Test Application").
					Return(nil, sql.ErrNoRows)
				mockStore.EXPECT().
					InsertApplication(ctx, application).
					Return(nil, errTest)
			},
			input:         application,
			expectedError: errTest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			appService := New(mockStore, mockEvironmentService)
			_, err := appService.AddApplication(ctx, tc.input)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestService_FetchAllApplications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockApplicationStore(ctrl)
	mockEvironmentService := service.NewMockEnvironmentService(ctrl)

	ctx := &gofr.Context{}

	expectedApplications := []store.Application{
		{
			ID:        1,
			Name:      "Test Application",
			CreatedAt: "2023-12-11T00:00:00Z",
		},
	}

	testCases := []struct {
		name          string
		mockBehavior  func()
		expectedError error
	}{
		{
			name: "success",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetALLApplications(ctx).
					Return(expectedApplications, nil)
			},
			expectedError: nil,
		},
		{
			name: "error fetching applications",
			mockBehavior: func() {
				mockStore.EXPECT().
					GetALLApplications(ctx).
					Return(nil, errTest)
			},
			expectedError: errTest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()

			appService := New(mockStore, mockEvironmentService)
			applications, err := appService.FetchAllApplications(ctx)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, expectedApplications, applications)
			}
		})
	}
}
