package main

import (
	"context"
	"fmt"
	"log"
	"new_it/app/configApp"
	"new_it/global"

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

	conapp.Get(ctx, "ddd")

	//fmt.Println(global.GLB_CFG_APP.DbType)
	//fmt.Println(global.GLB_CFG_APP.Mysql.Ip)
	// mysql, _ := conapp.GetConfigJWT(ctx)

	// fmt.Println(mysql)

	fmt.Println("hello new it !")
}
