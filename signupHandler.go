package main

import (
	"SampleApp/dbhandler"
	"SampleApp/models"
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

	//Checking If Username Already Exists In Database

	//Verifying Recaptcha
	if app.config.verifyRecaptcha {
		response, err := utils.VerifyRecaptchaToken(params.Token, app.config.recaptchaSecret)
		if err != nil {
			log.Println("error verifying recaptcha token", err)
			app.errorJson(w, "error verifying recaptcha token", http.StatusInternalServerError)
			return
		}

		if !response.Success {
			app.errorJson(w, "invalid/unverified token", http.StatusForbidden)
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

			usernameExists, err := dbhandler.UsernameExists(app.db, params.Username)
			if err != nil {
				log.Println("error checking if username exists", err)
				app.errorJson(w, "error checking if username exists", http.StatusInternalServerError)
				return
			}

			if !usernameExists {

				password, err := utils.HashPassword(params.Password)
				if err != nil {
					log.Println("error hashing password", err)
					app.errorJson(w, "error hashing password", http.StatusInternalServerError)
					return
				}

				//Creating a user
				user := &models.User{
					Name:          params.Name,
					Username:      params.Username,
					Email:         params.Email,
					Password:      password,
					EmailVerified: false,
				}

				id, err := dbhandler.CreateUser(app.db, user)
				if err != nil {
					log.Println("error inserting user", err)
					app.errorJson(w, "error inserting user", http.StatusInternalServerError)
					return
				}

				err = app.writeJson(w, "user", map[string]interface{}{
					"id": id,
				})

				if err != nil {
					log.Println("error writing json", err)
					return
				}

			} else {
				app.errorJson(w, "Username already exists", http.StatusOK)
			}
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
