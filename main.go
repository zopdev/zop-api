package main

import (
	"github.com/zopdev/zop-api/cloudaccounts/handler"
	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
	"github.com/zopdev/zop-api/provider/gcp"

	appHandler "github.com/zopdev/zop-api/applications/handler"
	appService "github.com/zopdev/zop-api/applications/service"
	appStore "github.com/zopdev/zop-api/applications/store"
	"github.com/zopdev/zop-api/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())

	gkeSvc := gcp.New()

	cloudAccountStore := store.New()
	cloudAccountService := service.New(cloudAccountStore, gkeSvc)
	cloudAccountHandler := handler.New(cloudAccountService)

	applicationStore := appStore.New()
	applicationService := appService.New(applicationStore)
	applicationHandler := appHandler.New(applicationService)

	app.POST("/cloud-accounts", cloudAccountHandler.AddCloudAccount)
	app.GET("/cloud-accounts", cloudAccountHandler.ListCloudAccounts)
	app.GET("/cloud-accounts/{id}/deployment-space/clusters", cloudAccountHandler.ListDeploymentSpace)
	app.GET("/cloud-accounts/{id}/deployment-space/namespaces", cloudAccountHandler.ListNamespaces)
	app.GET("/cloud-accounts/{id}/deployment-space/options", cloudAccountHandler.ListDeploymentSpaceOptions)

	app.POST("/applications", applicationHandler.AddApplication)
	app.GET("/applications", applicationHandler.ListApplications)

	app.Run()
}
