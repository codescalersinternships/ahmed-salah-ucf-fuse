# Fuse Filesystem

A simple filesystem to show the user some internal data from a running software.

## Overview

Fuse filesystem is a userspace file system, which lets the non-privileged users to create their own filesystem without editing the kernel code. This is done by running the actual filesystem code in the userspace while the FUSE interface works as a bridge between the userspace and the kernel.

## How to use it?

- Call in your program:

```go
Mount(mountPoint, appData)
```

that takes the mountPoint of the filesystem and data to be shown in this filesystem.

## Demo

```go
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
```
