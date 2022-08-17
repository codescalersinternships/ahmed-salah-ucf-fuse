package main

import (
	"log"

	"bazil.org/fuse"
	"bazil.org/fuse/fuseutil"
	"golang.org/x/net/context"
)


type File struct {
	Node
	data []byte
}

// newFile creates a new File object
func (fs *FS) newFile(fileName string, fileData []byte) *File {
	return &File{
		Node: Node{ inode: fs.nextInode(), name: fileName },
		data: fileData,
	}
}

// Attr provides the core information for the file
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	log.Println("Requested Attr for File", f.name, "has data size", len(f.data))
	a.Inode = f.inode
	a.Mode = 0777
	a.Size = uint64(len(f.data))
	return nil
}

// Read handles requests to read data from the file
func (f *File) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	log.Println("Requested Read on File", f.name)
	fuseutil.HandleRead(req, resp, f.data)
	return nil
}

// ReadAll reads all of the file
func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	log.Println("Reading all of file", f.name)
	return []byte(f.data), nil
}