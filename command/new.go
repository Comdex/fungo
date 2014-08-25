package command

import (
	//"fmt"
	"os"
)

//new命令创建资源目录
func NewCommand(dirName string) {
	err := os.Mkdir(dirName, os.ModePerm)
	CheckError(err)
	//err1 := os.MkdirAll(dirName+`/static/css`, os.ModePerm)
	//CheckError(err1)
	//err2 := os.MkdirAll(dirName+`/static/js`, os.ModePerm)
	//CheckError(err2)
	//err3 := os.MkdirAll(dirName+`/static/img`, os.ModePerm)
	//CheckError(err3)
	//写入配置文件
	fout, err4 := os.Create(dirName + "/site.conf")
	CheckError(err4)
	defer fout.Close()

	_, err5 := fout.WriteString(`[global]
#博客/网站名字
title = "fungo"
#网站介绍
introduce = "fungo"
#网站地址
url = "www.example.com"
#网站作者
author = "fungo"
#网站关键字
keywords = "fungo"
#作者邮箱 
email = "fungo@fungo.com"
[personal]
#首页显示文章条数
latest = 6
#首页文章摘要长度
summary = 50`)
	CheckError(err5)

	//写入演示文章
	fout, err6 := os.Create(dirName + "/1.article")
	CheckError(err6)
	defer fout.Close()

	_, err7 := fout.WriteString(`#文章标题
title:大家好,这是我的第一篇博客！
#标签
tag:大家好
#文章网页url,可选，默认为文件名
article-url:dajiahao.html
#内容
content:
这是我的第一篇博客！请大家多多支持啊，fungo是一个非常有趣的静态网页生成引擎，欢迎使用和提出建议!`)
	CheckError(err7)

}
