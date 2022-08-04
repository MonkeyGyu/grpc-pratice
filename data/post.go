package data

import (
	postpb "github.com/DonggyuLim/grpc-pratice/protos/v1/post"
)

type PostData struct {
	UserID string
	Posts  []*postpb.PostMessage
}

var UserPosts = []*PostData{
	{
		UserID: "1",
		Posts: []*postpb.PostMessage{
			{
				PostId: "1",
				Author: "", // Post 서비스는 자체적으로 유저의 이름은 알지 못한다
				Title:  "gRPC 구축하기 (1)",
				Body:   "gRPC 구축하려면 이렇게 하면 된다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "2",
				Author: "", // Post 서비스는 자체적으로 유저의 이름은 알지 못한다
				Title:  "gRPC 구축하기 (2)",
				Body:   "gRPC server란 이렇다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
		},
	},
}
