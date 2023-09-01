package images

import (
	"image"
	"io"
	"net/http"
	"os"

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

	err = s.Service.Save(newFilename, newImg)
	if err != nil {
		http.Error(w, "Error saving resized image", http.StatusInternalServerError)
		return
	}

	downloadURL := "/download/" + newFilename
	response := `
        <html>
        	Imagem rebaixada com sucesso.
        	<a id="downloadButton" href="` + downloadURL + `">
            	<button type="button">Download Resized Image</button>
        	</a>
        </html>
		<script>
    		document.getElementById("downloadButton").addEventListener("click", function(event) {
			event.preventDefault();
			var downloadLink = this.getAttribute("href");
			var link = document.createElement("a");
			link.href = downloadLink;
			link.download = "` + newFilename + `"; // Set the desired filename here
			link.style.display = "none";
			document.body.appendChild(link);
			link.click();
			document.body.removeChild(link);
		});
		</script>
    `

	w.Write([]byte(response))

}

func (s *Server) downloadImage(w http.ResponseWriter, r *http.Request) {

	filename := chi.URLParam(r, "filename")
	path := "./" + filename

	imgFile, err := os.Open(path)
	if err != nil {
		http.Error(w, "Error opening image file", http.StatusNotFound)
		return
	}
	defer imgFile.Close()

	w.Header().Set("Content-Type", "image/jpeg")

	_, err = io.Copy(w, imgFile)
	if err != nil {
		http.Error(w, "Error copying image to response", http.StatusInternalServerError)
		return
	}

}
