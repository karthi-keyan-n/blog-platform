package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id              uuid.UUID
	Title           string
	Content         string
	Author          string
	PublicationDate time.Time
	Tags            []string
}
