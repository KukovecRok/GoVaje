package Logic

import (
	"context"
)

func (c *Controller) GetHealth(ctx context.Context) (zdravje string, err error) {

	return c.db.GetHealth(ctx)

}
