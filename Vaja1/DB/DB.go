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

	GetHealth(ctx context.Context) (zdravje string, err error)
	InsertUser(ctx context.Context, user DataStructures.User) (err error)
	GetUserByName(ctx context.Context, username string) (user DataStructures.User, err error)
}
