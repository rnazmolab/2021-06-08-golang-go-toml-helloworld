package main

import (
	"fmt"

	"github.com/pelletier/go-toml/v2"
)

type MyConfig struct {
	Version       int
	Name          string
	Tags          []string
	FooBarVersion string `toml:"foo_bar_version"`
}

// Ref: https://pkg.go.dev/github.com/pelletier/go-toml/v2#Marshal
func marshal(cfg MyConfig) ([]byte, error) {
	return toml.Marshal(cfg)
}

// Ref: https://pkg.go.dev/github.com/pelletier/go-toml/v2#Unmarshal
func unmarshal(b []byte, cfg *MyConfig) error {
	return toml.Unmarshal(b, cfg)
}

func main() {
	fmt.Println("=================")

	cfg1 := MyConfig{
		Version:       2,
		Name:          "go-toml",
		Tags:          []string{"go", "toml"},
		FooBarVersion: "0.0.4",
	}
	b, err := marshal(cfg1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	fmt.Println("=================")

	doc := []byte(`
	version = 2
	name = "go-toml"
	tags = ["go", "toml"]
	foo_bar_version = "0.0.4"
	`)

	var cfg2 MyConfig
	err = unmarshal(doc, &cfg2)
	if err != nil {
		panic(err)
	}
	fmt.Println("version:", cfg2.Version)
	fmt.Println("name:", cfg2.Name)
	fmt.Println("tags:", cfg2.Tags)
	fmt.Println("foo_bar_version:", cfg2.FooBarVersion)

	fmt.Println("=================")
}
