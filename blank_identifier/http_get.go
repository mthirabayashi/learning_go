package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

// Go to shareagram.us and output all the html to stdout
func main() {
  res, _ := http.Get("http://www.shareagram.us")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", page)
}
