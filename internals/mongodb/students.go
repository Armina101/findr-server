package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thebravebyte/findr/internals"
)

var collection = "students"

// Todo: Implement the methods of the LinkerStore interface here that has to do with the students collection.

func (fd *FindrDB) CreateAccount(std internals.Students) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var res bson.M

	filter := bson.D{{Key: "email", Value: std.Email}}
	err := FindrData(fd.DB, collection).FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			std.ID = primitive.NewObjectID().Hex()
			_, err := FindrData(fd.DB, collection).InsertOne(ctx, std)
			if err != nil {
				fd.Logger.Error(fmt.Sprintf("Create Student Accout: %v", err.Error()))
				return false, false, errors.New("cannot registered this account")
			}
			return true, false, nil
		}
		fd.Logger.Error(err.Error())
		return false, false, errors.New("cannot registered this account")
	}
	return true, true, nil
}

// UpdateInfo this method helps to make the necessary update in the database specifically
// the user document
func (fd *FindrDB) UpdateDetails(id, token string) (bool, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	filter := bson.D{{Key: "_id", Value: id}}
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "jwt_token", Value: token},
		{Key: "updated_at", Value: updateTime},
	}}}
	_, err := FindrData(fd.DB, collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	fd.Logger.Debug("UpdateInfo Student: updating user details")

	return true, nil
}

// VerifyCode this method is implemented to verify user login credentials with respect
// to existing data
func (fd *FindrDB) VerifyDetails(email string) (internals.Students, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	var res internals.Students

	filter := bson.D{{Key: "email", Value: email}}
	err := FindrData(fd.DB, collection).FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fd.Logger.Error(fmt.Sprintf("VerifyCode: Error %v", err.Error()))
			return internals.Students{}, errors.New("unregistered account! sign up")
		}
		fd.Logger.Error(err.Error())
		return internals.Students{}, err
	}
	return res, nil
}
