//
//	TODO: Think about this a lot.  Do we want to have the FileSystem
//	create and return a SafeReader?
//

package filesystem

import (
	"io"
)

type mmapReader struct {
	file File
}

func (reader *mmapReader) Read(p []byte) (int, error) {
	return reader.file.Read(p)
	
}


// NewSafeReader takes in a File object and returns a reader that
// allows users to write to the file 
func NewSafeReader(f File) io.Reader {
	result := new(mmapReader)
	result.file = f
	return result
}
