package main

import (
	"flag"
	"fs"
	"log"
	"os"
	"time"
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

	var data = &MyData{
		Name: "Salah",
		Age:  22,
		Sub: SubStruct{
			SomeValue: 		3.14,
			SomeOtherValue: "some text...\n",
		},
	}

	// Testing procedures that alters data
	go updateAge(data)
	go updateName(data)

	if err := fs.Mount(mountPoint, data); err != nil {
		log.Fatal(err)
	}
}

func updateAge(data *MyData) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		data.Age += 1
	}
}

func updateName(data *MyData) {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		data.Name += "#"
	}
}

var usage = func() {
	log.Printf("Usage of %s:\n", os.Args[0])
	log.Printf("  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}