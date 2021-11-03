package Logic

import (
	"context"
	"todorokvaja1/DataStructures"
)

func (c *Controller) GetOpravilo(ctx context.Context) (opravilo []DataStructures.Opravilo, err error) {
	return c.db.GetOpravilo(ctx)
}
