package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type RecaptchaResponse struct {
	Success    bool     `json:"success"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

func VerifyRecaptchaToken(token string, secret string) (*RecaptchaResponse, error) {

	var response = &RecaptchaResponse{
		Success: false,
	}

	if len(token) != 0 {

		//Preparing Recaptcha Request Parameters
		data := url.Values{
			"secret":   {secret},
			"response": {token},
		}

		//Sending Request To Verify Recaptcha Token
		reqResponse, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", data)
		if err != nil {
			return response, err
		}

		//Decoding Response Into Struct
		decoder := json.NewDecoder(reqResponse.Body)
		err = decoder.Decode(&response)
		if err != nil {
			return response, err
		}

		return response, nil
	} else {
		return response, nil
	}
}
