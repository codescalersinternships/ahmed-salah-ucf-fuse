package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)
type Dir struct {
	Node
	files       *[]*File
	directories *[]*Dir
}

// newDir creates a new Dir object
func (fs *FS) newDir(dirName string) *Dir {
	return &Dir{
		Node:		 Node{ inode: fs.nextInode(), name: dirName},
		files: 		 &[]*File{ },
		directories: &[]*Dir{ },
	}
}

// Attr provides the core information for a directory
func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	log.Println("Requested Attr for Directory", d.name)
	a.Inode = d.inode
	a.Mode  = os.ModeDir | 0444
	a.Atime = time.Now()
	a.Mtime = time.Now()
	a.Ctime = time.Now()
	return nil
}

// Lookup provides the Node that matches that name, otherwise, return fuse.ENOENT.
// It could be either a File or a sub-Dir
func (d *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	log.Printf("%s", fmt.Sprintf("Requested lookup for %s in %s\n", name, d.name))
	if d.files != nil {
		for _, node := range *d.files {
			if node.name == name {
				log.Println("Found match for directory lookup with size", len(node.data))
				return node, nil
			}
		}
	}
	if d.directories != nil {
		for _, node := range *d.directories {
			if node.name == name {
				log.Println("Found match for directory lookup")
				return node, nil
			}
		}
	}
	return nil, fuse.ENOENT
}

// ReadDirAll returns a slice of fuse.Dirent
// for all Files and Dirs in the provided Dir
func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	log.Printf("%s", fmt.Sprintf("Reading all dirs of %s\n", d.name))
	var children []fuse.Dirent
	if d.files != nil {
		for _, file := range *d.files {
			children = append(children, fuse.Dirent{Inode: file.inode, Type: fuse.DT_File, Name: file.name})
		}
	}
	if d.directories != nil {
		for _, dir := range *d.directories {
			children = append(children, fuse.Dirent{Inode: dir.inode, Type: fuse.DT_Dir, Name: dir.name})
		}
		log.Printf("%s", fmt.Sprintf("%d children for dir %s", len(children), d.name))
	}
	return children, nil
}