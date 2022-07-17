package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var proxy = "ec2-18-132-120-106.eu-west-2.compute.amazonaws.com:443"

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		fmt.Println(usage())
		return
	}
	switch os.Args[1] {
	case "--help", "-h":
		fmt.Println(usage())
		return
	case "publish":
		if argsLen < 3 {
			fmt.Println(usage())
			return
		}
		dir := os.Args[2]
		conn, err := net.Dial("tcp", proxy)
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()
		fmt.Fprintf(conn, "PUB %s\n", dir)
	default:
		fmt.Println(usage())
	}
}

func usage() string {
	return `
ydy is a tool for generating a gallery of images from a directory.

Usage: 
	ydy <command> [arguments]

Commands:

	publish		temporarily publish your folder as a gallery.
	`
}
