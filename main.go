package main

import (
	"github.com/zopdev/zop-api/cloudaccounts/handler"
	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
	"github.com/zopdev/zop-api/deploymentspace/gke"

	appHandler "github.com/zopdev/zop-api/applications/handler"
	appService "github.com/zopdev/zop-api/applications/service"
	appStore "github.com/zopdev/zop-api/applications/store"

	envHandler "github.com/zopdev/zop-api/environments/handler"
	envService "github.com/zopdev/zop-api/environments/service"
	envStore "github.com/zopdev/zop-api/environments/store"

	"github.com/zopdev/zop-api/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())

	gkeSvc := gke.New()

	cloudAccountStore := store.New()
	cloudAccountService := service.New(cloudAccountStore, gkeSvc)
	cloudAccountHandler := handler.New(cloudAccountService)

	environmentStore := envStore.New()
	environmentService := envService.New(environmentStore)
	envrionmentHandler := envHandler.New(environmentService)

	applicationStore := appStore.New()
	applicationService := appService.New(applicationStore, environmentService)
	applicationHandler := appHandler.New(applicationService)

	app.POST("/cloud-accounts", cloudAccountHandler.AddCloudAccount)
	app.GET("/cloud-accounts", cloudAccountHandler.ListCloudAccounts)
	app.GET("/cloud-accounts/{id}/deployment-space/clusters", cloudAccountHandler.ListDeploymentSpace)
	app.GET("/cloud-accounts/{id}/deployment-space/namespaces", cloudAccountHandler.ListNamespaces)
	app.GET("/cloud-accounts/{id}/deployment-space/options", cloudAccountHandler.ListDeploymentSpaceOptions)

	app.POST("/applications", applicationHandler.AddApplication)
	app.GET("/applications", applicationHandler.ListApplications)
	app.GET("/applications/{id}", applicationHandler.GetApplication)

	app.POST("/environments", envrionmentHandler.AddEnvironment)
	app.PATCH("/environments", envrionmentHandler.UpdateEnvironments)
	app.GET("/applications/{id}/environments", envrionmentHandler.ListEnvironments)

	app.Run()
}
