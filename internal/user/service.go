package user

import "github.com/thanh2k4/Chat-app/proto/gen"

type UserServer struct {
	gen.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}
