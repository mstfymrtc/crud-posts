package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mstfymrtc/go-posts-api/app"
	"github.com/mstfymrtc/go-posts-api/controllers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("/api/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/authenticate", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Content-Type",
			"Authorization",
		},
		Debug: true, // Enable Debugging for testing, consider disabling in production
	}).Handler(router)

	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		fmt.Println(err)
	}
}
