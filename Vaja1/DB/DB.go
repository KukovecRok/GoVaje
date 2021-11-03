package DB

import (
	"context"
	"todorokvaja1/DataStructures"
)

type DB interface {
	Init(ctx context.Context) (err error)
	GetOpravilo(ctx context.Context) (opravilo []DataStructures.Opravilo, err error)
}
