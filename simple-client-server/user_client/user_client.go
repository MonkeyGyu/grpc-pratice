package user_client

import (
	"sync"

	"github.com/DonggyuLim/grpc-pratice/protos/v1/user"
	"google.golang.org/grpc"
)

var (
	once sync.Once
	cli  user.UserClient
)

func GetUserClient(serviceHost string) user.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost, grpc.WithInsecure(), grpc.WithBlock())
		cli = user.NewUserClient(conn)
	})
	return cli
}
