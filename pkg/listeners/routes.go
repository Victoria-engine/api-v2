package listeners

import "github.com/Victoria-engine/api-v2/pkg/middlewares"

func (server *Server) initializeRoutes() {
	//// Auth routes
	server.Router.HandleFunc("/api/auth/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")
	server.Router.HandleFunc("/api/auth/register", middlewares.SetMiddlewareJSON(server.Register)).Methods("POST")

	//// Content routes
	contentRoutes := server.Router.PathPrefix("/api/content").Subrouter()
	// Blog
	blogRoutes := contentRoutes.PathPrefix("/blog").Subrouter()

	blogRoutes.HandleFunc("",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(server.GetBlogData),
		),
	).Methods("GET")
	blogRoutes.HandleFunc("",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(server.CreateBlog),
		),
	).Methods("POST")

	//Post
	postRoutes := contentRoutes.PathPrefix("/post").Subrouter()

	// Posts routes
	postRoutes.HandleFunc("",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(server.SavePost),
		)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//// User routes
	userRoutes := server.Router.PathPrefix("/api/users").Subrouter()

	userRoutes.HandleFunc("/{id}",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(server.GetUserInfo),
		),
	).Methods("GET")

	userRoutes.HandleFunc("/{id}",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(server.DeleteUser),
		),
	).Methods("DELETE")

	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
}
