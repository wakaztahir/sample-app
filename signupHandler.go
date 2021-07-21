package main

import (
	"SampleApp/utils"
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

	//Verifying Recaptcha
	if app.config.verifyRecaptcha {
		response, err := utils.VerifyRecaptchaToken(params.Token, app.config.recaptchaSecret)
		if err != nil {
			app.errorJson(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !response.Success {
			app.errorJson(w, "invalid/unverified token", http.StatusBadRequest)
			return
		}
	}

	if params.Password == params.ConfirmPassword {

		//Validating Parameters
		nameVal := validate.Name(params.Name)
		usernameVal := validate.Username(params.Username)
		emailVal := validate.Email(params.Email)
		passwordVal := validate.Password(params.Password)

		if nameVal.IsValid && usernameVal.IsValid && emailVal.IsValid && passwordVal.IsValid {

			err := app.writeJson(w, "success", map[string]bool{
				"verified-request": true,
			})
			if err != nil {
				log.Println(err)
			}
			return

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
			if !nameVal.IsValid {
				validations = append(validations, nameVal)
			}
			if !usernameVal.IsValid {
				validations = append(validations, usernameVal)
			}
			if !passwordVal.IsValid {
				validations = append(validations, passwordVal)
			}

			err := app.writeJson(w, "errors", validations)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		app.errorJson(w, "password does not match confirm password", http.StatusBadRequest)
		return
	}
}
