// fungo project main.go
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	//"github.com/astaxie/beego/config"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fungo"
	app.Usage = "fungo is a static website generator by Comdex \n   fungo 是一个静态网站生成器 by Comdex"
	app.Version = "0.1.0"

	//全局选项
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "new", Value: "网站目录名", Usage: `"fungo new 网站目录名"命令会在fungo所在的当前目录下生成一个静态网页目录内含样式文件等`},
	}

	app.Action = func(c *cli.Context) {
		//全局帮助提示
		message := "Please use command \"fungo -h\" or \"fungo --help\" get information for help \n请使用命令 \"fungo -h\" 或 \"fungo --help\" 获取帮助的信息"

		//分析全局选项命令
		if len(c.Args()) > 0 {
			cliName := c.Args()[0]
			//获取目录名
			if cliName == "new" {
				dirName := c.Args()[1]
				if dirName != "" {
					//获取当前路径并在此路径下创建目录
					//path,_:=os.Getwd()
					err := os.Mkdir(dirName, 0777)
					checkError(err)
					err1 := os.MkdirAll(dirName+`/static/css`, 0777)
					checkError(err1)
					err2 := os.MkdirAll(dirName+`/static/js`, 0777)
					checkError(err2)
					err3 := os.MkdirAll(dirName+`/static/img`, 0777)
					checkError(err3)
					//写入配置文件
					fout, err4 := os.Create(dirName + "/site.conf")
					checkError(err4)
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
					checkError(err5)

					//写入演示文章
					fout, err6 := os.Create(dirName + "/1.article")
					checkError(err6)
					defer fout.Close()

					_, err7 := fout.WriteString(`#文章标题
title：大家好,这是我的第一篇博客！
#标签
tag:大家好
#文章网页url,可选，默认为文件名
article-url:dajiahao.html
#内容
content:
这是我的第一篇博客！请大家多多支持啊，fungo是一个非常有趣的静态网页生成引擎，欢迎使用和提出建议!`)
					checkError(err7)
				}

			}

		} else {
			fmt.Println(message)
			path, _ := os.Getwd()
			fmt.Println(path)
		}

	}
	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "add a task to the list",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:      "complete",
			ShortName: "c",
			Usage:     "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:      "template",
			ShortName: "r",
			Usage:     "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

//错误检测函数
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("Successful Executing!")
	}
}
