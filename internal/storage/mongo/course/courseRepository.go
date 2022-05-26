package course

import (
	"context"
	"time"

	"github.com/amchicas/go-course-srv/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "course"
)

type courseRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) domain.CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) CreateCourse(ctx context.Context, course *domain.Course) error {
	course.Created = time.Now().Unix()
	course.Modified = time.Now().Unix()
	c := r.db.Collection(collection)
	_, err := c.InsertOne(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
func (r *courseRepository) DeleteCourseById(ctx context.Context, id string) error {
	c := r.db.Collection(collection)
	_, err := c.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
func (r *courseRepository) GetCourseById(ctx context.Context, id string) (*domain.Course, error) {
	var course domain.Course
	c := r.db.Collection(collection)
	err := c.FindOne(ctx, bson.M{"uid": id}).Decode(&course)
	if err != nil {
		return &domain.Course{}, err
	}
	return &course, nil
}
func (r *courseRepository) UpdateCourse(ctx context.Context, course *domain.Course, id string) (*domain.Course, error) {
	var courseUpdate *domain.Course
	course.Modified = time.Now().Unix()
	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(true)
	c := r.db.Collection(collection)
	err := c.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": course}, ops).Decode(courseUpdate)
	if err != nil {
		return nil, err
	}

	return courseUpdate, nil
}
