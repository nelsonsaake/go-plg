package main

import (
	"fmt"
	"go-plg/ws-files/wsc"
	"strings"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
)

func die(_ *axios.Response, err error) {
	do.Die(err)
}

func verbose(res *axios.Response, err error) {

	// first handle error
	do.Die(err)

	// print som info about the request
	fmt.Printf("[%s] %s\n", strings.ToLower(res.Request().Method), res.Request().URL)

	// check if request was successful, i.e interpret status code
	fmt.Printf("status: %s\n", res.Response().Status)

	// extract information about the response
	obj, err := res.ObjMap()
	do.Die(err)

	// print response
	fmt.Println(obj)
}

func main() {

	var (
		res *axios.Response
		err error
	)

	res, err = wsc.Login()
	die(res, err)

	/*
		examples:

		// res, err = wsc.CreateFile()
		// verbose(res, err)

		// obj, err := res.ObjMap()
		// do.Die(err)

		// pretty.Print(obj.Get("data"))
	*/

	res, err = wsc.GetReviews()
	die(res, err)

	res, err = wsc.GetReview()
	die(res, err)

	res, err = wsc.CreateReview()
	die(res, err)

	res, err = wsc.UpdateReview()
	verbose(res, err)

	res, err = wsc.Delete()
	die(res, err)
}
