package distortion

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type server struct {
	port    string
	service service
	router  *mux.Router
}

func NewServer(port string, svc service) *server {
	r := mux.NewRouter()

	server := &server{
		port:    port,
		service: svc,
		router:  r,
	}

	r.HandleFunc("/", server.showUploadForm).Methods("GET")
	r.HandleFunc("/upload", server.uploadImage).Methods("POST")

	return server
}

func (s *server) Run() {
	log.Printf("server running on port %s", s.port)
	if err := http.ListenAndServe(":"+s.port, s.router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func (s *server) showUploadForm(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("assets/static/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("could not read index.html: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

func (s *server) uploadImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("could not read image: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode image: %v", err), http.StatusInternalServerError)
		return
	}

	transformationType := r.FormValue("transformation")

	newImg, err := s.service.transform(img, newTransformationType(transformationType))
	if err != nil {
		http.Error(w, fmt.Sprintf("could not transform image: %v", err), http.StatusInternalServerError)
		return
	}

	buffer, err := s.encodeImage(newImg, handler)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not encode image: %v", err), http.StatusInternalServerError)
		return
	}

	s.writeImage(w, buffer, transformationType, handler)
}

func (s *server) encodeImage(newImg image.Image, handler *multipart.FileHeader) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	format := strings.Split(handler.Filename, ".")[1]
	err := s.service.encodeImage(newImg, format, buffer)
	return buffer, err
}

func (s *server) writeImage(w http.ResponseWriter, buffer *bytes.Buffer, transformationType string, handler *multipart.FileHeader) {
	newFilename := transformationType + "_" + handler.Filename
	w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)
	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, fmt.Sprintf("could not write image: %v", err), http.StatusInternalServerError)
	}
}
