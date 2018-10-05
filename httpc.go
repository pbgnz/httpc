package main

import (
	"flag"
	"log"
)

func main() {

	h := flag.String("h", "", "header key:value pair")
	flag.Parse()

	URL := flag.Arg(1)
	if URL == "" {
		log.Fatal("URL is required")
	}

	params := RequestParameters{
		URL:         URL,
		HeaderLines: *h,
	}

	if err := Get(params); err != nil {
		log.Fatal(err)
	}
}
