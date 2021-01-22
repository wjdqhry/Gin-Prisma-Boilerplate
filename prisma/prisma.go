package prisma

import (
	"context"
	"go-prisma/db"
)

var Client *db.PrismaClient
var Ctx context.Context

func NewPrisma() *db.PrismaClient {
	Client = db.NewClient()
	err := Client.Connect()
	if err != nil {
		panic(err)
	}
	Ctx = context.Background()
	return Client
}
