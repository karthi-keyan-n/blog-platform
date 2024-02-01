package internal

import (
	"context"
	"github.com/blog-platform/handler/internal/model"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func Test_blogPostRepository_GetPost(t *testing.T) {
	postID := uuid.New()
	post := &model.Post{
		Id:              postID,
		Title:           "Getting started with golang",
		Content:         "The Go programming language is an open source project to make programmers more productive.",
		Author:          "Test user",
		PublicationDate: time.Now(),
		Tags:            []string{"Golang", "Beginner"},
	}
	persistedPosts[postID] = post

	type args struct {
		postID  uuid.UUID
		context context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Post
		wantErr bool
	}{{
		name:    "Test GetPost returns correct record when valid uuid is passed",
		args:    args{postID: postID, context: context.TODO()},
		want:    persistedPosts[postID],
		wantErr: false,
	},
		{
			name:    "Test GetPost returns error when valid uuid is passed",
			args:    args{postID: uuid.New(), context: context.TODO()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bpr := &blogPostRepository{}
			got, err := bpr.GetPost(tt.args.context, tt.args.postID)
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
