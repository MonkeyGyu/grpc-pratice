package main

import (
	"context"
	"log"
	"net"

	"github.com/DonggyuLim/grpc-pratice/data"
	"github.com/DonggyuLim/grpc-pratice/protos/v1/post"
	"github.com/DonggyuLim/grpc-pratice/protos/v1/user"
	"github.com/DonggyuLim/grpc-pratice/simple-client-server/user_client"
	"google.golang.org/grpc"
)

const portNumber = "9001"

type postServer struct {
	post.PostServer
	userCli user.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *post.ListPostsByUserIdRequest) (*post.ListPostsByUserIdResponse, error) {
	userID := req.UserId
	resp, err := s.userCli.GetUser(ctx, &user.GetUserRequest{UserId: userID})

	if err != nil {
		return nil, err
	}

	var postMessages []*post.PostMessage

	for _, up := range data.UserPosts {
		if up.UserID != userID {
			continue
		}
		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}
		postMessages = up.Posts
		break
	}
	return &post.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

func (s *postServer) ListPosts(ctx context.Context, req *post.ListPostsRequest) (*post.ListPostsResponse, error) {
	var postMessages []*post.PostMessage

	for _, up := range data.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &user.GetUserRequest{UserId: up.UserID})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}
	return &post.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	userCli := user_client.GetUserClient("localhost:9000") //connection 선언
	grpcServer := grpc.NewServer()
	post.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server:%s", err)
	}
}
