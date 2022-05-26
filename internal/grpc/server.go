package grpc

import (
	"net"

	"github.com/amchicas/go-course-srv/internal/adder"
	"github.com/amchicas/go-course-srv/internal/eraser"
	"github.com/amchicas/go-course-srv/internal/fetcher"
	"github.com/amchicas/go-course-srv/internal/modifier"
	"github.com/amchicas/go-course-srv/pkg/log"
	"github.com/amchicas/go-course-srv/pkg/pb"
	"google.golang.org/grpc"
)

type Server interface {
	Serve() error
}
type grpcServer struct {
	port   string
	aS     adder.Service
	fS     fetcher.Service
	mS     modifier.Service
	eS     eraser.Service
	logger *log.Logger
}

func NewServer(
	port string,
	aS adder.Service,
	fS fetcher.Service,
	mS modifier.Service,
	eS eraser.Service,
	logger *log.Logger,
) Server {

	return &grpcServer{port: port, aS: aS, mS: mS, fS: fS, eS: eS, logger: logger}

}
func (s *grpcServer) Serve() error {

	grpcServer := grpc.NewServer()
	srv := NewHandler(s.aS, s.mS, s.fS, s.eS, s.logger)
	pb.RegisterCourseServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		s.logger.Error("Failed at server" + err.Error())
	}

	return grpcServer.Serve(lis)

}
