package taskApp

import (
	"context"
	"net/http"
	"new_it/app/taskApp/model"
	"new_it/global"
	"os"

	"github.com/ServiceWeaver/weaver"
)

type TaskcentApp interface {
	RegisterTables(ctx context.Context) error
	RegisterRouter(ctx context.Context) error
	RegisterDBdata(ctx context.Context) error
}

type taskcent_App struct {
	weaver.Implements[TaskcentApp]
}

//var uApi api.usercentApi

func (u *taskcent_App) RegisterTables(ctx context.Context) error {

	var err error
	log := u.Logger()
	//建表,每个模块都有自己的Tables
	for _, t := range Tables {
		err = global.GLB_DB.AutoMigrate(t)
		if nil != err {
			if err != nil {
				log.Error("register taskcent_App table failed", err)
				os.Exit(0)
			}
			return err
		}

		log.Info("register " + t.(model.TableName_).TableName() + " table success")
	}

	log.Info("register taskcent_App table success")

	return err
}

func (u *taskcent_App) RegisterRouter(ctx context.Context) error {

	var err error
	log := u.Logger()
	//每个组件或APP都有自己的路由
	for _, t := range RouterList {

		http.HandleFunc(t.router, t.handler)

		log.Info("register " + t.router + " success")
	}

	log.Info("register taskcent_App Router success")

	return err
}

func (u *taskcent_App) RegisterDBdata(ctx context.Context) error {
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
