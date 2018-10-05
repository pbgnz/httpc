package main

import (
	"flag"
	"log"
)

func main() {
	var hl HeaderLines = map[string]string{}
	flag.Var(hl, "h", "Associates headers to HTTP Request with the format 'key:value'.")
	flag.Parse()

	URL := flag.Arg(1)
	if URL == "" {
		log.Fatal("URL is required")
	}

	params := RequestParameters{
		URL:         URL,
		HeaderLines: hl,
	}

	if err := Get(params); err != nil {
		log.Fatal(err)
	}
}
