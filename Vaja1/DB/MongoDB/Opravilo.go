package MongoDB

import (
	"context"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"todorokvaja1/DataStructures"
)

func (dbo *MongoDB) GetOpravilo(ctx context.Context) (opravila []DataStructures.Opravilo, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("opravila").Find(ctx, bson.M{})
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	opravila = make([]DataStructures.Opravilo, 0)

	err = cursor.All(ctx, &opravila)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}
	return
}

func (dbo *MongoDB) GetOpraviloById(ctx context.Context, id primitive.ObjectID) (opravilo DataStructures.Opravilo, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("opravila").FindOne(ctx, bson.M{"_id": id}).Decode(&opravilo)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}
func (dbo *MongoDB) InsertOpravilo(ctx context.Context, opravilo DataStructures.Opravilo) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("opravila").InsertOne(ctx, opravilo)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}
func (dbo *MongoDB) RemoveOpravilo(ctx context.Context, OpraviloID primitive.ObjectID) (err error) {

	_, err = dbo.Client.Database(dbo.Database).Collection("opravila").DeleteOne(ctx, bson.M{"_id": OpraviloID})
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}

func (dbo *MongoDB) UpdateOpravilo(ctx context.Context, id primitive.ObjectID, opravilo DataStructures.Opravilo) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"naslov":              opravilo.Naslov,
		"opis":                opravilo.Opis,
		"previden_datum_dela": opravilo.PredvidenDatumDela},
	}

	_, err = dbo.Client.Database(dbo.Database).Collection("opravila").UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return err
	}
	return
}
