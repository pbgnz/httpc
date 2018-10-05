package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var rh RequestHeader = map[string]string{}
	flag.Var(rh, "h", "Associates headers to HTTP Request with the format 'key:value'.")
	v := flag.Bool("v", false, "Prints the detail of the response such as protocol,status,and headers.")
	d := flag.String("d", "", "Associates an inline data to the body HTTP POST request.")
	f := flag.String("f", "", "Associates the content of a file to the body HTTP POST.")
	flag.Parse()

	URL := flag.Arg(1)
	if URL == "" {
		log.Fatal("URL is required")
	}

	params := RequestParameters{
		URL:           URL,
		RequestHeader: rh,
		Verbose:       *v,
		Data:          *d,
		File:          *f,
	}

	if *d == "" {
		if data, err := ioutil.ReadFile(*f); err == nil {
			params.Data = string(data)
		}
	}

	if err := Post(params); err != nil {
		log.Fatal(err)
	}
}
