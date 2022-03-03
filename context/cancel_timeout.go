package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// cancel due to timeout
func main() {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("request error", err)
		return
	}
	fmt.Println("response status code", res.StatusCode)
}
