package main

import (
	"log/slog"

	"github.com/adrien3d/fizzbuzz/server"
)

func main() {
	err := server.SetupRouter()
	if err != nil {
		slog.Error("unable to setup router")
	}
}
