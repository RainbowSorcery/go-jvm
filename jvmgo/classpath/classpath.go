package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}

	cp.parseBootClassPathAndExtClassPath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

func (t *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	bootClassPathData, bootClassPathEntry, err := t.bootClasspath.ReadClass(className)
	if err == nil {
		return bootClassPathData, bootClassPathEntry, err
	}

	userClassPathData, userClassPathEntry, err := t.userClasspath.ReadClass(className)
	if err == nil {
		return userClassPathData, userClassPathEntry, err
	}

	return t.extClasspath.ReadClass(className)
}

func (t *Classpath) String() string {
	return t.userClasspath.String()
}

func (t *Classpath) parseBootClassPathAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	t.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	t.extClasspath = newWildcardEntry(jreExtPath)

}

func (t *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	t.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	// 优先获取用户设置jre目录 如果目录存在那么直接返回
	if jreOption != "" && exist(jreOption) {
		return jreOption
	}

	// 其次判断当前目录下的jre目录 如果目录也存在 则直接返回
	if exist("./jre") {
		abs, err := filepath.Abs("./jre")
		if err != nil {
			panic(err)
		}
		return abs
	}

	// 如果当前目录下也没有jre目录 那么直接根据环境遍历找jre目录
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		return javaHome
	}

	panic("can not find jre folder!")
}

func exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
