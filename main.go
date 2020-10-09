package main

import (
  "fmt"
  "github.com/g5becks/easyapi/internal"
  "strings"

  "io/ioutil"
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
