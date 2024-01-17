package gofc

import (
	"os"
	"strings"
)

// GetAllFiles 列出文件夹下所有文件，返回map
// length-层级，=1则只返回一层，=2则返回两层
// 支持文件后缀匹配
func GetAllFiles(dir string) ([]string, error) {
	dirPath, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	sep := string(os.PathSeparator)
	for _, fi := range dirPath {
		if fi.IsDir() { // 如果还是一个目录，则递归去遍历
			subFiles, err := GetAllFiles(dir + sep + fi.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			// 过滤指定格式的文件
			ok := strings.HasSuffix(fi.Name(), ".md")
			if ok {
				files = append(files, dir+sep+fi.Name())
			}
			// files = append(files, dir+sep+fi.Name())
		}
	}
	return files, nil
}
