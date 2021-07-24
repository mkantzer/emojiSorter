package api

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	healthStatusOk    = "ok"
	healthStatusError = "error"
)

// swagger:model
type healthResponse struct {
	// The health of the service instance
	// Required: true
	// Example: "ok" or "error"
	Status string `json:"status"`
	// Optional message for error responses
	Message string `json:"reason,omitempty"`
}

// swagger:route GET /healthz healthStatus
//
// Returns server health status
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: body:healthResponse
func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, healthResponse{
		Status: healthStatusOk,
	})
}

// swagger:route GET /unhealthz unhealthStatus
//
// Returns a 500, useful for testing
//
//     Produces:
//     - application/json
//
//     Responses:
//       500: body:healthResponse
func (s *Server) UnhealthCheck(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, healthResponse{
		Status:  healthStatusError,
		Message: "This seems Not Fine",
	})
}
