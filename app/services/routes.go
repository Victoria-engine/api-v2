package services

import "github.com/Victoria-engine/api-v2/app/middlewares"

func (s *Server) initializeRoutes() {
	//// Auth routes
	s.Router.HandleFunc("/api/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/auth/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

	//// Content routes
	contentRoutes := s.Router.PathPrefix("/api/content").Subrouter()
	// Blog
	blogRoutes := contentRoutes.PathPrefix("/blog").Subrouter()

	blogRoutes.HandleFunc("/",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.GetBlogData),
		),
	).Methods("GET")

	// Post
	//postRoutes := contentRoutes.PathPrefix("/post").Subrouter()

	// Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//// User routes
	userRoutes := s.Router.PathPrefix("/api/users").Subrouter()

	userRoutes.HandleFunc("/{id}",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.GetUserInfo),
		),
	).Methods("GET")

	userRoutes.HandleFunc("/{id}",
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(s.DeleteUser),
		),
	).Methods("DELETE")

	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
}
