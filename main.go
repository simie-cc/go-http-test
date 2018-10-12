package main

import (
	"errors"
	"fmt"
	"net"

	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	success, fail := 0, 0
	for i := 0; i < 20; i++ {
		err := http_test()
		if err == nil {
			success++
			fmt.Printf("o")
		} else {
			fail++
			fmt.Printf("x")
		}
	}

	fmt.Printf("  success=%d  fail=%d", success, fail)

	fmt.Println("\n")
}

func http_test() error {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout:   timeout,
		Transport: netTransport,
	}
	defer netTransport.CloseIdleConnections()

	resp, err := client.Get(os.Args[1])
	if err != nil {
		return errors.New("fail get")
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("fail read")
	}

	return nil
}
