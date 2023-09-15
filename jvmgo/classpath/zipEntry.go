package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type zipEntry struct {
	absPath string // class文件绝对路径
}

func newZipEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{absPath: absPath}
}

func (t *zipEntry) ReadClass(className string) ([]byte, Entry, error) {
	// 遍历zip中的所有文件 判断文件名称是否与传入的文件名称一致 如果一致则读出来并返回 如果在读取的过程中出现了错误则捕获错误并返回

	reader, err := zip.OpenReader(t.absPath)

	if err != nil {
		panic(err)
	}

	defer func(reader *zip.ReadCloser) {
		err := reader.Close()
		if err != nil {
			panic(err)
		}
	}(reader)

	for _, file := range reader.File {
		if file.Name == className {
			open, err := file.Open()

			if err != nil {
				return nil, nil, err
			}

			open.Close()

			data, err := io.ReadAll(open)

			if err != nil {
				return nil, nil, err
			}
			return data, t, nil
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (t *zipEntry) String() string {
	return t.absPath
}
