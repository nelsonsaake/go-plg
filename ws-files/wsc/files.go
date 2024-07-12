package wsc

import (
	"path/filepath"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/obj"
	"github.com/nelsonsaake/go/ufs"
)

func GetFiles() (obj.Map, error) {

	url := "files"
	client := newClient()

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return res.ObjMap()
}

func DelFiles() (obj.Map, error) {

	url := "files/4"
	client := newClient()

	res, err := client.Delete(url)
	if err != nil {
		return nil, err
	}

	return res.ObjMap()
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

	req := obj.Map{
		"name":          filepath.Base(filename),
		"content":       do.RawToB64(raw),
		"parent_dir_id": 3,
	}

	return client.Post(url, req)
}

func UpdateFile() (obj.Map, error) {

	var (
		url      = "files/5"
		client   = newClient()
		filename = "README.md"
	)

	raw, err := ufs.ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	req := obj.Map{
		"content":  do.RawToB64(raw),
		"filename": filepath.Base(filename),
	}

	res, err := client.Patch(url, req)
	if err != nil {
		return nil, err
	}

	return res.ObjMap()
}
