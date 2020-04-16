package listeners

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Victoria-engine/api-v2/pkg/auth"
	"github.com/Victoria-engine/api-v2/pkg/presenters"
	"github.com/Victoria-engine/api-v2/pkg/repo"
	"github.com/Victoria-engine/api-v2/pkg/utils/responses"
)

// SavePost : Creates a new Post
func (server *Server) SavePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Method not allowed
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Decode the token
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		log.Fatalf("Failed to extract token from request: %v", err)
	}

	// Get the user blog
	userModel := repo.User{}
	owner, err := userModel.FindUserByID(server.DB, userID)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Post owner (user) does not exist: "+err.Error()))
		return
	}

	freshPost := repo.Post{}

	err = json.Unmarshal(body, &freshPost)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	freshPost.Prepare()

	err = freshPost.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// If the user doesnt have a blog, he cannot create posts
	if owner.BlogID <= 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("User does not belong to a blog"))
		return
	}

	freshPost.Author = *owner
	freshPost.BlogID = owner.BlogID

	_, err = freshPost.SavePost(server.DB)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error saving the post: "+err.Error()))
		return
	}

	res := presenters.SavePostPresenter(freshPost)

	responses.JSON(w, http.StatusOK, res)
}
