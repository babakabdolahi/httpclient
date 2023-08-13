package main

import (
	"flag"
	"fmt"

	"github.com/babakabdolahi/httpclient.git/fetch"
)

var (
	h   = flag.Bool("h", false, "help")
	b   = flag.Bool("b", false, "returning the body of the message")
	hd  = flag.Bool("hd", false, "returning the header of the message")
	req = flag.Bool("req", false, "returning the request of the message")

	msg = `This is an http client cli app. Here are the flags:
	-h: help
	-b: message body of the provided urls
	-hd: message header
	`
)

func main() {
	flag.Parse()
	if *h {
		fmt.Println(msg)
		return
	}
	if *b {
		for _, url := range flag.Args() {
			fmt.Printf("%s\n", fetch.Body(url))
		}
	}
	if *hd {
		for _, url := range flag.Args() {
			fmt.Printf("%v\n", fetch.Header(url))
		}
	}
	if *req {
		for _, url := range flag.Args() {
			fmt.Printf("%+v\n", fetch.Req(url))
		}
	}
}
