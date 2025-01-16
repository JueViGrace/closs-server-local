package main

import (
	"fmt"

	"github.com/JueViGrace/closs-server-local/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.New()

	if err := server.Init(); err != nil {
		panic(fmt.Sprintf("cannot start server at: %s", err))
	}
}
