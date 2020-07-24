package router

import (
	"fmt"
	"os"
	"strconv"

	rpc "github.com/fajarhide/skeleton/grpc/servers"
	"github.com/fajarhide/skeleton/helper"
	userPresenter "github.com/fajarhide/skeleton/modules/user/presenter"
	log "github.com/sirupsen/logrus"
)

// GRPCDefaultPort default port for GRPC
const GRPCDefaultPort = 8082

// GRPCServerMain - function for initializing main GRPC server
func (s *Service) GRPCServerMain() {

	// set GRPC port
	port := GRPCDefaultPort
	portGRPC, ok := os.LookupEnv("PORT_GRPC")
	if ok {
		intPort, _ := strconv.Atoi(portGRPC)
		port = intPort
	}

	helper.Log(log.InfoLevel, fmt.Sprintf("GRPC server will be running on port %d", port), "grpc_main", "initiate_grpc")

	uGRPCHandler := userPresenter.NewGRPCHandler(s.UserUseCase)

	grpcServer := rpc.NewGRPCServer(uGRPCHandler)

	err := grpcServer.Serve(uint(port))

	if err != nil {
		err = fmt.Errorf("error in Startup: %s", err.Error())
		helper.Log(log.ErrorLevel, err.Error(), "grpc_main", "serve_grpc")
		os.Exit(1)
	}

}
