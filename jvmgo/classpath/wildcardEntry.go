package classpath

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	var wildcardEntry []Entry

	basePath := path[:len(path)-1]
	err := filepath.Walk(basePath, func(path string, info fs.FileInfo, err error) error {
		// 跳过子目录 通配符只能匹配以及目录
		if info.IsDir() && path != basePath {
			return filepath.SkipDir
		}

		// 判断文件是否是已jar结尾或是以zip结尾 如果是则新家zipEntry
		if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
			jarEntry := newZipEntry(path)
			wildcardEntry = append(wildcardEntry, jarEntry)
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return wildcardEntry
}
