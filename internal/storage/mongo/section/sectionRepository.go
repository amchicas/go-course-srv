package mongo

import (
	"context"

	"github.com/amchicas/go-course-srv/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sectionRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) domain.SectionRepository {
	return &sectionRepository{db: db}
}

func (r *sectionRepository) SaveSection(ctx context.Context, idCourse string, section *domain.Section) error {
	c := r.db.Collection("course")
	_, err := c.InsertOne(ctx, section)
	if err != nil {
		return err
	}

	return nil
}
func (r *sectionRepository) DeleteSectionByID(ctx context.Context, idCourse, idSection string) error {
	c := r.db.Collection("course")
	_, err := c.UpdateOne(ctx, bson.M{"id": idCourse},
		bson.M{
			"$pull": bson.M{
				"id": idSection,
			}},
	)
	if err != nil {
		return err
	}
	return nil
}
func (r *sectionRepository) GetSectionByID(ctx context.Context, idSection string) (*domain.Section, error) {
	var section domain.Section
	c := r.db.Collection("course")
	filter := bson.M{"section": bson.M{"id": idSection}}
	err := c.FindOne(ctx, filter).Decode(&section)
	if err != nil {
		return &domain.Section{}, err
	}
	return &section, nil
}
func (r *sectionRepository) UpdateSection(ctx context.Context, section *domain.Section, idCourse, idSection string) error {
	c := r.db.Collection("section")
	filter := bson.M{"id": idSection}
	update := bson.M{
		"$set": bson.M{
			"title":       section.Title,
			"description": section.Description,
			"status":      section.Status,
			"created":     section.Created,
			"modified":    section.Modified,
		}}
	_, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
