package main

import (
	"fmt"
	"github.com/fahimimam/letsgo/cmd/web/config"
	"github.com/fahimimam/letsgo/cmd/web/routes"
)

func init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
}

func main() {
	config.InitServerConfig().StartServer(routes.AppRoutes(config.GetApp()))
}
