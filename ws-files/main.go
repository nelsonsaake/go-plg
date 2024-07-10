package main

import (
	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/pretty"
)

var (
	baseUrl = "http://localhost:3004/api"
	client  = &axios.Client{BaseUrl: baseUrl}
)

func login() {

	url := "auth/admin/login"

	redBody := map[string]any{
		"email":    "admin@example.com",
		"password": "password",
	}

	res, err := client.Post(url, redBody)
	do.Die(err)

	payload, err := res.Json()
	do.Die(err)

	pretty.Print(payload)
}

func main() {

	login()
}
