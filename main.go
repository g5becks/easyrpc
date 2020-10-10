package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/g5becks/easyrpc/internal"
)

func main() {
	file, _ := ioutil.ReadFile("try.varlink")
	i, err := internal.NewIDL(string(file))
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(i.Name)
	fmt.Println(len(i.Methods))
	fmt.Println(strings.Split(i.Doc, "#"))
}
