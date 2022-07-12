package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go_test/internal/user"
	"go_test/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user. ERROR: %v", err)
	}
	d.logger.Debug("convert InsertedID to ObjectID")
	objId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objId.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex. probably ObjId: %s", objId)
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid. Hex: %s", id)
	}
	filter := bson.M{"_id": objId}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("ErrEntityNotFound")
		}
		return u, fmt.Errorf("failed to find one user by id: %s. Error: %v", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from db: %s. Error: %v", id, err)
	}
	return u, nil
}

func (d *db) FindAll(ctx context.Context) (u []user.User, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to find all users. Error: %v", err)
	}
	if err := cursor.All(ctx, &u); err != nil {
		return nil, fmt.Errorf("failed to read all documents from cursor. Erros: %v", err)
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert objectid to hex. UserId: %s", user.ID)
	}
	filter := bson.M{"_id": objId}
	userBytes, err := bson.Marshal(user)

	if err != nil {
		return fmt.Errorf("failer to marshal user. Error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes. Error: %v", err)
	}
	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query. Error: %v", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}
	d.logger.Tracef("Mached %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user id to ObjId. UserId: %s", id)
	}
	filter := bson.M{"_id": objId}
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. Error: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	d.logger.Tracef("Deleted %d documents", result.DeletedCount)
	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
