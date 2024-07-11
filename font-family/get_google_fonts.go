package main

import (
	_ "embed"
	"fmt"

	"github.com/nelsonsaake/go/axios"
)

//go:embed resources/config/apikey
var apiKey string

func getGoogleFonts() ([]string, error) {

	url := "https://www.googleapis.com/webfonts/v1/webfonts?key=" + apiKey
	die := func(err error) ([]string, error) {
		return nil, err
	}

	client := axios.Client{}

	res, err := client.Get(url)
	if err != nil {
		return die(err)
	}

	resbody, err := res.Json()
	if err != nil {
		return die(err)
	}

	fonts, ok := resbody["items"].([]any)
	if !ok {
		return die(fmt.Errorf("error getting items in response body"))
	}

	fontfamilies := []string{}
	for _, v := range fonts {

		font, ok := v.(map[string]any)
		if !ok {
			return die(fmt.Errorf("error getting parsing `any` to `map[string]any`"))
		}

		fontfamilies = append(fontfamilies, font["family"].(string))
	}

	return fontfamilies, nil
}
