package appinit

import (
	"context"
	"oms/database"
)

func Initialize(ctx context.Context) {
	initializeDb(ctx)
}

func initializeDb(ctx context.Context) {
	database.ConnectMongo(ctx)
}
