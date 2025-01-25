package tsoptions

import (
	"testing/fstest"

	"github.com/microsoft/typescript-go/internal/tspath"
	"github.com/microsoft/typescript-go/internal/vfs"
	"github.com/microsoft/typescript-go/internal/vfs/vfstest"
)

func fixRoot(path string) string {
	rootLength := tspath.GetRootLength(path)
	if rootLength == 0 {
		return path
	}
	if len(path) == rootLength {
		return "."
	}
	return path[rootLength:]
}

type VfsParseConfigHost struct {
	fs               vfs.FS
	currentDirectory string
}

var _ ParseConfigHost = (*VfsParseConfigHost)(nil)

func (h *VfsParseConfigHost) FS() vfs.FS {
	return h.fs
}

func (h *VfsParseConfigHost) GetCurrentDirectory() string {
	return h.currentDirectory
}

func NewVFSParseConfigHost(files map[string]string, currentDirectory string) *VfsParseConfigHost {
	fs := fstest.MapFS{}
	for name, content := range files {
		fs[fixRoot(name)] = &fstest.MapFile{
			Data: []byte(content),
		}
	}
	return &VfsParseConfigHost{
		fs:               vfstest.FromMapFS(fs, true /*useCaseSensitiveFileNames*/),
		currentDirectory: currentDirectory,
	}
}
