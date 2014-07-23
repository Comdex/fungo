package command

import (
	"errors"
	"fmt"
	"fungo/process"
	//"fungo/command"
)

func ComplieCommand(dirName string) {
	//先提取目录下的所有article存进Article数组
	process.GetAllArticles(dirName)
	if process.FileArray == nil {
		CheckError(errors.New("编译出错!"))
	}
	var articles []*process.Article
	articles = make(articles, len(process.FileArray))

}
