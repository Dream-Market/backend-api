package auth

import (
	"fmt"

	"github.com/Dream-Market/backend-api/config"
	"github.com/Dream-Market/backend-api/pkg/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c config.Config) pb.AuthServiceClient {
	// Add SSL
	url := c.AuthSrvcHost + ":" + c.AuthSrvcPort
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
