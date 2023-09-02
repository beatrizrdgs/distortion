package images

import (
	"bytes"
	"image"
	"image-messer/pkg/cerrors"
	"image-messer/pkg/utils"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type Server struct {
	Router  chi.Router
	Service Handler
}

func NewServer(s Handler) *Server {

	svr := &Server{
		Router:  chi.NewRouter(),
		Service: s,
	}

	svr.Router.Get("/", svr.showUploadForm)
	svr.Router.Post("/upload", svr.uploadImage)

	return svr

}

func (s *Server) Routes() chi.Router {

	return s.Router

}

func (s *Server) showUploadForm(w http.ResponseWriter, r *http.Request) {

	html := homePage()
	w.Write([]byte(html))

}

var errWritingResponse = cerrors.NewCError("error writing image to response", http.StatusInternalServerError)

func (s *Server) uploadImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.HandleError(w, r, ErrParsingImage)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		utils.HandleError(w, r, ErrDecodeImage)
		return
	}

	transformationType := r.FormValue("transformation")

	newImg, err := s.Service.Transform(img, NewTransformationType(transformationType))

	newFilename := transformationType + "_" + handler.Filename

	buffer := new(bytes.Buffer)

	format := strings.Split(handler.Filename, ".")[1]
	if err := s.Service.EncodeImage(newImg, format, buffer); err != nil {
		utils.HandleError(w, r, err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+newFilename)

	if _, err := w.Write(buffer.Bytes()); err != nil {
		utils.HandleError(w, r, errWritingResponse)
		return
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
