package dictApp

import (
	"context"
	"net/http"
	"new_it/app/dictApp/model"
	"new_it/global"
	"os"

	"github.com/ServiceWeaver/weaver"
)

type DictcentApp interface {
	RegisterTables(ctx context.Context) error
	RegisterRouter(ctx context.Context) error
	RegisterDBdata(ctx context.Context) error
}

type dictcent_App struct {
	weaver.Implements[DictcentApp]
}

//var uApi api.usercentApi

// RegisterTables
//
//	@receiver u
//	@param ctx
//	@return error
func (u *dictcent_App) RegisterTables(ctx context.Context) error {

	var err error
	log := u.Logger()
	//建表,每个模块都有自己的Tables
	for _, t := range Tables {
		err = global.GLB_DB.AutoMigrate(t)
		if nil != err {
			if err != nil {
				log.Error("register dictcent_App table failed", err)
				os.Exit(0)
			}
			return err
		}

		log.Info("register " + t.(model.TableName_).TableName() + " table success")
	}

	log.Info("register dictcent_App table success")

	return err
}

func (u *dictcent_App) RegisterRouter(ctx context.Context) error {

	var err error
	log := u.Logger()
	//每个组件或APP都有自己的路由
	for _, t := range RouterList {

		http.HandleFunc(t.router, t.handler)

		log.Info("register " + t.router + " success")
	}

	log.Info("register dictcent_App Router success")

	return err
}

func (u *dictcent_App) RegisterDBdata(ctx context.Context) error {
	var err error
	log := u.Logger()

	for _, data := range Dbdatas {
		if data.CheckDataExist() {
			log.Info("init dbdata " + data.TableName() + " exist")
			continue
		} else {
			data.Initialize()
			log.Info("init dbdata " + data.TableName() + " success")
		}
	}
	return err
}
