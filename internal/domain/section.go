package domain

import (
	"context"
	"time"
)

type Section struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Created     int64  `json:"created"`
	Modified    int64  `json:"modified"`
}

func NewSection(title, description string, status Status) *Section {

	return &Section{
		Title:       title,
		Description: description,
		Status:      status,
		Created:     time.Now().Unix(),
		Modified:    time.Now().Unix(),
	}

}

type SectionRepository interface {
	SaveSection(ctx context.Context, idCourse string, section *Section) error
	UpdateSection(ctx context.Context, section *Section, idCourse, idSection string) error
	GetSectionByID(ctx context.Context, id string) (*Section, error)
	DeleteSectionByID(ctx context.Context, idCourse, idSection string) error
}
