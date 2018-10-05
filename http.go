package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

// RequestParameters needed
type RequestParameters struct {
	URL     string
	Headers string // Request Headers
}

// Get request
func Get() error {
	return request("httpbin.org", "GET /status/418 HTTP/1.0\r\nHost: httpbin.org\r\n\r\n")
}

func request(host string, requestMessage string) error {

	// establish a TCP connection to a particular port on a server (port 80)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, 80))
	if err != nil {
		return fmt.Errorf("Error establishing TCP connection to \"%s:%d\" : %v", host, 80, err)
	}
	defer conn.Close()

	// sends HTTP request message to the server (through the TCP socket).
	_, err = fmt.Fprint(conn, requestMessage)
	if err != nil {
		return fmt.Errorf("Error sending HTTP request message: %v", err)
	}

	// receive response from the server
	response, err := ioutil.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("Error getting response: %v", err)
	}

	fmt.Println(string(response))

	return nil
}
