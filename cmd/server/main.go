package main

import "go-backend/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}
