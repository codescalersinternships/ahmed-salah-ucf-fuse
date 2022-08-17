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

func newFS() *FS {
	return &FS{
		inode: 0,
		root: &Dir{
			Node: Node{name: "root", inode: 1},
			files: &[]*File{ },
			directories: &[]*Dir{ },
		},
	}
}

func (f *FS) nextInode() uint64 {
	return (f.inode + 1)
}

// Root is called to obtain the Node for the file system root.
func (f *FS) Root() (fs.Node, error) {
	return f.root, nil
}