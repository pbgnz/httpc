package main

import (
	"flag"
	"fmt"
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
		URL:     URL,
		Headers: *h,
	}

	fmt.Println(params)
	// httpc get [-v] [-h key:value] URL
	if err := Get(); err != nil {
		log.Fatal(err)
	}
}
