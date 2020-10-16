package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/luk3skyw4lker/social-go/api/auth"
	"github.com/luk3skyw4lker/social-go/api/models"
	"github.com/luk3skyw4lker/social-go/api/responses"
	formaterror "github.com/luk3skyw4lker/social-go/api/utils/errors"
)

// Login is...
func (server *Server) Login(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	user.Prepare()
	err = user.Validate("login")

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)

		return
	}

	token, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(res, http.StatusBadRequest, formattedError)

		return
	}

	responses.JSON(res, http.StatusOK, token)
}

// SignIn is...
func (server *Server) SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
