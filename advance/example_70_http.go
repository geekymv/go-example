package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func sendHttp(url string) (string, error) {
	// 发送 http 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s:%s", url, resp.Status)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s:%s", url, resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return string(content), fmt.Errorf("%s:%s", url, err.Error())
	}

	return string(content), nil
}

func main() {

	url := "http://www.baidu.com1"

	if content, err := sendHttp(url); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	} else {

		fmt.Println(content)
	}

}
