package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// RequestParameters needed
type RequestParameters struct {
	URL           string
	RequestHeader RequestHeader
	Verbose       bool
	Data          string
	File          string
	Output        string
}

// RequestHeader map
type RequestHeader map[string]string

// Get request
func Get(params RequestParameters) error {
	host, path := parseURL(params.URL)
	requestLine := fmt.Sprintf("GET %s HTTP/1.0", path)
	requestMessage := fmt.Sprintf("%s\r\n%s\r\n", requestLine, params.RequestHeader)
	return request(host, requestMessage, params)
}

// Post request
func Post(params RequestParameters) error {
	host, path := parseURL(params.URL)
	requestLine := fmt.Sprintf("POST %s HTTP/1.0", path)
	params.RequestHeader["Host"] = host
	params.RequestHeader["Content-Length"] = strconv.Itoa(len([]byte(params.Data)))
	requestMessage := fmt.Sprintf("%s\r\n%s\r\n%s", requestLine, params.RequestHeader, params.Data)
	return request(host, requestMessage, params)
}

func request(host string, requestMessage string, params RequestParameters) error {

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

	if params.Verbose {
		fmt.Println(string(res))
	} else {
		if params.Output != "" {
			ioutil.WriteFile(params.Output, res, 0666)
		} else {
			// split the header and body of the response message
			resMsg := strings.Split(string(res), "\r\n\r\n")
			fmt.Println(resMsg[1])
		}

	}

	return nil
}

func parseURL(u string) (host string, path string) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", ""
	}
	return URL.Hostname(), URL.EscapedPath()
}

// String implements the flag.Value interface
func (rh RequestHeader) String() string {
	s := ""
	for k, v := range rh {
		s += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	return s
}

// Set implements the flag.Value interface
func (rh RequestHeader) Set(s string) error {
	indexes := regexp.MustCompile(":").FindStringIndex(s)
	if len(indexes) < 2 {
		return fmt.Errorf("Header value must contain a key:value pair")
	}
	rh[s[:indexes[0]]] = s[indexes[1]:]
	return nil
}
