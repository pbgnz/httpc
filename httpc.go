package main

import (
	"flag"
	"log"
)

func main() {
	var rh RequestHeader = map[string]string{}
	flag.Var(rh, "h", "Associates headers to HTTP Request with the format 'key:value'.")
	v := flag.Bool("v", false, "Prints the detail of the response such as protocol,status,and headers.")
	d := flag.String("d", "", "Associates an inline data to the body HTTP POST request.")
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
	}

	if err := Post(params); err != nil {
		log.Fatal(err)
	}
}
