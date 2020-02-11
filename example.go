// +build run

package main

import (
	"errors"
	"os"

	"github.com/zetamatta/go-windows-junction"
)

func main1() error {
	if len(os.Args) < 3 {
		return errors.New("go run example.go TARGET MOUNTPT")
	}
	if err := junction.Create(os.Args[1], os.Args[2]); err != nil {
		return err
	}
	result, err := os.Readlink(os.Args[2])
	if err != nil {
		return err
	}
	println(os.Args[2], " is linked to ", result)
	return nil
}

func main() {
	if err := main1(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
