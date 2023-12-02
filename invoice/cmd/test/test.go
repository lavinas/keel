package main

import (
	"fmt"

	"github.com/lavinas/keel/invoice/pkg/kerror"
)

func main() {
	err := kerror.NewKError(kerror.BadRequest, "client id already exists")
	err.SetPrefix("client ")
	fmt.Println(err)
}
