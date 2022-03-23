package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// cancel due to timeout
func main() {

	// Create a new request
	req, _ := http.NewRequest("GET", "https://www.google.com", nil)

	// Create a context with a timeout of 5 seconds.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// Bind the new context into the request.
	req = req.WithContext(ctx)

	// Make the web call and return any error. Do will handle the context level timeout.
	res, err := http.DefaultClient.Do(req)

	//client := &http.Client{}
	//res, err := client.Do(req)

	if err != nil {
		fmt.Println("request error", err)
		return
	}
	// Close the response body on the return.
	defer res.Body.Close()
	fmt.Println("response status code", res.StatusCode)

	// Write the response to stdout.
	io.Copy(os.Stdout, res.Body)
}
