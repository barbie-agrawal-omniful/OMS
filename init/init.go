package main

import (
	"context"
	"oms/database"
)

func main() {
	ctx := context.Background()
	database.ConnectMongo(ctx)
}
