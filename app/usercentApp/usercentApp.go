package usercentApp

import (
	"context"
	"new_it/app/usercentApp/model"
	"new_it/global"
	"os"

	"github.com/ServiceWeaver/weaver"
)

type UsercentApp interface {
	RegisterTables(ctx context.Context) error
}

type usercent_App struct {
	weaver.Implements[UsercentApp]
}

func (u *usercent_App) RegisterTables(ctx context.Context) error {

	err := global.GLB_DB.AutoMigrate(
		model.SysUsers{},
	)

	log := u.Logger()
	if err != nil {
		log.Error("register table failed", err)
		os.Exit(0)
	}
	log.Info("register table success")

	return err
}
