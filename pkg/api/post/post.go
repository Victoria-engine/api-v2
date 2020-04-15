package post

import (
	"errors"
	"github.com/Victoria-engine/api-v2/pkg/api/post/repo"
	ur "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
)

// SavePost : Creates a new Post
func (p Post) SavePost(post repo.PostModel, owner ur.UserModel) (repo.PostModel, error) {

	/*	//TODO: Move to different method (User Model)
		owner, err := userModel.FindUserByID(server.DB, userID)
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Post owner (user) does not exist: "+err.Error()))
			return
		}
	*/
	post.Prepare()

	err := post.Validate()
	if err != nil {
		return post, err
	}

	// If the user doesnt have a blog, he cannot create posts
	if post.BlogID <= 0 {
		return post, errors.New("user does not belong to a blog")
	}

	post.Author = owner
	post.BlogID = owner.BlogID

	return post, nil
}
