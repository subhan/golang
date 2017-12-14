package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	//data := make([]byte, 99999)
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//resp.Body.Read(data)
	//fmt.Println(string(data))
	io.Copy(os.Stdout, resp.Body)
}
