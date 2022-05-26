package domain

import (
	"context"
	"time"
)

type Content struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Url      string `json:"url"`
	Article  string `json:"article"`
	Status   Status `json:"status"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
}

func NewContent(title, subtitle, url, article string, status Status) *Content {

	return &Content{
		Title:    title,
		Subtitle: subtitle,
		Url:      url,
		Article:  article,
		Status:   status,
		Created:  time.Now().Unix(),
		Modified: time.Now().Unix(),
	}

}

type ContentRepository interface {
	SaveContent(ctx context.Context, content *Content) error
	UpdateContent(ctx context.Context, content *Content, id string) error
	GetContentById(ctx context.Context, id string) (*Content, error)
	DeleteContentById(ctx context.Context, id string) error
}
