package internal

import (
	"github.com/blog-platform/handler/internal/model"
	bp "github.com/blog-platform/proto/go/blogplatformservice"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BuildPost(req *bp.CreatePostRequest) *model.Post {
	return &model.Post{
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		PublicationDate: req.GetPublicationDate().AsTime(),
		Tags:            req.GetTags(),
	}
}

func BuildPostResponse(post *model.Post) *bp.PostResponse {
	return &bp.PostResponse{
		PostID:          post.Id.String(),
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: timestamppb.New(post.PublicationDate),
		Tags:            post.Tags,
	}
}

func BuildUpdatePost(uuid uuid.UUID, req *bp.UpdatePostRequest) *model.Post {
	return &model.Post{
		Id:              uuid,
		Title:           req.GetTitle(),
		Content:         req.GetContent(),
		Author:          req.GetAuthor(),
		PublicationDate: req.GetPublicationDate().AsTime(),
		Tags:            req.GetTags(),
	}
}
