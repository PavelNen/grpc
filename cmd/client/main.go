package main

import (
	"context"
	"log"

	"github.com/PavelNen/grpc/pkg/api/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := example.NewExampleClient(conn)

	var header, trailer metadata.MD

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("client-user-id", "123"))

	resp, err := client.CreatePost(ctx, &example.CreatePostRequest{
		Title:    "Go is awesome",
		Content:  "Go is a programming language that makes it easy to build simple, reliable, and efficient software.",
		AuthorId: "1",
	}, grpc.Header(&header), grpc.Trailer(&trailer))


	if err != nil {
		switch status.Code(err) {
		case codes.InvalidArgument:
			log.Println("Некорректный запрос")
		default:
			log.Fatal(err)
		}

		if st, ok := status.FromError(err); ok {
			log.Println("code:", st.Code(), "messgae:", st.Message(), "details:", st.Details())
		} else {
			log.Println("not grpc error")
		}
	}

	log.Println("header:", header)
	log.Println("trailer:", trailer)

	log.Println(resp.GetPostId())

	bytes, err := protojson.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bytes))

}
