package main

import (
	"fmt"
	"go-plg/ws-files/wsc"
	"strings"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/pretty"
)

func die(_ *axios.Response, err error) {
	do.Die(err)
}

func verbose(res *axios.Response, err error) {

	do.Die(err)
	fmt.Printf("[%s] %s\n", strings.ToLower(res.Request().Method), res.Request().URL)

	obj, err := res.ObjMap()
	do.Die(err)

	fmt.Println(pretty.JSON(obj))
	fmt.Println("")
}

func main() {

	var (
		res *axios.Response
		err error
	)

	// res, err = wsc.Login()
	// die(res, err)

	res, err = wsc.CreateFile()
	verbose(res, err)

	res, err = wsc.GetOneDir()
	verbose(res, err)
}
