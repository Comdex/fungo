// fungo project main.go
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	iconv "github.com/djimenez/iconv-go"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fungo"
	app.Usage = "fungo is a static website generator by Comdex \n   fungo 是一个静态网站生成器 by Comdex"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "new", value: "网站目录名", Usage: `"fungo new 网站目录名"命令会在fungo所在的当前目录下生成一个静态网页目录内含样式文件等`},
	}

	app.Action = func(c *cli.Context) {
		//全局帮助提示
		message := "Please use command \"fungo -h\" or \"fungo --help\" get information for help \n请使用命令 \"fungo -h\" 或 \"fungo --help\" 获取帮助的信息"
		cmessage, _ := iconv.ConvertString(message, "UTF-8", "GBK")

		//分析全局选项命令
		var dirName string
		if len(c.Args()) > 0 {
			dirName = c.Args()[0]

		}

		fmt.Println(cmessage)
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
