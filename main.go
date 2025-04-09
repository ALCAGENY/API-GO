package main

import "api-go/src/server"

func main() {
	srv := server.NewServer("localhost", "8080")
	srv.Run()

}