// Package main is the grpc server of the application.
package main

import (
	"github.com/zhufuyi/transfer-example/cmd/transfer/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.InitApp()
	servers := initial.CreateServices()
	closes := initial.Close(servers)

	a := app.New(servers, closes)
	a.Run()
}
