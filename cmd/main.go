package main

import (
	"log"
	"os"

	"github.com/Odvin/go-hex-boilerplate/internal/adapters/app/api"
	"github.com/Odvin/go-hex-boilerplate/internal/adapters/core/arithmetic"
	rpc "github.com/Odvin/go-hex-boilerplate/internal/adapters/framework/left/grpc"
	"github.com/Odvin/go-hex-boilerplate/internal/adapters/framework/right/db"
	"github.com/Odvin/go-hex-boilerplate/internal/ports"
)

func main() {
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	var core ports.ArithmaticPort
	var appAdapter ports.APIPort
	var dbAdapter ports.DbPort
	var gRPCAdapter ports.GRPCPort

	dbAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbAdapter, core)

	gRPCAdapter = rpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
