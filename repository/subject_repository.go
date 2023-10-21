package repository

import (
	"context"

	"github.com/Bicycle-Guild/kaktys/helpers"
	"github.com/Bicycle-Guild/kaktys/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client = helpers.ConnectDB()

func Create_subject(new_subject models.CreateSubject) models.Subject {
	var inserted_subject models.Subject

	subject_collection := db.Database("kaktus").Collection("subject")
	result, _ := subject_collection.InsertOne(context.TODO(), new_subject)

	subject_collection.FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: result.InsertedID}},
	).Decode(&inserted_subject)

	return inserted_subject
}

func Get_subjects() []models.Subject {
	var results []models.Subject

	subject_collection := db.Database("kaktus").Collection("subject")
	cursor, _ := subject_collection.Find(context.TODO(), bson.D{})
	cursor.All(context.TODO(), &results)

	return results
}

func Get_Subject(subjectID string) (models.Subject, error) {
	var result models.Subject

	id, _ := primitive.ObjectIDFromHex(subjectID)

	subject_collection := db.Database("kaktus").Collection("subject")
	err := subject_collection.FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: id}},
	).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Subject{}, err
		}
	}

	return result, nil
}
