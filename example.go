// +build run

package main

import (
	"os"

	"github.com/zetamatta/go-windows-junction"
)

func main() {
	if len(os.Args) < 3 {
		println("go run example.go TARGET MOUNTPT")
		return
	}
	if err := junction.Create(os.Args[1], os.Args[2]); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
