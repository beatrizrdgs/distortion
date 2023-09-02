package images

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nfnt/resize"
)

type Server struct {
	Router  chi.Router
	Service *Service
}

func NewServer(s *Service) *Server {

	r := chi.NewRouter()

	svr := &Server{
		Router:  r,
		Service: s,
	}

	r.Mount("/", svr.Routes())

	return svr

}

func (s *Server) showUploadForm(w http.ResponseWriter, r *http.Request) {

	html := `
    <html>
        <body>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <input type="file" name="image" accept="image/*">
                <input type="submit" value="Upload and Resize">
            </form>
        </body>
    </html>
    `
	w.Write([]byte(html))

}

func (s *Server) uploadImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error parsing image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	newFilename := "rebaixado_" + handler.Filename

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	newImg := resize.Resize(uint(img.Bounds().Dx()*5), uint(img.Bounds().Dy()), img, resize.Lanczos3)

	var imgBuffer bytes.Buffer
	if err := jpeg.Encode(&imgBuffer, newImg, nil); err != nil {
		http.Error(w, "Error encoding resized image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)
	w.Header().Set("Content-Type", "image/jpeg")

	if _, err := w.Write(imgBuffer.Bytes()); err != nil {
		http.Error(w, "Error writing image to response", http.StatusInternalServerError)
		return
	}

}
