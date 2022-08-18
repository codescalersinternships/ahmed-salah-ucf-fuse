package fs

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fuseutil"
	"golang.org/x/net/context"
	"github.com/fatih/structs"
)


type File struct {
	Node
	data []byte
	filePath []string
	appData any
}

// newFile creates a new File object
func (fs *FS) newFile(fileName string, fileData []byte, filePath []string) *File {
	return &File{
		Node: Node{ inode: fs.nextInode(), name: fileName },
		data: fileData,
		filePath: filePath,
		appData: fs.appData,
	}
}

// Attr provides the core information for the file
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	log.Printf("%s", fmt.Sprintf("Requested Attr for File %s has data size %d", f.name, len(f.data)))
	a.Inode = f.inode
	a.Mode  = 0444
	a.Size  = uint64(len(f.data))
	a.Atime = time.Now()
	a.Mtime = time.Now()
	a.Ctime = time.Now()
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

	reflectedData := structs.Map(f.appData)

	return f.ReadFileContent(reflectedData), nil
}


func (f *File) ReadFileContent(data map[string]any) []byte {
	var content []byte
	var traverseReflectedData func(m map[string]any, idx int)

	traverseReflectedData = func (data map[string]any, idx int) {
		if idx == len(f.filePath) {
			content = []byte(fmt.Sprintln(reflect.ValueOf(data[f.name])))
		} else {
			traverseReflectedData(data[f.filePath[idx]].(map[string]any), idx+1)
		}
	}

	traverseReflectedData(data, 0)

	return content
}