package lib

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Error    bool     `json:"error"`
	Category string   `json:"category"`
	Type     string   `json:"type"`
	Joke     string   `json:"joke"`
	Flags    []string `json:"Flags"`
	Id       string   `json:"id"`
	Safe     bool     `json:"safe"`
	Lang     string   `json:"lang"`
}

func GetJoke() string {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any?type=single")
	if err != nil {
		print(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	//Convert the body to type string
	sb := string(body)
	var res Response
	json.Unmarshal([]byte(sb), &res)
	return res.Joke
}
