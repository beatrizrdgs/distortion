package images

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Server struct {
	ListenAddr string
	Service    Service
}

func NewServer(listenAddr string, svc Service) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Service:    svc,
	}
}

func (s *Server) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", s.showUploadForm).Methods("GET")
	r.HandleFunc("/upload", s.uploadImage).Methods("POST")

	log.Println("Starting server on", s.ListenAddr)

	if err := http.ListenAndServe(s.ListenAddr, r); err != nil {
		log.Fatal("error starting server: ", err)
	}
}

func (s *Server) showUploadForm(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("assets/static/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("could not read index.html: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(html))
}

func (s *Server) uploadImage(w http.ResponseWriter, r *http.Request) {
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

	newImg, err := s.Service.Transform(img, NewTransformationType(transformationType))
	if err != nil {
		http.Error(w, fmt.Sprintf("could not transform image: %v", err), http.StatusInternalServerError)
		return
	}

	newFilename := transformationType + "_" + handler.Filename

	buffer := new(bytes.Buffer)

	format := strings.Split(handler.Filename, ".")[1]
	if err := s.Service.EncodeImage(newImg, format, buffer); err != nil {
		http.Error(w, fmt.Sprintf("could not encode image: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)

	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, fmt.Sprintf("could not write image: %v", err), http.StatusInternalServerError)
		return
	}
}
