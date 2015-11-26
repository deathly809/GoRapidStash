
package filesystem

import (
	"io"
)

// FileOffset is the relative offset in the file from
// a given  position
type FileOffset int

// Beginning
// Current
// End
const (
	Beginning FileOffset = iota
	Current
	End
)

// File is an interface to represent basic operations of a file?
type File interface {
	Close() error
	Write(data []byte) (int, error)
	Read(data []byte) (int, error)
	Seek(int,FileOffset)
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