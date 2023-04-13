package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := global.GLB_WEAVER_ROOT.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

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
	user.RegisterRouter(ctx)

	http.Serve(lis, nil)
}
