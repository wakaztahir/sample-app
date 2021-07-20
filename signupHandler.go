package main

import (
	"SampleApp/validate"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type signupParams struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string `json:"token"`
}

// SignupHandler handles the requests for signups / new accounts
func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	app.configureResponse(w, r)

	var params *signupParams

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		app.errorJson(w, "Invalid signup parameters", http.StatusBadRequest)
		return
	}

	//Stripping Spaces From Parameters
	params.Name = strings.Trim(params.Name, " ")
	params.Username = strings.Trim(params.Username, " ")
	params.Email = strings.Trim(params.Email, " ")

	//todo verify recaptcha token here

	if params.Password == params.ConfirmPassword {

		//Validating Parameters
		nameVal := validate.Name(params.Name)
		usernameVal := validate.Username(params.Username)
		emailVal := validate.Email(params.Email)
		passwordVal := validate.Password(params.Password)

		if nameVal.IsValid && usernameVal.IsValid && emailVal.IsValid && passwordVal.IsValid {

			//todo Convert password to salted hash

			//Creating a user
			//user := &models.User{
			//	Name:     params.Name,
			//	Username: params.Username,
			//	Email:    params.Email,
			//}

			//todo Save User To Database
		} else {
			var validations []validate.ValidationResult
			validations = append(validations, nameVal)
			validations = append(validations, usernameVal)
			validations = append(validations, passwordVal)
			err := app.writeJson(w, "error", validations)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		app.errorJson(w, "password does not match confirm password", http.StatusBadRequest)
	}
}
