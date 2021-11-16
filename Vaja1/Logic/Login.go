package Logic

import (
	"context"
	"todorokvaja1/DataStructures"
)

func (c *Controller) InsertUser(ctx context.Context, user DataStructures.User) (err error) {

	return c.db.InsertUser(ctx, user)
}

func (c *Controller) GetUserByName(ctx context.Context, username string) (user DataStructures.User, err error) {

	return c.db.GetUserByName(ctx, username)
}
func (c *Controller) Login(ctx context.Context, user DataStructures.User) (pravilno bool, err error) {
	return c.db.Login(ctx, user)
}
