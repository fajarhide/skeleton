package servers

import (
	"fmt"
	"net"

	"github.com/fajarhide/skeleton/grpc/middleware"
	"google.golang.org/grpc"

	userPresenter "github.com/fajarhide/skeleton/modules/user/presenter"
	userPB "github.com/fajarhide/skeleton/protogo/user"
)

// Server data structure for grpc server
type Server struct {
	userGRPCHandler *userPresenter.UserGRPCHandler
}

// NewGRPCServer function for creating GRPC server
func NewGRPCServer(userGrpcHandler *userPresenter.UserGRPCHandler) *Server {
	return &Server{
		userGRPCHandler: userGrpcHandler,
	}
}

// Serve insecure server/ no server side encryption
func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		//Unary interceptor
		grpc.UnaryInterceptor(middleware.Auth),
	)

	//Register all sub server here
	userPB.RegisterUserServiceServer(server, s.userGRPCHandler)
	//end register server

	err = server.Serve(l)

	if err != nil {
		return err
	}

	return nil
}
