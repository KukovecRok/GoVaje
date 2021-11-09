package MongoDB

import (
	"context"
	"github.com/getsentry/sentry-go"
	"log"
)

func (dbo *MongoDB) GetHealth(ctx context.Context) (zdravje string, err error) {
	zdravje = "zdrav"
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}
