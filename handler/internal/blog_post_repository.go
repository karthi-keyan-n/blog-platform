package internal

import (
	"context"
	"errors"
	"github.com/blog-platform/handler/internal/model"
	"github.com/google/uuid"
)

type blogPostRepository struct {
}

//go:generate mockery --name BlogPostRepository --output ../testutils/mock/ --filename blog_post_respository.go --outpkg mock
type BlogPostRepository interface {
	CreatePost(cxt context.Context, post *model.Post) (*model.Post, error)
	GetPost(cxt context.Context, postID uuid.UUID) (*model.Post, error)
	UpdatePost(cxt context.Context, post *model.Post) (*model.Post, error)
	DeletePost(cxt context.Context, postID uuid.UUID) bool
}

func NewBlogPostRepository() BlogPostRepository {
	return &blogPostRepository{}
}

var persistedPosts = make(map[uuid.UUID]*model.Post)

func (bpr blogPostRepository) CreatePost(cxt context.Context, post *model.Post) (*model.Post, error) {
	post.Id = uuid.New()
	persistedPosts[post.Id] = post
	return post, nil
}

func (bpr blogPostRepository) GetPost(cxt context.Context, postID uuid.UUID) (*model.Post, error) {
	post, exists := persistedPosts[postID]
	if exists {
		return post, nil
	}

	return nil, errors.New("post does not exists")
}

func (bpr blogPostRepository) UpdatePost(cxt context.Context, post *model.Post) (*model.Post, error) {
	p, exists := persistedPosts[post.Id]
	if exists {
		p.Title = post.Title
		p.Content = post.Content
		p.PublicationDate = post.PublicationDate
		p.Tags = post.Tags

		return post, nil
	}
	return nil, errors.New("post does not exists")
}

func (bpr blogPostRepository) DeletePost(cxt context.Context, postID uuid.UUID) bool {
	delete(persistedPosts, postID)
	return true
}
