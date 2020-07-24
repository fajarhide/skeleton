package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuthInterceptor(t *testing.T) {

	t.Run("PositiveTest", func(t *testing.T) {
		md := metadata.Pairs("authorization", "rcl-9cc18122bcb544031798a8b1b9003c38")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "rcl", nil
		}

		_, err := Auth(ctx, "JUST TEST", unaryInfo, unaryHandler)

		assert.NoError(t, err)
	})

	t.Run("NegativeTest", func(t *testing.T) {
		md := metadata.Pairs("authorization", "invalid-auth")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "rcl", nil
		}

		_, err := Auth(ctx, "JUST TEST", unaryInfo, unaryHandler)

		assert.Error(t, err)
	})

	t.Run("MissingAuthorization", func(t *testing.T) {
		md := metadata.Pairs("authorization", "")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "rcl", nil
		}

		_, err := Auth(ctx, "JUST TEST", unaryInfo, unaryHandler)

		assert.Error(t, err)
	})
}
