package main

import "Gin-Prisma-Boilerplate/server"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
func run() error {
	client := prisma.NewPrisma()
	server.Init()

	defer func() {
		client.Disconnect()
	}()

	return nil
}
