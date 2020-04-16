package listeners

import (
	"net/http"

	"github.com/Victoria-engine/api-v2/pkg/utils/responses"
)

// GetBlogData : Retuns the blog data
func (server *Server) GetBlogData(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusNoContent, "Not implemented")
}
