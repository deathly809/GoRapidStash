
package filesystem

import (
	"io"
)

// File is an interface to represent basic operations of a file?
type File interface {
	Close() error
	Write(data []byte, pos int) (int, error)
	Read(pos, len, offset int) ([]byte, error)
	IsNew() bool
	Name() string
}

// FileSystem is an interface into your brain
type FileSystem interface {
	GetWriter() io.Writer
	GetReader() io.Reader
	Shutdown()
	Lock(File)
	Unlock(File)
}

// OpenFileSystem is the best FileSystem
func OpenFileSystem(filename string) FileSystem {
	return nil
}