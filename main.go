package main

import (
	"flag"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type MyData struct {
    Name string
	Age int
	Sub SubStruct
}

type SubStruct struct {
    SomeValue float64
    SomeOtherValue string
}

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

	data := MyData{
        Name: "Salah",
		Age: 22,
		Sub: SubStruct{
			SomeValue: 3.14,
			SomeOtherValue: "some text...\n",
		},
    }

	fs := newFS(data)
	fs.reflectDataIntoFS(fs.data, fs.root)

	if err := server.Serve(fs); err != nil {
		log.Fatal(err)
	}
}