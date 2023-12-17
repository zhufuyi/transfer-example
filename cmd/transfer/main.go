// Package main is the grpc server of the application.
package main

import (
	"transfer/cmd/transfer/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.InitApp()
	servers := initial.CreateServices()
	closes := initial.Close(servers)

	a := app.New(servers, closes)
	a.Run()
}
