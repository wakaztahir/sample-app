package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type recaptchaResponse struct {
	Success    bool     `json:"success"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`
}

func VerifyRecaptchaToken(token string, secret string) (bool, error) {

	if len(token) != 0 {

		//Preparing Recaptcha Request Parameters
		requestBody, err := json.Marshal(map[string]string{
			"secret":   secret,
			"response": token,
		})
		bodyReader := bytes.NewBuffer(requestBody)
		if err != nil {
			return false, err
		}

		//Sending Request To Verify Recaptcha Token
		reqResponse, err := http.Post("https://www.google.com/recaptcha/api/siteverify", "application/json", bodyReader)
		if err != nil {
			return false, err
		}

		//Decoding Response Into Struct
		var response recaptchaResponse

		decoder := json.NewDecoder(reqResponse.Body)
		err = decoder.Decode(&response)
		if err != nil {
			return false, err
		}

		//Returning Response
		return response.Success, nil
	} else {
		return false, nil
	}
}
