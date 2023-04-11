package main

import (
	"context"
	"fmt"
	"log"
	"new_it/app/configApp"
	"new_it/app/usercentApp"
	"new_it/global"
	"new_it/initialize"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Initialize the Service Weaver application.
	global.GLB_WEAVER_ROOT = weaver.Init(context.Background())

	// Get a client to the Reverser component.
	conapp, err := weaver.Get[configApp.ConfigApp](global.GLB_WEAVER_ROOT)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	global.GLB_CFG_INFO, _ = conapp.GetConfigConfigInfo(ctx)
	global.GLB_DB = initialize.Gorm_mysql()

	//for test
	jwt, _ := conapp.GetConfigJWT(ctx)
	mysql, _ := conapp.GetConfigMysql(ctx)
	fmt.Println(jwt)
	fmt.Println(mysql)
	fmt.Println(global.GLB_CFG_INFO)
	//end test

	fmt.Println("--------------------------")

	user, err2 := weaver.Get[usercentApp.UsercentApp](global.GLB_WEAVER_ROOT)
	if err2 != nil {
		log.Fatal(err2)
	}

	user.RegisterTables(ctx)
}
