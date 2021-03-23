package main

import "clientApi/api"

func main() {
	api.InitRepository()
	api.StartServer()
}
