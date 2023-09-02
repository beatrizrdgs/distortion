package images

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

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
				<label for="transformation">Choose Transformation</label>
				<select name="transformation" id="transformation">
					<option value="rebaixado">Rebaixado</option>
					<option value="pixelado">Pixelado</option>
				</select>
                <input type="submit" value="Upload and Transform">
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

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	transformationType := r.FormValue("transformation")
	newFilename := transformationType + "_" + handler.Filename
	var newImg image.Image

	switch transformationType {
	case "rebaixado":
		newImg = resize.Resize(uint(img.Bounds().Dx()*5), uint(img.Bounds().Dy()), img, resize.Lanczos3)
	case "pixelado":
		pixelSize := 15
		newImg = resize.Resize(uint(img.Bounds().Dx()/pixelSize), uint(img.Bounds().Dy()/pixelSize), img, resize.NearestNeighbor)
	}

	var imgBuffer bytes.Buffer

	format := strings.Split(handler.Filename, ".")[1]
	if err := encodeImage(&imgBuffer, newImg, format); err != nil {
		http.Error(w, "Error encoding image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)

	if _, err := w.Write(imgBuffer.Bytes()); err != nil {
		http.Error(w, "Error writing image to response", http.StatusInternalServerError)
		return
	}

}

func encodeImage(imgBuffer *bytes.Buffer, img image.Image, format string) error {
	switch format {
	case "png":
		return png.Encode(imgBuffer, img)
	case "jpeg", "jpg":
		return jpeg.Encode(imgBuffer, img, nil)
	default:
		return fmt.Errorf("Unsupported image format: %s", format)
	}
}
