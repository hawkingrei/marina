package main

import (
	"fmt"
	"github.com/hawkingrei/marina/bitbucket"
	"github.com/hawkingrei/marina/conf"
)

func main() {
	conf := conf.NewConfig()
	httpServer := bitbucket.NewHTTPServer(conf)
	if err := httpServer.Start(); err != nil {
		fmt.Println("Error: '%s'", err.Error())
	}
}
