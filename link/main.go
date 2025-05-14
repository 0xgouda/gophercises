package main

import (
	"flag"
	"fmt"
	"link_parser/link"
)

func main() {
	file_name := flag.String("file", "ex2.html", "html file name")
	flag.Parse()

	links := link.Parse(*file_name)		
	fmt.Println(links)
}