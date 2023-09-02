package utils

import (
	"image-messer/pkg/cerrors"
	"net/http"

	"github.com/go-chi/render"
)

func HandleError(w http.ResponseWriter, r *http.Request, err error) error {

	switch cerr := err.(type) {
	case cerrors.CError:
		return render.Render(w, r, cerr)
	}

	cerr := cerrors.CError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	return render.Render(w, r, cerr)

}
