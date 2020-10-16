package controllers

import "github.com/luk3skyw4lker/social-go/api/middlewares"

func (server *Server) initializeRoutes() {
	// Home
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	// Users
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuth(server.DeleteUser)).Methods("DELETE")

	// Posts
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(middlewares.SetMiddlewareJSON(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuth(server.DeletePost)).Methods("DELETE")
}
