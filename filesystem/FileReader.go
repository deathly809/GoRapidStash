package filesystem

import (
	"io"
)

type mmapReader struct {
	file File
}

func (reader *mmapReader) Read(p []byte) (n int, err error) {
	return reader.file.Read(p)
}


// NewReader takes in a File object and returns a reader that
// allows users to write to the file 
func NewReader(f File) io.Reader {
	result := new(mmapReader)
	result.file = f
	return result
}
