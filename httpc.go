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
	o := flag.String("o", "", "Write the body of the response to the specified file instead of the console.")
	p := flag.Int("p", 80, "Specifies the port number of the server.")
	flag.Parse()

	method := flag.Arg(0)
	URL := flag.Arg(1)
	if URL == "" {
		log.Fatal("Error: URL is required")
	}

	if method != "get" && method != "post" {
		log.Fatalf("Error: unsuported method: %v.", method)
	}

	if method == "get" && *d != "" {
		log.Fatalln("Error: get method should not used with the flag -d.")
	}

	if method == "get" && *f != "" {
		log.Fatalln("Error: get method should not used with the flag -f.")
	}

	if *d != "" && *f != "" {
		log.Fatalln("Error: post method should have either -d or -f but not both flags.")
	}

	params := RequestParameters{
		URL:           URL,
		RequestHeader: rh,
		Verbose:       *v,
		Data:          *d,
		File:          *f,
		Output:        *o,
		Port:          *p,
	}

	if *f != "" {
		if data, err := ioutil.ReadFile(*f); err == nil {
			params.Data = string(data)
		}
	}

	switch method {
	case "get":
		if err := Get(params); err != nil {
			log.Fatal(err)
		}
	case "post":
		if err := Post(params); err != nil {
			log.Fatal(err)
		}
	}
}
