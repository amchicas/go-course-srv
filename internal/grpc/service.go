package grpc

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/adder"
	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/internal/eraser"
	"github.com/amchicas/go-course-srv/internal/fetcher"
	"github.com/amchicas/go-course-srv/internal/modifier"
	"github.com/amchicas/go-course-srv/pkg/log"
	"github.com/amchicas/go-course-srv/pkg/pb"
)

type courseHandler struct {
	aS     adder.Service
	fS     fetcher.Service
	mS     modifier.Service
	eS     eraser.Service
	logger *log.Logger
}

func NewHandler(adderService adder.Service, modifieService modifier.Service, fetcherService fetcher.Service, eraserService eraser.Service, logger *log.Logger) pb.CourseServiceServer {

	return &courseHandler{
		aS:     adderService,
		mS:     modifieService,
		fS:     fetcherService,
		eS:     eraserService,
		logger: logger,
	}
}
func (s *courseHandler) CreateCourse(ctx context.Context, req *pb.CourseReq) (*pb.CreateCourseResp, error) {

	course, err := s.aS.AddCourse(ctx,
		req.Course.Title,
		req.Course.Subtitle,
		req.Course.Description,
		domain.Status(req.Course.Status),
	)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.CreateCourseResp{}, err
	}
	return &pb.CreateCourseResp{CourseId: course.Id}, nil

}

func (s *courseHandler) UpdateCourse(ctx context.Context, req *pb.CourseReq) (*pb.UpdateCourseResp, error) {
	course, err := s.fS.GetCourse(ctx, req.Course.Id)
	if err != nil {

		return &pb.UpdateCourseResp{}, err
	}

	course.Title = req.Course.Title
	course.Subtitle = req.Course.Subtitle
	course.Description = req.Course.Description

	courseUpdated, err := s.mS.UpdateCourse(ctx,
		req.Course.Id,
		course.Title,
		course.Subtitle,
		course.Description,
		domain.Status(req.Course.Status),
	)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.UpdateCourseResp{}, err

	}
	c := &pb.Course{
		Id:          courseUpdated.Id,
		Title:       courseUpdated.Title,
		Subtitle:    courseUpdated.Subtitle,
		Description: courseUpdated.Description,
	}
	return &pb.UpdateCourseResp{Course: c}, nil
}
func (s *courseHandler) FindCourse(ctx context.Context, req *pb.FindCourseReq) (*pb.CourseResp, error) {
	course, err := s.fS.GetCourse(ctx, req.CourseId)
	if err != nil {

		return &pb.CourseResp{}, err
	}

	return &pb.CourseResp{
		Course: &pb.Course{
			Id:          course.Id,
			Title:       course.Title,
			Subtitle:    course.Subtitle,
			Description: course.Description,
			Created:     course.Created,
			Modified:    course.Modified,
		},
	}, nil
}
func (s *courseHandler) DeleteCourse(ctx context.Context, req *pb.DeleteCourseReq) (*pb.DeleteCourseResp, error) {
	err := s.eS.DeleteCourseById(ctx, req.CourseId)
	if err != nil {

		return &pb.DeleteCourseResp{}, err
	}

	return &pb.DeleteCourseResp{}, nil
}
