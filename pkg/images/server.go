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

	html := homePage()
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

func homePage() string {

	return `
    <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    text-align: center;
                    background-color: #f0f0f0;
                }

                form {
                    margin: 20px auto;
                    padding: 20px;
                    border: 1px solid #ccc;
                    background-color: #fff;
                    max-width: 400px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                .input-container {
                    text-align: center;
                }

                input[type="file"] {
                    display: none;
                }

                .file-upload-box {
                    border: 2px dashed #ccc;
                    padding: 20px;
                    margin: 10px 0;
                    cursor: pointer;
                }

                .file-upload-box:hover {
                    border-color: #007bff;
                }

                .file-upload-box label {
                    cursor: pointer;
                }

                select {
                    margin: 10px 0;
                    padding: 5px;
                    width: 100%;
                }

                input[type="submit"] {
                    margin: 15px 0;
                    padding: 10px 20px;
                    background-color: #007bff;
                    color: #fff;
                    border: none;
                    cursor: pointer;
                }

                input[type="submit"]:hover {
                    background-color: #0056b3;
                }
            </style>
        </head>
        <body>
            <form action="/upload" method="post" enctype="multipart/form-data">
                <div class="file-upload-box">
                    <label for="image" id="file-label">Upload Image</label>
                    <input type="file" name="image" id="image" accept="image/*" onchange="handleFileSelect(event)">
                </div>
                <br>
                <label for="transformation">Choose Transformation:</label>
                <select name="transformation" id="transformation">
                    <option value="rebaixado">Rebaixado</option>
                    <option value="pixelado">Pixelado</option>
                </select>
                <br>
                <input type="submit" value="Upload and Transform">
            </form>
            <script>
                function handleFileSelect(event) {
                    const fileLabel = document.getElementById('file-label');
                    if (event.target.files.length > 0) {
                        fileLabel.textContent = event.target.files[0].name;
                    } else {
                        fileLabel.textContent = 'Drag file here or click to upload';
                    }
                }
            </script>
        </body>
    </html>
    `

}
