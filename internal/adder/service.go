package adder

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/pkg/log"
)

type Service interface {
	AddCourse(ctx context.Context, title, subtitlle, description string, status domain.Status) (*domain.Course, error)
}

type service struct {
	courseRepo domain.CourseRepository
	logger     *log.Logger
}

func (s *service) AddCourse(ctx context.Context,
	title, subtitle, description string, status domain.Status,
) (*domain.Course, error) {
	course := domain.NewCourse(title, subtitle, description, status)
	err := s.courseRepo.CreateCourse(ctx, course)
	if err != nil {

		return nil, err
	}
	return course, err
}
func NewService(courseRepo domain.CourseRepository, logger *log.Logger) Service {

	return &service{
		courseRepo: courseRepo,
		logger:     logger,
	}

}
