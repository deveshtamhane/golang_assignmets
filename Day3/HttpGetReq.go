package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	url:="http://golang.org/"
	getData(url)
}

func getData(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", string(contents))
		fmt.Println("Response Code: ",response.StatusCode)
	}
}
