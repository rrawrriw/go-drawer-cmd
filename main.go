package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rrawrriw/go-venv-lib"
)

func main() {
	curr, err := os.Getwd()
	handleError(err)

	src, err := venv.FindSrcDir(curr)
	handleError(err)

	bin := path.Join(src, "bin")
	err = venv.AppendEnvList("PATH", bin)
	handleError(err)

	fmt.Printf("export GOPATH=%v;\n", src)
	fmt.Printf("export PATH=%v;", os.Getenv("PATH"))

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
