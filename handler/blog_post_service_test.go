package handler

import (
	"context"
	"fmt"
	"github.com/blog-platform/handler/internal/model"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"
	"time"

	"github.com/blog-platform/handler/testutils/mock"
	bp "github.com/blog-platform/proto/go/blogplatformservice"
)

func TestBlogPostService_GetPost(t *testing.T) {
	mockBlogPostRepo := mock.NewBlogPostRepository(t)
	postID := uuid.New()
	inValidPostID := uuid.New()
	post := &model.Post{
		Id:              postID,
		Title:           "Getting started with golang",
		Content:         "The Go programming language is an open source project to make programmers more productive.",
		Author:          "Test user",
		PublicationDate: time.Now(),
		Tags:            []string{"Golang", "Beginner"},
	}
	mockBlogPostRepo.On("GetPost", context.TODO(), postID).Return(post, nil)
	mockBlogPostRepo.On("GetPost", context.TODO(), inValidPostID).Return(nil, fmt.Errorf("invalid PostID"))

	type fields struct {
		blogPostRepository *mock.BlogPostRepository
		bp.UnimplementedBlogPlatformServiceServer
	}
	type args struct {
		cxt context.Context
		req *bp.GetPostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bp.PostResponse
		wantErr bool
	}{{
		name:   "Test GetPost returns correct blog response when valid uuid is passed",
		fields: fields{blogPostRepository: mockBlogPostRepo},
		args: args{
			cxt: context.TODO(),
			req: &bp.GetPostRequest{
				PostID: postID.String(),
			},
		},
		want: &bp.PostResponse{
			PostID:          postID.String(),
			Title:           post.Title,
			Content:         post.Content,
			Author:          post.Author,
			PublicationDate: timestamppb.New(post.PublicationDate),
			Tags:            post.Tags,
		},
		wantErr: false,
	}, {
		name:   "Test GetPost returns error when uuid is not present",
		fields: fields{blogPostRepository: mockBlogPostRepo},
		args: args{
			cxt: context.TODO(),
			req: &bp.GetPostRequest{
				PostID: inValidPostID.String(),
			},
		},
		want:    nil,
		wantErr: true,
	}, {
		name:   "Test GetPost returns error when invalid uuid is passed",
		fields: fields{blogPostRepository: mockBlogPostRepo},
		args: args{
			cxt: context.TODO(),
			req: &bp.GetPostRequest{
				PostID: "93c1faa6-e97e-4ff1-8f9e-0e8a69",
			},
		},
		want:    nil,
		wantErr: true,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bps := &BlogPostService{
				blogPostRepository:                     tt.fields.blogPostRepository,
				UnimplementedBlogPlatformServiceServer: tt.fields.UnimplementedBlogPlatformServiceServer,
			}
			got, err := bps.GetPost(tt.args.cxt, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPost() got = %v, want %v", got, tt.want)
			}
		})
	}
}
