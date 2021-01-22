package prisma

import (
	"Gin-Prisma-Boilerplate/db"
	"context"
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
