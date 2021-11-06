package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorokvaja1/DataStructures"
)

type DB interface {
	Init(ctx context.Context) (err error)
	GetOpravilo(ctx context.Context) ([]DataStructures.Opravilo, error)
	GetOpraviloById(ctx context.Context, id primitive.ObjectID) (DataStructures.Opravilo, error)
	InsertOpravilo(ctx context.Context, opravilo DataStructures.Opravilo) (err error)
	RemoveOpravilo(ctx context.Context, opraviloID primitive.ObjectID) (err error)
	UpdateOpravilo(ctx context.Context, id primitive.ObjectID, opravilo DataStructures.Opravilo) (err error)
}
