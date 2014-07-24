package process

import (
	"os"
	"path/filepath"
	"strings"
)

//要编译的文章模板
var FileArray []string

//根据输入的内容和摘要长度获取文章摘要
func GetSummary(content string, length int) string {
	return string(content[0 : length+1])
}

//根据输入的文章总数和分页条数返回分页数
func GetPageNum(articleNum, num int) int {
	i := articleNum % num
	var pageNum int
	if i != 0 {
		pageNum = articleNum / num
		pageNum = pageNum + 1
	} else {
		pageNum = articleNum / num
	}
	return pageNum
}

//根据输入的目录获取所有后缀为.article的文件路径名切片
func GetAllArticles(dirName string) {
	filepath.Walk(dirName, walkFunc)
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".article") {
		FileArray = append(FileArray, path)
	}
	return nil
}

//func endWith(s, pattern string) bool {
//	pos := strings.LastIndex(s, pattern)
//	if pos == -1 {
//		return false
//	}
//	if (pos + len(pattern)) == len(s) {
//		return true
//	} else {
//		return false
//	}
//}
