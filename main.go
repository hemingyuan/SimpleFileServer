package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		dir, IP string
		port    int
	)

	flag.StringVar(&dir, "d", "./", "USAGE:[-d DIR] DEFAULT: Current Directory")
	flag.StringVar(&IP, "i", "0.0.0.0", "USAGE:[-i IP] DEFAULT: 0.0.0.0")
	flag.IntVar(&port, "p", 23333, "USAGE:[-p PORT] DEFAULT: 23333")
	flag.Parse()

	ListenAddr := fmt.Sprintf("%s:%d", IP, port)

	log.Printf("SimpleFileServer is running on http://%s \n", ListenAddr)

	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(ListenAddr, nil)
	fmt.Println(err)
}
