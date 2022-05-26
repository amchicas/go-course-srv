package modifier

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/pkg/log"
)

type Service interface {
	UpdateCourse(ctx context.Context, id, title, subtitle, description string, status domain.Status) (*domain.Course, error)
}

type service struct {
	courseRepo domain.CourseRepository
	logger     *log.Logger
}

func (s *service) UpdateCourse(ctx context.Context, id, title, subtitle, description string, status domain.Status) (*domain.Course, error) {
	var course *domain.Course
	course.Title = title
	course.Subtitle = subtitle
	course.Description = description
	course.Status = status
	courseUpdated, err := s.courseRepo.UpdateCourse(ctx, course, id)
	if err != nil {
		s.logger.Error(err.Error())
		return &domain.Course{}, err
	}
	return courseUpdated, nil
}

func NewService(courseRepo domain.CourseRepository, logger *log.Logger) Service {

	return &service{
		courseRepo: courseRepo,
		logger:     logger,
	}

}
