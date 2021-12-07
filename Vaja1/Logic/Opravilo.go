package Logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todorokvaja1/DataStructures"
)

func (c *Controller) GetOpravilo(ctx context.Context) (opravilo []DataStructures.Opravilo, err error) {

	return c.db.GetOpravilo(ctx)
}
func (c *Controller) GetOpraviloById(ctx context.Context, opraviloID primitive.ObjectID) (opravilo DataStructures.Opravilo, err error) {

	return c.db.GetOpraviloById(ctx, opraviloID)
}
func (c *Controller) InsertOpravilo(ctx context.Context, opravilo DataStructures.Opravilo) (err error) {

	return c.db.InsertOpravilo(ctx, opravilo)
}
func (c *Controller) RemoveOpravilo(ctx context.Context, OpraviloID primitive.ObjectID) (err error) {

	return c.db.RemoveOpravilo(ctx, OpraviloID)
}
func (c *Controller) UpdateOpravilo(ctx context.Context, id primitive.ObjectID, opravilo DataStructures.Opravilo) (err error) {

	return c.db.UpdateOpravilo(ctx, id, opravilo)

}
