package renderer

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

var (
	renderer = render.New(render.Options{IndentJSON: true})
)

func HandleError(ctx context.Context, w io.Writer, err error) {
	log.Print(err)
	RenderError(w, http.StatusInternalServerError, err.Error())

	// Sentry Handling
	log.Print("Sentry Handling")
}

func RenderJSON(w io.Writer, status int, v interface{}) {
	renderer.JSON(w, status, v)
}

func RenderError(w io.Writer, status int, msg string) {
	renderer.JSON(w, status, newResponseError(status, msg))
}

type responseError struct {
	Status int       `json:"status"`
	Error  respError `json:"error"`
}

type respError struct {
	Message string `json:"message"`
}

func newResponseError(status int, message string) *responseError {
	return &responseError{
		Status: status,
		Error: respError{
			Message: message,
		},
	}
}
