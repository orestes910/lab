package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type YamlStct struct {
	A string `yaml:"env"`
	B string `yaml:"branch"`
}

var data = `
env: dev
branch: develop
`

func main() {

	// Create the yaml struct
	yamlData := YamlStct{}

	// unmarshal the transformed data (what does []byte() do?)
	// into a pointer to the struct. Because we're pointing
	// to the struct, we only need to return an error
	err := yaml.Unmarshal([]byte(data), &yamlData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(data)
	fmt.Println(yamlData)
	fmt.Println(yamlData.A)

	output, err := yaml.Marshal(&yamlData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("\n%s", output)

}
