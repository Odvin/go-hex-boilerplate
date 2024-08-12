package main

import (
	"fmt"

	"github.com/Odvin/go-hex-boilerplate/internal/adapters/core/arithmetic"
	"github.com/Odvin/go-hex-boilerplate/internal/ports"
)

func main() {
	var core ports.ArithmaticPort

	core = arithmetic.NewAdapter()

	res, err := core.Addition(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
