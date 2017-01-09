package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	// return Response and Error
	// treat Error as blank identifier, as I don't care about errors for now
	res, _ := http.Get("http://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
