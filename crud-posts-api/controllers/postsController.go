package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mstfymrtc/go-posts-api/models"
	u "github.com/mstfymrtc/go-posts-api/utils"
	"net/http"
	"strconv"
)

var CreatePost = func(w http.ResponseWriter, r *http.Request) {

	//get user that sends request (from auth.go 74th line)
	currentUserId := r.Context().Value("user_id").(uint)
	post := &models.Post{}

	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	post.AuthorID = currentUserId
	resp := post.Create()
	u.Respond(w, resp)
}
var GetPost = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//get id from query params
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "The request is not valid"))
		return
	}
	data := models.GetPost(uint(id))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}
var DeletePost = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//get id from query params
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "The request is not valid"))
		return
	}
	resp := models.DeletePost(uint(id))
	u.Respond(w, resp)
}

var GetPosts = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetPosts()
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}
