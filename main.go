package main

import (
	"image-messer/pkg/images"
	"net/http"
)

func main() {

	svc := images.NewService()
	s := images.NewServer(svc)

	http.ListenAndServe(":8080", s.Router)

}
