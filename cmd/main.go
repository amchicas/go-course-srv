package main

import (
	"context"

	"github.com/amchicas/go-course-srv/config"
	"github.com/amchicas/go-course-srv/internal/adder"
	"github.com/amchicas/go-course-srv/internal/domain"
	"github.com/amchicas/go-course-srv/internal/eraser"
	"github.com/amchicas/go-course-srv/internal/fetcher"
	"github.com/amchicas/go-course-srv/internal/grpc"
	"github.com/amchicas/go-course-srv/internal/modifier"
	"github.com/amchicas/go-course-srv/internal/storage/mongo"
	"github.com/amchicas/go-course-srv/internal/storage/mongo/course"
	"github.com/amchicas/go-course-srv/pkg/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := log.New("Course", "dev")
	c, err := config.LoadConfig()
	if err != nil {

		logger.Error("Failed at config" + err.Error())
	}
	repo := newMongo(c.MongoHost, c.MongoPort, c.Database)
	adderService := adder.NewService(repo, logger)
	fetcherService := fetcher.NewService(repo, logger)
	modifierService := modifier.NewService(repo, logger)
	eraserService := eraser.NewService(repo, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		srv := grpc.NewServer(c.Port, adderService, fetcherService, modifierService, eraserService, logger)
		return srv.Serve()

	})

	logger.Fatal(g.Wait().Error())

}

func newMongo(host, port, database string) domain.CourseRepository {
	db, cancel := mongo.NewConn(host, port, database)
	defer cancel()
	return course.NewMongo(db)
}
