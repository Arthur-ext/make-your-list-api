package main

import "wedding_gifts/api"

func main() {
	srv := api.NewServer()
	srv.Start()
}