package appinit

import (
	"context"
	"oms/database"
)

func Initialize(ctx context.Context) {
	initializeDb(ctx)
	initializeSqs(ctx)
}

func initializeDb(ctx context.Context) {
	database.ConnectMongo(ctx)
}

func initializeSqs(ctx context.Context) {
	database.ConnectSqs(ctx)
}
