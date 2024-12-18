package main

import "api-mock/internal/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
