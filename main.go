package main

import (
	"painellembretes/config"
	"painellembretes/routes"
)

func main() {
	config.LoadEnvFile()

	routes.SetRoutes()
}
