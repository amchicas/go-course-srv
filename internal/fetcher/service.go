package fetcher

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/pkg/log"
)

type Service interface {
	GetCourse(ctx context.Context, id string) (*domain.Course, error)
}
type service struct {
	courseRepo domain.CourseRepository
	logger     *log.Logger
}

func (s *service) GetCourse(ctx context.Context, id string) (*domain.Course, error) {

	course, err := s.courseRepo.GetCourseById(ctx, id)
	if err != nil {

		return &domain.Course{}, err
	}

	return course, nil
}
func NewService(courseRepo domain.CourseRepository, logger *log.Logger) Service {

	return &service{
		courseRepo: courseRepo,
		logger:     logger,
	}

}
