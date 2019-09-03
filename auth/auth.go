package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

type Credential struct {
	Email    string `json:"e_mail"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

const authUrl string = ""

func Authenticate(email string, password string) (string, error) {

	body, err := json.Marshal(&Credential{
		email,
		password,
	})

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPut, authUrl, strings.NewReader(string(body)))
	request.Header.Add("content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var token Token
	json.Unmarshal([]byte(content), &token)
	return response.Header.Get("set-cookie") + ";" + " accessToken=" + token.Token, nil
}
