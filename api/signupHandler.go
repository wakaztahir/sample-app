package api

import (
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
func (server *Server) signupHandler(w http.ResponseWriter, r *http.Request) {
	server.configureResponse(w, r)

	var params *signupParams

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		errorJson(w, "Invalid signup parameters", http.StatusBadRequest)
		return
	}

	//Stripping Spaces From Parameters
	params.Name = strings.Trim(params.Name, " ")
	params.Username = strings.Trim(params.Username, " ")
	params.Email = strings.Trim(params.Email, " ")

	//Checking If Username Already Exists In Database

	//Verifying Recaptcha
	if server.config.VerifyRecaptcha {
		response, err := utils.VerifyRecaptchaToken(params.Token, server.config.RecaptchaSecret)
		if err != nil {
			log.Println("error verifying recaptcha token", err)
			errorJson(w, "error verifying recaptcha token", http.StatusInternalServerError)
			return
		}

		if !response.Success {
			errorJson(w, "invalid/unverified token", http.StatusForbidden)
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

			emailExists, err := server.handler.EmailExists(params.Email)
			if err != nil {
				log.Println("error checking if email exists", err)
				errorJson(w, "error checking if email exists", http.StatusInternalServerError)
				return
			}

			if !emailExists {

				usernameExists, err := server.handler.UsernameExists(params.Username)
				if err != nil {
					log.Println("error checking if username exists", err)
					errorJson(w, "error checking if username exists", http.StatusInternalServerError)
					return
				}

				if !usernameExists {

					password, err := utils.HashPassword(params.Password)
					if err != nil {
						log.Println("error hashing password", err)
						errorJson(w, "error hashing password", http.StatusInternalServerError)
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

					id, err := server.handler.CreateUser(user)
					if err != nil {
						log.Println("error inserting user", err)
						errorJson(w, "error inserting user", http.StatusInternalServerError)
						return
					}

					err = writeJson(w, "user", map[string]interface{}{
						"id": id,
					})

					if err != nil {
						log.Println("error writing json", err)
						return
					}

				} else {
					errorJson(w, "Username already exists", http.StatusOK)
				}
			} else {
				errorJson(w, "Email already registered", http.StatusOK)
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

			err := writeJson(w, "errors", validations)
			if err != nil {
				log.Println(err)
			}
		}
	} else {
		errorJson(w, "password does not match confirm password", http.StatusBadRequest)
		return
	}
}
