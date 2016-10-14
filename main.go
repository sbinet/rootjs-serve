package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
	dirFlag  = flag.String("dir", ".", "directory of files to serve")
	browser  = flag.String("browser", "", "path to browser to launch if not empty")
)

func main() {
	flag.Parse()
	dir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}

	if *browser != "" {
		go func() {
			log.Printf("launching browser at localhost%s/...\n", *addrFlag)
			cmd := exec.Command(*browser, "localhost"+*addrFlag+"/")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	serve(dir)
}

func serve(dir string) {
	log.Printf("serving at %q...\n", *addrFlag)
	err := http.ListenAndServe(*addrFlag, http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal(err)
	}
}
