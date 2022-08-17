package main

type FS struct {
	inode uint64
	root *Dir
}

type Node struct {
	inode uint64
	name  string
}