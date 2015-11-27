// The concrete package contains concrete implementations
// of the FileSystem interface.

package concrete

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"time"
	"errors"

	"github.com/deathly809/gorapidstash/fs/mmapfile"
)

const (
	// Major version of the filesystem
	Major = int32(0)
	// Minor version of the filesystem
	Minor = int32(1)
	// Patch version of the filesystem
	Patch = int32(0)
)

// Header information
var _Signature = []byte{0xD, 0xE, 0xA, 0xD, 0xB, 0xE, 0xE, 0xF}

const (
	_SignatureSize  = 8
	_VersionBytes   = 12
	_FileCountBytes = 12
	_SizeBytes      = 12
	_FirstFreeBytes = 12
	_MajorVersion   = 2

	// 	The header layout contains a signature, version, number of files, filesystem size,
	//	and first block of free list.
	// signature 	= 8 bytes
	// version   	= 12 bytes
	// number files = 16 bytes
	// size			= 16 bytes
	// first free	= 16 bytes
	_HeaderSize = _SignatureSize + _VersionBytes + _FileCountBytes + _SizeBytes + _FirstFreeBytes
)

// Each file in the FileSystem is represented by a linked
// list structure
type fileNode struct {
	data []byte
	next *fileNode
}

// This is the header for each file in the filesystem
type fileSystemFile struct {
	pos          int
	size         int
	first        *fileNode
	created      time.Time
	lastModified time.Time
}

// The actual implementation
type fileSystemImpl struct {
	numFiles  int64
	firstFree fileNode
	safeFiles map[string]File
	files     map[string]fileNode
	mFile     File
}

func (fSys *fileSystemImpl) readHeader() error {
	header := make([]byte, _HeaderSize)

	fSys.mFile.Seek(0, Beginning)
	fSys.mFile.Read(header)

	buffer := bytes.NewReader(header)

	signature := make([]byte, _SignatureSize)
	buffer.Read(signature)

	var major, minor, patch int32
	binary.Read(buffer, binary.BigEndian, &major)
	binary.Read(buffer, binary.BigEndian, &minor)
	binary.Read(buffer, binary.BigEndian, &patch)

	if major != Major {
		msg := fmt.Sprintf("Trying to load an incompatible filesystem version: %d.%d.%d", major, minor, patch)
		return errors.New(msg)
	}

	var numFiles, size, firstFree int32
	binary.Read(buffer, binary.BigEndian, &numFiles)
	binary.Read(buffer, binary.BigEndian, &size)
	binary.Read(buffer, binary.BigEndian, &firstFree)

	return nil
}

// Initializes the filesystem after the MMAPFile has been
// opened
func (fSys *fileSystemImpl) init() error {
	//	Read filesystem header
	err := fSys.readHeader()
	if err != nil {
		return err
	}
}

// Open returns the berst FileSystem
func Open(filename string) FileSystem,error {
	result := new(fileSystemImpl)
	result.mFaile,err := mmapfile.NewFile(filename)
	
	if err != nil {
		return nil,errors.New("Could not open filesystem: " + err.Error())
	}

	return result,nil
}
