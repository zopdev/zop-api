package main

import (
	"github.com/zopdev/zop-api/cloudaccounts/handler"
	"github.com/zopdev/zop-api/cloudaccounts/service"
	"github.com/zopdev/zop-api/cloudaccounts/store"
	"github.com/zopdev/zop-api/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())

	cloudAccountStore := store.New()
	cloudAccountService := service.New(cloudAccountStore)
	cloudAccountHandler := handler.New(cloudAccountService)

	app.POST("/cloud-accounts", cloudAccountHandler.AddCloudAccount)
	app.GET("/cloud-accounts", cloudAccountHandler.ListCloudAccounts)

	app.Run()
}
