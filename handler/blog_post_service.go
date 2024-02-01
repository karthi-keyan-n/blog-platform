package handler

import (
	"context"
	"fmt"
	repo "github.com/blog-platform/handler/internal"
	bp "github.com/blog-platform/proto/go/blogplatformservice"
	"github.com/google/uuid"
)

type BlogPostService struct {
	blogPostRepository repo.BlogPostRepository
	bp.UnimplementedBlogPlatformServiceServer
}

func NewBlogPostService() bp.BlogPlatformServiceServer {
	return &BlogPostService{blogPostRepository: repo.NewBlogPostRepository()}
}

func (bps *BlogPostService) CreatePost(cxt context.Context, req *bp.CreatePostRequest) (*bp.PostResponse, error) {
	postToAdd := repo.BuildPost(req)
	post, err := bps.blogPostRepository.CreatePost(cxt, postToAdd)
	if err != nil {
		fmt.Printf("error occured while creating a post %+v", err)
		return nil, err
	}
	return repo.BuildPostResponse(post), nil
}

func (bps *BlogPostService) GetPost(cxt context.Context, req *bp.GetPostRequest) (*bp.PostResponse, error) {
	postID, pErr := uuid.Parse(req.GetPostID())
	if pErr != nil {
		fmt.Printf("error occured while trying to parse postID - %+v", pErr)
		return nil, pErr
	}
	post, err := bps.blogPostRepository.GetPost(cxt, postID)
	if err != nil {
		fmt.Printf("error occured while fetching post %+v", err)
		return nil, err
	}
	return repo.BuildPostResponse(post), nil
}

func (bps *BlogPostService) UpdatePost(cxt context.Context, req *bp.UpdatePostRequest) (*bp.PostResponse, error) {
	postID, pErr := uuid.Parse(req.GetPostID())
	if pErr != nil {
		fmt.Printf("error occured while trying to parse postID - %+v", pErr)
		return nil, pErr
	}
	postToAdd := repo.BuildUpdatePost(postID, req)
	post, err := bps.blogPostRepository.UpdatePost(cxt, postToAdd)
	if err != nil {
		fmt.Printf("error occured while updating post %+v", err)
		return nil, err
	}
	return repo.BuildPostResponse(post), nil
}

func (bps *BlogPostService) DeletePost(cxt context.Context, req *bp.DeletePostRequest) (*bp.DeleteResponse, error) {
	postID, pErr := uuid.Parse(req.GetPostID())
	if pErr != nil {
		fmt.Printf("error occured while trying to parse postID - %+v", pErr)
		return nil, pErr
	}
	isDeleted := bps.blogPostRepository.DeletePost(cxt, postID)
	if !isDeleted {
		return nil, fmt.Errorf("unable to delete post")
	}
	return &bp.DeleteResponse{}, nil
}
