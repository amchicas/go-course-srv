package eraser

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/pkg/log"
)

type Service interface {
	DeleteCourseById(ctx context.Context, id string) error
}
type service struct {
	courseRepo domain.CourseRepository
	logger     *log.Logger
}

func (s *service) DeleteCourseById(ctx context.Context, id string) error {

	err := s.courseRepo.DeleteCourseById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func NewService(courseRepo domain.CourseRepository, logger *log.Logger) Service {

	return &service{
		courseRepo: courseRepo,
		logger:     logger,
	}

}
