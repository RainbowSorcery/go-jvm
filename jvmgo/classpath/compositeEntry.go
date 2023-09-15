package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	var compositeEntry []Entry

	for _, subPath := range strings.Split(path, pathListSeparator) {
		entry := newEntry(subPath)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (t CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range t {
		class, e, err := entry.ReadClass(className)
		if err == nil {
			return class, e, err
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (t CompositeEntry) String() string {
	StringList := make([]string, len(t))

	for i, entry := range t {
		StringList[i] = entry.String()
	}

	return strings.Join(StringList, pathListSeparator)
}
