package transport

import (
	"encoding/json"
	"errors"
	"github.com/Victoria-engine/api-v2/pkg/api/post"
	"github.com/Victoria-engine/api-v2/pkg/api/post/repo"
	repo2 "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"github.com/Victoria-engine/api-v2/pkg/utl/jwtauth"
	"github.com/Victoria-engine/api-v2/pkg/utl/middleware/authmiddleware"
	jsonm "github.com/Victoria-engine/api-v2/pkg/utl/middleware/json"
	"github.com/Victoria-engine/api-v2/pkg/utl/responses"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP represents http service
type HTTP struct {
	// The Auth service
	svc post.Service
}

func NewHTTP(svc post.Service, mux *mux.Router) {
	h := HTTP{svc}

	postRoutes := mux.PathPrefix("/api/content").Subrouter()

	postRoutes.HandleFunc("/post",
		authmiddleware.SetMiddlewareAuthentication(
			jsonm.SetMiddlewareJSON(h.createPost),
		))
}

func (h *HTTP) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Method not allowed
		return
	}

	// Decode the token
	userID, err := jwtauth.Service{}.ExtractTokenID(r)
	if err != nil {
		log.Fatalf("Failed to extract token from request: %v", err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var p repo.PostModel

	err = json.Unmarshal(body, &p)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//TODO: Call the UserModel.getUserByID(userID)
	var owner repo2.UserModel

	freshPost, err := h.svc.SavePost(p, owner)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error saving the post: "+err.Error()))
		return
	}

	res := post.SavePostPresenter(freshPost)

	responses.JSON(w, http.StatusOK, res)

}
