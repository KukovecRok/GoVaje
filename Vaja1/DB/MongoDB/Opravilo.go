package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"todorokvaja1/DataStructures"
)

func (dbo *MongoDB) GetOpravilo(ctx context.Context) (opravila []DataStructures.Opravilo, err error) {
	cursor, err := dbo.Client.Database(dbo.Database).Collection("opravila").Find(ctx, bson.M{})
	if err != nil {
		return
	}

	opravila = make([]DataStructures.Opravilo, 0)

	err = cursor.All(ctx, &opravila)
	if err != nil {
		return
	}
	return
}
