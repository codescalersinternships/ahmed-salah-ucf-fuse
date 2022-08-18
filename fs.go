package fs

import (
	"fmt"
	"log"
	"reflect"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

type FS struct {
	inode uint64
	root  *Dir
	appData any
}

type Node struct {
	inode uint64
	name  string
}

// newFS creates new FS object
func newFS(appData any) *FS {
	return &FS{
		inode: 0,
		root: &Dir{
			Node: 		 Node{name: "root", inode: 1},
			files: 		 []*File{ },
			directories: []*Dir{ },
		},
		appData: appData,
	}
}

// nextInode determines the next Inode number that can be
// assigned to the next new node in the file system
func (fs *FS) nextInode() uint64 {
	return (fs.inode + 1)
}

// Root is called to obtain the Node for the file system root.
func (fs *FS) Root() (fs.Node, error) {
	dataMap := structs.Map(fs.appData)
	fs.reflectDataIntoFS(dataMap, fs.root, []string{})
	return fs.root, nil
}

func Mount(mountPoint string, appData any) error {
	conn, err := fuse.Mount(mountPoint, fuse.ReadOnly())
	defer func () error {
		if err := conn.Close(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}()
	
	server := fs.New(conn, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func () error {
		if err := fuse.Unmount(mountPoint); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}()


	if err := server.Serve(newFS(appData)); err != nil {
		log.Fatal(err)
	}

	return nil
}

// reflectDataIntoFS reflects data given as argument into file system
func (fs *FS) reflectDataIntoFS(data any, currentDir *Dir, currentPath []string) {
	for key, val := range data.(map[string]any) {
		if reflect.TypeOf(val).Kind() == reflect.Map {
			newDir := fs.newDir(key)
			currentDir.directories = append(currentDir.directories, newDir)
			fs.reflectDataIntoFS(val, newDir, append(currentPath, key))
		} else {
			filePath := make([]string, len(currentPath))
			copy(filePath, currentPath)
			newFile := fs.newFile(key, []byte(fmt.Sprint(reflect.ValueOf(val))), filePath)
			currentDir.files = append(currentDir.files, newFile)
		}
	}
}