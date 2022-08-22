package fs

import (
	"fmt"
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
	attributes fuse.Attr
}

// newFile creates a new File object
func (fs *FS) newFile(fileName string, data []byte, filePath []string) *File {
	return &File{
		Node: Node{ inode: fs.nextInode(), name: fileName },
		data: data,
		filePath: filePath,
		appData: fs.appData,
		attributes: fuse.Attr{
			Inode: fs.inode,
			Mode: 0444,
			Size: uint64(len(data)),
			Atime: time.Now(),
			Mtime: time.Now(),
			Ctime: time.Now(),
		},
	}
}

// Attr provides the core information for the file
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {	
	*a = f.attributes
	return nil
}

// Read handles requests to read data from the file
func (f *File) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	reflectedData := structs.Map(f.appData)
	fuseutil.HandleRead(req, resp, f.ReadFileContent(reflectedData))

	return nil
}

// ReadAll reads all of the file
func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	reflectedData := structs.Map(f.appData)

	return []byte(fmt.Sprintf("%s\n", string(f.ReadFileContent(reflectedData)))), nil
}

// ReadFileContent gets content of associated receiver from provided data
func (f *File) ReadFileContent(data map[string]any) []byte {
	for _, part := range f.filePath {
		data = data[part].(map[string]any)
	}

	content := []byte(fmt.Sprintln(reflect.ValueOf(data[f.name])))
	f.attributes.Size = uint64(len(content))

	return content
}