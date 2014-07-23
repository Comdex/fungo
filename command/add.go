package command

import (
	"os"
)

//Add命令创建文章模板
func AddCommand(filename, dirname string) {
	file, err := os.Create(dirname + "/" + filename + ".article")
	CheckError(err)
	_, err1 := file.WriteString(`#文章标题
title：大家好,这是我的第一篇博客！
#标签
tag:大家好
#文章网页url,可选，默认为文件名
article-url:dajiahao.html
#内容
content:
这是我的第一篇博客！请大家多多支持啊，fungo是一个非常有趣的静态网页生成引擎，欢迎使用和提出建议!`)
	CheckError(err1)
}
