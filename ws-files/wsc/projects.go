package wsc

import "github.com/nelsonsaake/go/axios"

func GetProjects() (*axios.Response, error) {

	res, err := client.Get("projects")
	if err != nil {
		return nil, err
	}

	return res, nil
}
