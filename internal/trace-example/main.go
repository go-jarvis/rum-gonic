package main

import (
	"github.com/go-jarvis/rum-gonic/rum"
)

func main() {
	r := rum.Default()

	r.Register(&Index{})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
