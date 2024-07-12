package wsc

import (
	"fmt"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/obj"
)

var (
	baseUrl       = "http://localhost:3004/api"
	client        *axios.Client
	loginResponse obj.Map
)

func newClient() *axios.Client {
	return &axios.Client{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", loginResponse.Get("data.access_token")),
		},
	}
}

func saveLoginResponse(v obj.Map) {
	loginResponse = v
	client = newClient()
}

func init() {
	client = newClient()
}
