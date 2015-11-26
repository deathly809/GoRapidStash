//
//	TODO: Think about this a lot.  Do we want to have the FileSystem
//	create and return a SafeWriter?
//


package filesystem

import (
	"io"
)

type mmapWriter struct {
	file File
}

func (writer *mmapWriter) Write(data []byte) (written int, err error) {
	return writer.file.Write(data)
}


// NewSafeWriter takes in a File object and returns a writer that
// allows users to Write to the file 
func NewSafeWriter(f File) io.Writer {
	return new(mmapWriter)
}
