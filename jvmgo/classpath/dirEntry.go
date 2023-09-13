package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absPath string // class文件绝对路径
}

func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{absPath: absPath}
}

func (t *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(className, t.absPath)

	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return data, t, err
}

func (t *DirEntry) String() string  {
	return t.absPath
}