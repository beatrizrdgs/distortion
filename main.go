package main

import (
	"distortion/internal/images"
)

func main() {
	svc := images.NewService()
	svr := images.NewServer(":8080", svc)

	svr.Run()
}
