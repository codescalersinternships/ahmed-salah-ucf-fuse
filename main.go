package main

import (
	"flag"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

var usage = func() {
	log.Printf("Usage of %s:\n", os.Args[0])
	log.Printf("  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}

// func (f *FS) 
func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	mountPoint := flag.Arg(0)
	conn, err := fuse.Mount(mountPoint, fuse.ReadOnly())
	if err != nil {
		log.Fatal(err)
	}
	server := fs.New(conn, nil)
	fs := newFS()
	
	if err := server.Serve(fs); err != nil {
		log.Fatal(err)
	}
}