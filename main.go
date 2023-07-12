package main

import (
	"fmt"

	"github.com/S5477/go-redis/route"
)

const PORT = ":8000"

func main() {
	serveApplication()
}

func serveApplication() {
	router := route.Route()
	router.Run(PORT)
	fmt.Println("Server running on port" + PORT)
}
