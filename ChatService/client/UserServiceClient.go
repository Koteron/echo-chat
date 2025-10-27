package client

import (
	"context"
	"time"

	pb "ChatService/UserService"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceGRPC struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func NewUserServiceClient(addr string) (*UserServiceGRPC, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)
	return &UserServiceGRPC{
		conn:   conn,
		client: client,
	}, nil
}

func (c *UserServiceGRPC) Close() error {
	return c.conn.Close()
}

func (c *UserServiceGRPC) GetDisplayNames(ids []string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	names, err := c.client.GetDisplayName(ctx, &pb.GetDisplayNameRequest{UserId: ids})
	return names.GetDisplayName(), err
}
