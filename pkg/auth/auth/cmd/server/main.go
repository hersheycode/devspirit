package main

import (
	"log"

	handlers "apppathway.com/pkg/user/auth/internals/submod/app/server/handlers"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	handlers.Listen()
}
