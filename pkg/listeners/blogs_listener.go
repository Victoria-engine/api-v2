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
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/Victoria-engine/api-v2/pkg/utils/responses"
)

// CreateBlog : Creates a new blog
func (server *Server) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // method not allowed
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
		log.Fatalf("failed to extract token from request: %v", err)
	}

	// Get the user blog
	userModel := repo.User{}
	owner, err := userModel.FindUserByID(server.DB, userID)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("post owner (user) does not exist: "+err.Error()))
		return
	}

	// Check if user already has a blog
	if owner.BlogID != 0 {
		responses.ERROR(w, http.StatusBadRequest, errors.New("user already has a blog"))
		return
	}

	// Create a new blog entry
	freshBlog := repo.Blog{}

	err = json.Unmarshal(body, &freshBlog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = freshBlog.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	freshBlog.Prepare()

	uuid, err := uuid.NewRandom()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	freshBlog.APIKey = uuid.String()
	freshBlog.AuthorID = owner.ID

	createdBlog, err := freshBlog.Save(server.DB)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error saving the blog: "+err.Error()))
		return
	}

	createdBlog.Author = *owner

	res := presenters.SaveBlogPresenter(createdBlog)

	responses.JSON(w, 200, res)
}

// GetBlog : Request made when accessing a blog
func (server *Server) GetBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405) // method not allowed
		return
	}

	vars := mux.Vars(r)
	apiKey := vars["apiKey"]

	if apiKey == "" {
		log.Println("missing api key in request")
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("missing api id in request"))
		return
	}

	blog := repo.Blog{}

	// Find a blog by api key
	foundBlog, err := blog.GetByAPIKey(server.DB, apiKey)
	if err != nil {
		log.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error getting the blog: "+err.Error()))
		return
	}

	res := presenters.GetBlogPresenter(foundBlog)

	responses.JSON(w, 200, res)
}
