package wsc

import (
	"path/filepath"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
)

func GetFiles() (*axios.Response, error) {

	url := "files"
	client := newClient()

	return client.Get(url)
}

func DelFiles() (*axios.Response, error) {

	url := "files/4"
	client := newClient()

	return client.Delete(url)
}

func CreateFile() (*axios.Response, error) {

	var (
		url      = "files"
		client   = newClient()
		filename = "README.md"
	)

	raw, err := ufs.ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	req := map[string]any{
		"name":          filepath.Base(filename),
		"content":       do.RawToB64(raw),
		"parent_dir_id": 3,
	}

	return client.Post(url, req)
}

func UpdateFile() (*axios.Response, error) {

	var (
		url      = "files/5"
		client   = newClient()
		filename = "README.md"
	)

	raw, err := ufs.ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	req := map[string]any{
		"content":  do.RawToB64(raw),
		"filename": filepath.Base(filename),
	}

	return client.Patch(url, req)
}
