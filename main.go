package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"new_it/app/configApp"
	"new_it/app/dictApp"
	"new_it/app/taskApp"
	"new_it/app/usercentApp"
	"new_it/global"
	"new_it/initialize"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Initialize the Service Weaver application.
	global.GLB_WEAVER_ROOT = weaver.Init(context.Background())
	global.GLB_LOG = global.GLB_WEAVER_ROOT.Logger()

	// Get a client to the Reverser component.
	conapp, err := weaver.Get[configApp.ConfigApp](global.GLB_WEAVER_ROOT)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	global.GLB_CFG_INFO, _ = conapp.GetConfigConfigInfo(ctx)
	global.GLB_DB = initialize.Gorm_mysql()

	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := global.GLB_WEAVER_ROOT.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------------------------")
	dict, err2 := weaver.Get[dictApp.DictcentApp](global.GLB_WEAVER_ROOT)
	if err2 != nil {
		log.Fatal(err2)
	}

	dict.RegisterTables(ctx)
	dict.RegisterRouter(ctx)
	dict.RegisterDBdata(ctx)

	user, err2 := weaver.Get[usercentApp.UsercentApp](global.GLB_WEAVER_ROOT)
	if err2 != nil {
		log.Fatal(err2)
	}

	user.RegisterTables(ctx)
	user.RegisterRouter(ctx)
	user.RegisterDBdata(ctx)

	task, err3 := weaver.Get[taskApp.TaskcentApp](global.GLB_WEAVER_ROOT)
	if err3 != nil {
		log.Fatal(err3)
	}

	task.RegisterTables(ctx)
	task.RegisterRouter(ctx)
	task.RegisterDBdata(ctx)

	http.Serve(lis, nil)
}
