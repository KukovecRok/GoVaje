package MongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"todorokvaja1/DataStructures"
)

type MongoDB struct {
	Client        *mongo.Client
	User          string
	Pass          string
	IP            string
	Port          int
	Database      string
	AuthDB        string
	AuthMechanism string
}

const (
	// Timeout po 5000 MS
	connectionStringTemplate = "mongodb://%s:%s@%s:%d/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=%s&authMechanism=%s"
)

func (dbo *MongoDB) Init(ctx context.Context) (err error) {
	connectionURI := fmt.Sprintf(connectionStringTemplate, dbo.User, dbo.Pass, dbo.IP, dbo.Port, dbo.AuthDB, dbo.AuthMechanism)

	dbo.Client, err = mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		return
	}
	err = dbo.Client.Connect(ctx)
	if err != nil {

		return
	}
	err = dbo.DoInit(ctx)

	return
}

func (dbo *MongoDB) DoInit(ctx context.Context) (err error) {

	if count, err := dbo.Client.Database(dbo.Database).Collection("opravila").CountDocuments(ctx, bson.M{}); count == 0 {
		if err != nil {
			return err
		}
		_, err = dbo.Client.Database(dbo.Database).Collection("opravila").InsertOne(ctx, DataStructures.Opravilo{
			Naslov:             "Testni naslov",
			Opis:               "Opravilo dodano med DoInit()",
			PredvidenDatumDela: time.Date(2021, 11, 10, 23, 30, 33, 021, time.UTC),
			DatumDodajanja:     time.Now(),
		})
		if err != nil {
			return err
		}
	}

	err = dbo.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
