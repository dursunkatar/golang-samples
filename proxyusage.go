package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
    // proxy ip and port
	os.Setenv("HTTP_PROXY", "11.111.111.11:8080")
	resp, err := http.Get("https://target.com/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
