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
	RegisterRouter(ctx context.Context) error
}

type usercent_App struct {
	weaver.Implements[UsercentApp]
}

func (u *usercent_App) RegisterTables(ctx context.Context) error {

	var err error
	log := u.Logger()
	//建表,每个模块都有自己的Tables
	for _, t := range Tables {
		err = global.GLB_DB.AutoMigrate(t)
		if nil != err {
			if err != nil {
				log.Error("register usercent_App table failed", err)
				os.Exit(0)
			}
			return err
		}

		log.Info("register " + t.(model.TableName_).TableName() + "success")
	}

	log.Info("register usercent_App table success")

	return err
}

func (u *usercent_App) RegisterRouter(ctx context.Context) error {

	return nil
}
