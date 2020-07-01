package main

import (
	"fmt"

	"github.com/astj/renovate-sandbox-go-mod-replace/module2"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println(module2.Version)
	fmt.Println(errors.New("ohno"))
}
