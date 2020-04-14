package services

import "github.com/Victoria-engine/api-v2/pkg/utl/middleware"

func (s *Server) initializeRoutes() {
	//// Auth routes
	s.Router.HandleFunc("/api/authmiddleware/login", middleware.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/authmiddleware/register", middleware.SetMiddlewareJSON(s.Register)).Methods("POST")

	//// Content routes
	contentRoutes := s.Router.PathPrefix("/api/content").Subrouter()
	// Blog
	blogRoutes := contentRoutes.PathPrefix("/blog").Subrouter()

	blogRoutes.HandleFunc("/",
		middleware.SetMiddlewareJSON(
			middleware.SetMiddlewareAuthentication(s.GetBlogData),
		),
	).Methods("GET")

	//Post
	postRoutes := contentRoutes.PathPrefix("/post").Subrouter()

	// Posts routes
	postRoutes.HandleFunc("",
		middleware.SetMiddlewareJSON(
			middleware.SetMiddlewareAuthentication(s.SavePost),
		)).Methods("POST")
	// s.Router.HandleFunc("/posts", middleware.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middleware.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//// User routes
	userRoutes := s.Router.PathPrefix("/api/users").Subrouter()

	userRoutes.HandleFunc("/{id}",
		middleware.SetMiddlewareJSON(
			middleware.SetMiddlewareAuthentication(s.GetUserInfo),
		),
	).Methods("GET")

	userRoutes.HandleFunc("/{id}",
		middleware.SetMiddlewareJSON(
			middleware.SetMiddlewareAuthentication(s.DeleteUser),
		),
	).Methods("DELETE")

	// s.Router.HandleFunc("/users", middleware.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
}
