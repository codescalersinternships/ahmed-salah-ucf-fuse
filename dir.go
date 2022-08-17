package main

type Dir struct {
	Node
	files       *[]*File
	directories *[]*Dir
}