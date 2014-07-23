// fungo project main.go
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	//"github.com/astaxie/beego/config"
	"com"
	"fungo/command"
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
		if len(c.Args()) == 2 {
			cliName := c.Args()[0]

			if cliName == "new" {
				dirName := c.Args()[1]
				if dirName != "" {
					//获取当前路径并在此路径下创建目录
					//path,_:=os.Getwd()
					command.NewCommand(dirName)
				}
			} else {

			}

		} else {
			fmt.Println(message)
			path, _ := os.Getwd()
			//fmt.Println(path)
			a, _ := com.GetAllSubDirs(path + "/blog")
			fmt.Println(a)

		}
	}

	//子命令
	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "在指定目录添加一个用你输入的名字命名的文章模板",
			Action: func(c *cli.Context) {
				command.AddCommand(c.Args().First(), c.Args().Get(1))

			},
		},
		{
			Name:      "complie",
			ShortName: "c",
			Usage:     "根据输入的站点目录编译生成静态网站",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
