package fs

import (
	"io"
)

// FileSystem is an interface into your brain
type FileSystem interface {

	// GetSafeWriter takes in a File and returns a SafeWriter
	GetSafeWriter(File) io.Writer

	// GetSafeReader takes in a File and returns a SafeReader
	GetSafeReader(File) io.Reader

	//	Shutdown will abort all reads and writes in the
	//	FileSystem.  Any future calls to any function will
	//	have no effect.
	Shutdown()

	//	Lock will lock a file so that only the thread
	//	which has locked the file can Read or Write to
	//	it
	Lock(File)

	// Unlock unlocks a file for Reading and Writing.
	Unlock(File)

	//	Open takes in a filename and returns a File.  If
	//	the file does not exist then it is created.
	Open(string) File

	// Exists returns true if a file exists in the filesystem, false otherwise
	Exists(string) bool

	// EnableMVCC takes in a boolean value which  determines if MVCC is enabled
	EnableMVCC(bool)

	// Delete takes a filename and deletes the file from the filesystem if it
	// exists
	Delete(string)
}
