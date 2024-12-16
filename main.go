package main

import (
	appHandler "github.com/zopdev/zop-api/applications/handler"
	appService "github.com/zopdev/zop-api/applications/service"
	appStore "github.com/zopdev/zop-api/applications/store"
	"github.com/zopdev/zop-api/cloudaccounts/handler"
	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
	clStore "github.com/zopdev/zop-api/deploymentspace/cluster/store"
	"github.com/zopdev/zop-api/provider/gcp"

	envHandler "github.com/zopdev/zop-api/environments/handler"
	envService "github.com/zopdev/zop-api/environments/service"
	envStore "github.com/zopdev/zop-api/environments/store"

	deployHandler "github.com/zopdev/zop-api/deploymentspace/handler"
	deployService "github.com/zopdev/zop-api/deploymentspace/service"
	deployStore "github.com/zopdev/zop-api/deploymentspace/store"

	clService "github.com/zopdev/zop-api/deploymentspace/cluster/service"

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

	deploymentStore := deployStore.New()
	clusterStore := clStore.New()
	clusterService := clService.New(clusterStore)
	deploymentService := deployService.New(deploymentStore, clusterService)

	deploymentHandler := deployHandler.New(deploymentService)

	environmentStore := envStore.New()
	environmentService := envService.New(environmentStore, deploymentService)
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

	app.POST("/applications/{id}/environments", envrionmentHandler.Add)
	app.PATCH("/environments", envrionmentHandler.Update)
	app.GET("/applications/{id}/environments", envrionmentHandler.List)

	app.POST("/environments/{id}/deploymentspace", deploymentHandler.Add)

	app.Run()
}
