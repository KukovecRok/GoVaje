package MongoDB

import (
	"context"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"todorokvaja1/DataStructures"
)

func (dbo *MongoDB) InsertUser(ctx context.Context, user DataStructures.User) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("user").InsertOne(ctx, user)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}
func (dbo *MongoDB) GetUserByName(ctx context.Context, username string) (user DataStructures.User, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("user").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return
	}

	return
}
