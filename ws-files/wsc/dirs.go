package wsc

import (
	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/fake"
)

func CreateDir() (*axios.Response, error) {

	req := map[string]any{
		"name":          fake.Word(),
		"parent_dir_id": 3,
	}

	return client.Post("dirs", req)
}

func GetOneDir() (*axios.Response, error) {

	return client.Get("dirs/3")
}

func GetDirs() (*axios.Response, error) {

	return client.Get("dirs")
}

func UpdateDir() (*axios.Response, error) {

	req := map[string]any{
		"name": fake.Word(),
	}

	return client.Patch("dirs/1", req)
}

func DeleteDir() (*axios.Response, error) {

	return client.Delete("dirs/2")
}
