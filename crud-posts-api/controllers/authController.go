package controllers

import (
	"encoding/json"
	"github.com/mstfymrtc/go-posts-api/models"
	u "github.com/mstfymrtc/go-posts-api/utils"
	"net/http"
)

var Register = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := user.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.UserName, user.Password)
	u.Respond(w, resp)
}
