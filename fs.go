package main

import (
	"fmt"
	"reflect"

	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

type FS struct {
	inode uint64
	root *Dir
	data map[string]any
}

type Node struct {
	inode uint64
	name  string
}

// newFS creates new FS object
func newFS(data MyData) *FS {
	dataMap := structs.Map(data)

	return &FS{
		data:  dataMap,
		inode: 0,
		root: &Dir{
			Node: 		 Node{name: "root", inode: 1},
			files: 		 &[]*File{ },
			directories: &[]*Dir{ },
		},
	}
}

// nextInode determines the next Inode number that can be
// assigned to the next new node in the file system
func (fs *FS) nextInode() uint64 {
	return (fs.inode + 1)
}

// Root is called to obtain the Node for the file system root.
func (fs *FS) Root() (fs.Node, error) {
	return fs.root, nil
}

// reflectDataIntoFS reflects data given as argument into file system
func (fs *FS) reflectDataIntoFS(data any, currentDir *Dir) {
	for key, val := range data.(map[string]any) {
		if reflect.TypeOf(val).Kind() == reflect.Map {
			newDir := fs.newDir(key)
			*(currentDir.directories) = append(*(currentDir.directories), newDir)
			fs.reflectDataIntoFS(val, newDir)
		} else {
			newFile := fs.newFile(key, []byte(fmt.Sprint(reflect.ValueOf(val))))
			*(currentDir.files) = append(*(currentDir.files), newFile)
		}
	}
}