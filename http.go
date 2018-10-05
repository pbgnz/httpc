package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"regexp"
)

// RequestParameters needed
type RequestParameters struct {
	URL         string
	HeaderLines HeaderLines
}

// Get request
func Get(params RequestParameters) error {
	host, path, err := parseURL(params.URL)
	if err != nil {
		log.Fatal(err)
	}
	requestLine := fmt.Sprintf("GET %s HTTP/1.0", path)
	fmt.Println(requestLine)
	requestMessage := fmt.Sprintf("%s\r\n%s\r\n", requestLine, params.HeaderLines)
	return request(host, requestMessage)
}

// HeaderLines map
type HeaderLines map[string]string

// String implements the flag.Value interface
func (h HeaderLines) String() string {
	s := ""
	for key, value := range h {
		s += fmt.Sprintf("%s: %v\r\n", key, value)
	}
	return s
}

// Set implements the flag.Value interface
func (h HeaderLines) Set(s string) error {
	indexes := regexp.MustCompile(":").FindStringIndex(s)
	if len(indexes) < 2 {
		return fmt.Errorf("Header value must contain a key:value pair")
	}
	h[s[:indexes[0]]] = s[indexes[1]:]
	return nil
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

func parseURL(u string) (host string, path string, err error) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", "", fmt.Errorf("Error parsing the URL: %v", err)
	}
	return URL.Hostname(), URL.EscapedPath(), nil
}
