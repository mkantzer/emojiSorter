package api

import (
	"net/http"

	"github.com/go-chi/render"
)

const helloWorld = "Hello World!\n"

// swagger:route GET / Hello
//
// Returns Hello World
//
// Produces:
//   - plain/text
//
// Responses:
//   200: description:Body contains "Hello World!"
func (s *Server) HelloServer(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(s.Deps.Logger, r)
	logger.Info("Hello World!")

	render.Status(r, http.StatusOK)
	render.PlainText(w, r, helloWorld)
}
