package Logic

import (
	"os"
	"todorokvaja1/DB"
)

type Controller struct {
	db     DB.DB
	secret []byte
}

func NewController(db DB.DB, secret []byte) *Controller {
	return &Controller{
		db:     db,
		secret: secret,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
