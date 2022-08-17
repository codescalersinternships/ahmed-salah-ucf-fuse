package main

import "bazil.org/fuse/fs"

type FS struct {
	inode uint64
	root *Dir
}

type Node struct {
	inode uint64
	name  string
}

// Root is called to obtain the Node for the file system root.
func (f *FS) Root() (fs.Node, error) {
	return f.root, nil
}