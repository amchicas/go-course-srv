package domain

import (
	"context"
	"time"
)

type Course struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Votes       uint64 `json:"votes"`
	Students    uint64 `json:"students"`
	Status      Status `json:"status"`
	Created     int64  `json:"created"`
	Modified    int64  `json:"modified"`
}

func NewCourse(title, subtitle, description string, status Status) *Course {

	return &Course{
		Title:       title,
		Subtitle:    subtitle,
		Description: description,
		Status:      status,
		Created:     time.Now().Unix(),
		Modified:    time.Now().Unix(),
	}

}

type CourseRepository interface {
	CreateCourse(ctx context.Context, course *Course) error
	UpdateCourse(ctx context.Context, course *Course, id string) (*Course, error)
	GetCourseById(ctx context.Context, id string) (*Course, error)
	DeleteCourseById(ctx context.Context, id string) error
}
