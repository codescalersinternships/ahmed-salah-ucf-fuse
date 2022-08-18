package main

import (
	"flag"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

type MyData struct {
    Name string
	Age  int
	Sub  SubStruct
}

type SubStruct struct {
    SomeValue 	   float64
    SomeOtherValue string
}

var data = &MyData{
	Name: "Salah",
	Age:  22,
	Sub: SubStruct{
		SomeValue: 		3.14,
		SomeOtherValue: "some text...\n",
	},
}

var fileSystem = newFS(data)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}

	mountPoint := flag.Arg(0)
	if err := os.MkdirAll(mountPoint, os.ModeDir | 0444); err != nil {
		log.Fatal(err)
	}
	conn, err := fuse.Mount(mountPoint, fuse.ReadOnly())
	if err != nil {
		log.Fatal(err)
	}
	defer fuse.Unmount(mountPoint)

	server := fs.New(conn, nil)

	dataMap := structs.Map(data)
	fileSystem.reflectDataIntoFS(dataMap, fileSystem.root)

	if err := server.Serve(fileSystem); err != nil {
		log.Fatal(err)
	}
}

var usage = func() {
	log.Printf("Usage of %s:\n", os.Args[0])
	log.Printf("  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}