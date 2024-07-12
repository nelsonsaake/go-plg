package wsc

import "github.com/nelsonsaake/go/axios"

func Login() (*axios.Response, error) {

	var (
		url = "auth/admin/login"
		req = map[string]any{
			"email":    "admin@example.com",
			"password": "password",
		}
	)

	res, err := client.Post(url, req)
	if err != nil {
		return nil, err
	}

	obj, err := res.ObjMap()
	if err != nil {
		return nil, err
	}

	// for subsequent auth
	saveLoginResponse(obj)

	return res, nil
}
