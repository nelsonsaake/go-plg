package main

import (
	_ "embed"
	"fmt"

	"github.com/nelsonsaake/go/do"
	"gopkg.in/yaml.v3"
)

//go:embed data.yaml
var data string

func main() {

	config := map[string]any{}

	err := yaml.Unmarshal([]byte(data), &config)
	do.Die(err)

	// deps, _ := config["dependencies"].(map[string]any)
	// for k, v := range deps {

	// 	_, ok := v.(string)
	// 	if ok {
	// 		fmt.Printf("flutter pub add %v\n", k)
	// 	}
	// }

	devdeps, _ := config["dev_dependencies"].(map[string]any)
	for k, v := range devdeps {

		_, ok := v.(string)
		if ok {
			fmt.Printf("flutter pub add dev:%v\n", k)
		}
	}
}
