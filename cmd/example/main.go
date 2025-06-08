package main

import (
	"context"
	"io"
	"log"
	"math/rand/v2"
	"net"
	"sync"

	"buf.build/go/protovalidate"
	"github.com/PavelNen/grpc/pkg/api/example"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}

	server := grpc.NewServer()
	service := &ExampleService{storage: make(map[uint64]*Post), validator: validator}
	example.RegisterExampleServer(server, service)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(server)

	log.Println("gRPC server is running on port 8082")
	if err := server.Serve(lis); err != io.EOF {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Post struct {
	ID       uint64
	Title    string
	Content  string
	AuthorID string
}

type ExampleService struct {
	example.UnimplementedExampleServer

	validator protovalidate.Validator
	storage   map[uint64]*Post
	mx        sync.RWMutex
}

func (s *ExampleService) CreatePost(ctx context.Context, req *example.CreatePostRequest) (*example.CreatePostResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "metadata not found")
	}

	log.Println("metadata:", md)

	if err := s.validator.Validate(req); err != nil {
		st := status.New(codes.InvalidArgument, codes.InvalidArgument.String())
		st, _ = st.WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "request",
					Description: err.Error(),
				},
			},
		})
		return nil, st.Err()
	}

	id := rand.Uint64()
	post := &Post{
		ID:       id,
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
		AuthorID: req.GetAuthorId(),
	}

	s.mx.Lock()
	s.storage[id] = post
	s.mx.Unlock()

	header := metadata.Pairs("x-user-id", "123")
	grpc.SetHeader(ctx, header)
	grpc.SetTrailer(ctx, header)

	return &example.CreatePostResponse{PostId: id}, nil
}
func (s *ExampleService) ListPosts(ctx context.Context, req *example.ListPostsRequest) (*example.ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
