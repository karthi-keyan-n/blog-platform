package main

import (
	"context"
	"fmt"
	bp "github.com/blog-platform/proto/go/blogplatformservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := bp.NewBlogPlatformServiceClient(cc)
	createRequest := &bp.CreatePostRequest{
		Title:           "Getting started with golang",
		Content:         "The Go programming language is an open source project to make programmers more productive.",
		Author:          "Test user",
		PublicationDate: timestamppb.New(time.Now()),
		Tags:            []string{"Golang", "Beginner"},
	}

	postCreated, err := client.CreatePost(context.TODO(), createRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post created => %+v \n\n", postCreated)

	post, err := client.GetPost(context.TODO(), &bp.GetPostRequest{PostID: postCreated.GetPostID()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post retrieved => %+v \n\n", post)

	updateRequest := &bp.UpdatePostRequest{
		PostID: postCreated.GetPostID(),
		Title:  postCreated.GetTitle(),
		Content: "The Go programming language is an open source project to make programmers more productive." +
			"Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs " +
			"that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction",
		Author:          postCreated.GetAuthor(),
		PublicationDate: postCreated.GetPublicationDate(),
		Tags:            postCreated.GetTags(),
	}
	updatedPost, err := client.UpdatePost(context.TODO(), updateRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post updated => %+v \n\n", updatedPost)

	_, err = client.DeletePost(context.TODO(), &bp.DeletePostRequest{PostID: postCreated.GetPostID()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Post deleted \n\n")

	deletedPost, err := client.GetPost(context.TODO(), &bp.GetPostRequest{PostID: postCreated.GetPostID()})
	if err != nil {
		fmt.Printf("error occured/ %s post not exists \n\n", postCreated.GetPostID())
	} else {
		fmt.Printf("Post retrieved => %+v \n\n", deletedPost)
	}
}
