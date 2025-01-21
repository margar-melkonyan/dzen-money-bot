package app

import (
	"log"

	env "github.com/margar-melkonyan/watch-later/pkg/env-loader"
)

func init() {
	if err := env.MustLoad(); err != nil {
		log.Fatal("The application can't load .env file.")
	}
}
