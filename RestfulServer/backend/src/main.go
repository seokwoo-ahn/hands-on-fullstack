package main

import (
	"hands-on/RestfulServer/backend/src/rest"
	"log"
)

func main() {
	log.Print("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
