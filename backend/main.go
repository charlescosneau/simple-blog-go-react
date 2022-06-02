package main

import (
	"backend/router"
	_ "github.com/lib/pq"
)

func main() {
	router.HandleRequest()
}
