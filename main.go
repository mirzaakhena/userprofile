package main

import (
	"flag"
	"fmt"
	"userprofile/application"
	"userprofile/application/registry"
)

func main() {

	appMap := map[string]func() application.RegistryContract{
		"simplememory": registry.NewSimplememory(),
		"usingdb":      registry.NewUsingdb(),
	}

	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		application.Run(app())
	} else {
		fmt.Println("You may try this app name:")
		for appName := range appMap {
			fmt.Printf("userprofile %s\n", appName)
		}
	}

}
