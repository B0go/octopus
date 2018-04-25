package system

import "os"

//FileSystemReader reads the fs data
type FileSystemReader interface {
	FileStatus(string) (os.FileInfo, error)
	IsNotExist(error) bool
}

//OSFileSystemReader reads the current OS fs data
type OSFileSystemReader struct {
}

//DirectoryStatus retrieves directory data or error if it do not exists
func (fsReader OSFileSystemReader) FileStatus(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

//IsNotExist validates if the error is NotExist
func (fsReader OSFileSystemReader) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}
