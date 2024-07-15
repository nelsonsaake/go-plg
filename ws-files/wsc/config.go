package wsc

import (
	"fmt"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/obj"
)

var (
	baseUrl       = "http://localhost:3004/api"
	client        *axios.Client
	loginResponse *obj.Map
)

func AccessToken() any {
	if loginResponse != nil {
		return loginResponse.Get("data.access_token")
	}
	return ""
}

func newClient() *axios.Client {
	return &axios.Client{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", AccessToken()),
		},
	}
}

func saveLoginResponse(v *obj.Map) {
	loginResponse = v
	client = newClient()
}

func init() {
	client = newClient()
}
