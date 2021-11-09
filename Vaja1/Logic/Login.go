package Logic

import (
	"context"
	"todorokvaja1/DataStructures"
)

func (c *Controller) InsertUser(ctx context.Context, user DataStructures.User) (err error) {

	return c.db.InsertUser(ctx, user)
}
