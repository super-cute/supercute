package main

import (
	"github.com/super-cute/supercute/api"
)

func main() {
	if err := execute(); err != nil {
		panic(err)
	}
}

func execute() error {
	api := api.New()

	return api.Start()
}
