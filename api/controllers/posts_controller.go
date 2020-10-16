package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/luk3skyw4lker/social-go/api/auth"
	"github.com/luk3skyw4lker/social-go/api/models"
	"github.com/luk3skyw4lker/social-go/api/responses"
	formaterror "github.com/luk3skyw4lker/social-go/api/utils/errors"
)

// CreatePost is...
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)

		return
	}

	post := models.Post{}

	err = json.Unmarshal(body, &post)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)

		return
	}

	post.Prepare()

	err = post.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)

		return
	}

	uid, err := auth.ExtractTokenID(r)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Invalid token"))

		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Invalid user id"))

		return
	}

	postCreated, err := post.Create(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)

		return
	}

	responses.JSON(w, http.StatusCreated, postCreated)
}

// GetPosts is...
func (server *Server) GetPosts(res http.ResponseWriter, req *http.Request) {
	post := models.Post{}

	posts, err := post.FindAllPosts(server.DB)

	if err != nil {
		responses.ERROR(res, http.StatusInternalServerError, err)

		return
	}

	if len(*posts) == 0 {
		responses.ERROR(res, http.StatusNotFound, errors.New("No posts found"))

		return
	}

	responses.JSON(res, http.StatusOK, posts)
}

// GetPost is...
func (server *Server) GetPost(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(res, http.StatusInternalServerError, err)

		return
	}

	post := models.Post{}

	foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	responses.JSON(res, http.StatusOK, foundPost)
}

// UpdatePost is...
func (server *Server) UpdatePost(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	tokenID, err := auth.ExtractTokenID(req)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	post := models.Post{}
	foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	if tokenID != foundPost.AuthorID {
		responses.ERROR(res, http.StatusUnauthorized, errors.New("Only the author can update the post"))

		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	postUpdate := models.Post{}

	err = json.Unmarshal(body, &postUpdate)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()

	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)

		return
	}

	postUpdate.ID = foundPost.ID

	postUpdated, err := postUpdate.UpdatePost(server.DB)

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())

		responses.ERROR(res, http.StatusInternalServerError, formatedError)

		return
	}

	responses.JSON(res, http.StatusOK, postUpdated)
}

// DeletePost is...
func (server *Server) DeletePost(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)

		return
	}

	post := models.Post{}
	foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	tokenID, err := auth.ExtractTokenID(req)

	if err != nil {
		responses.ERROR(res, http.StatusInternalServerError, err)

		return
	}

	if tokenID != foundPost.AuthorID {
		responses.ERROR(res, http.StatusUnauthorized, errors.New("Only the author can delete the post"))

		return
	}

	_, err = foundPost.DeletePost(server.DB, foundPost.ID, tokenID)

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())

		responses.ERROR(res, http.StatusInternalServerError, formatedError)

		return
	}

	responses.JSON(res, http.StatusNoContent, "")
}
