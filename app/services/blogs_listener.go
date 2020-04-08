package services

import (
	"net/http"

	"github.com/Victoria-engine/api-v2/app/responses"
)

// GetBlogData : Retuns the blog data
func (s *Server) GetBlogData(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusNoContent, "Not implemented")
}
