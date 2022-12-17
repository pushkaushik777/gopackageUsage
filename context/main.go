package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeOutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	fmt.Println("Testing the context")
	//req, err := http.NewRequest(http.MethodGet, "https://placeholder.com/wp-content/uploads/2018/10/placeholder.com-logo2.jpg", nil)
	req, err := http.NewRequestWithContext(timeOutContext, http.MethodGet, "https://placeholder.com/wp-content/uploads/2018/10/placeholder.com-logo2.jpg", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imageData, err := ioutil.ReadAll(res.Body)

	fmt.Printf("File Downloaded of size %d", len(imageData))

}
