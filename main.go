package main

import (
	"fmt"
	"os"
	"path"

	"github.com/tochti/go-drawer-lib"
)

const (
	GOPATH = "GOPATH"
)

func main() {
	curr, err := os.Getwd()
	handleError(err)

	oldBin := path.Join(os.Getenv(GOPATH), "bin")

	src, err := drawer.FindSrcDir(curr)
	handleError(err)

	newBin := path.Join(src, "bin")
	pathEnv := os.Getenv("PATH")
	newPath := drawer.NewPath(pathEnv, oldBin, newBin)

	fmt.Printf("export GOPATH=%v;\n", src)
	fmt.Printf("export PATH=%v;\n", newPath)

	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
