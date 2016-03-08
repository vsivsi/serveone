/*****************************************************************************
*     Copyright (C) 2016 by Vaughn Iverson
*     serveone is free software released under the MIT/X11 license.
*     See included LICENSE file for details.
*****************************************************************************/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n %s ", path.Base(os.Args[0]))
	flag.VisitAll(func(f *flag.Flag) { fmt.Fprint(os.Stderr, "[-", f.Name, " val] ") })
	fmt.Fprint(os.Stderr, "filename\n")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage
	port := flag.Uint("port", 8088, "Port to bind")
	flag.Parse()

	filename := flag.Arg(0) // Should be a filename

	file, err := os.Open(filename)
	if err != nil {
		panic("Aborting, could not open file!")
	} else {
		err = file.Close()
		if err != nil {
			panic("Aborting, could not close file!")
		}
	}

	fmt.Printf("Serving file %s on port %d\n", path.Base(filename), *port)

	handler := func (rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, filename)
	}

	http.HandleFunc(fmt.Sprintf("/%s", path.Base(filename)), handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
