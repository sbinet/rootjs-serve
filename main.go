package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
	dirFlag  = flag.String("dir", ".", "directory of files to serve")
)

func main() {
	flag.Parse()
	dir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(*addrFlag, http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal(err)
	}
}
