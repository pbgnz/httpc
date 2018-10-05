package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
)

// RequestParameters needed
type RequestParameters struct {
	URL         string
	HeaderLines string // Request Headers
}

// Get request
func Get(params RequestParameters) error {
	host, path := parseURL(params.URL)
	requestLine := fmt.Sprintf("GET %s HTTP/1.0", path)
	requestMessage := fmt.Sprintf("%s\r\n%s\r\n", requestLine, "Host: httpbin.org\r\n")
	return request(host, requestMessage)
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
	res, err := ioutil.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("Error getting response: %v", err)
	}

	fmt.Println(string(res))

	return nil
}

func parseURL(u string) (host string, path string) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", ""
	}
	return URL.Hostname(), URL.EscapedPath()
}
