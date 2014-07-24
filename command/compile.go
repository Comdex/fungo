package command

import (
	"errors"
	"fmt"
	"fungo/process"
	"html/template"
	"io/ioutil"
	"os"
	//"fungo/command"
)

func ComplieCommand(dirName string) {
	//先提取目录下的所有article存进Article数组
	process.GetAllArticles(dirName)
	if process.FileArray == nil {
		CheckError(errors.New("编译出错!"))
	}
	fmt.Println(process.FileArray)
	var articles []*process.Article
	articles = make([]*process.Article, len(process.FileArray))

	for i := 0; i < len(process.FileArray); i++ {
		file, err := os.Open(process.FileArray[i])
		CheckError(err)
		fileByte, err2 := ioutil.ReadAll(file)
		CheckError(err2)
		t := process.ParseArticle(string(fileByte))
		if t != nil {
			articles[i] = t
		} else {
			panic("编译出错!")
		}
	}

	//先读取配置文档信息
	site := process.ParseSiteConf(dirName + "\\site.conf")

	//解析summary进Article数组
	for i := 0; i < len(articles); i++ {
		articles[i].Summary = process.GetSummary(articles[i].Content, site.Summary)
	}

	for i := 0; i < len(articles); i++ {
		file, err := os.Create(dirName + "/" + articles[i].Article_url)
		fmt.Println(articles[i].Article_url)
		CheckError(err)
		//t := template.New(articles[i].Title)
		//tp, err2 := t.Parse(`<!DOCTYPE html>
		//<html lang='zh-CN'>
		//<head>
		//<meta charset='utf-8'>
		//</head>
		//<body>
		//<div id="article-container" class="clear">
		//  <article class="clear 9" id="article-{{.Title}}">
		//      <h3 class="title"><a href="" title="{{.Title}}">{{.Title}}</a></h3>
		//      <section class="content markdown">{{.Content}}</section>
		//      <p class="info clear">{{if .Tag}}
		//          <span class="author inline-block"><i class="fa fa-user"></i>{{.Tag}}</span>{{end}}
		//          <span class="time inline-block"><i class="fa fa-clock-o"></i>niin</span>nini
		//          <span class="tag inline-block"><i class="fa fa-tags"></i>nini<a href="u">nini</a>nini</span>nini
		//          <span class="views inline-block right">阅读&nbsp;&nbsp;<span>{{.Tag}}</span>&nbsp;&nbsp;次</span>
		//      </p>
		//  </article>
		//</div>
		//<body>
		//<html>`)
		tp, err2 := template.ParseFiles(dirName + "/article.tmpl")
		fmt.Println(dirName + "\\article.html")
		CheckError(err2)
		fmt.Println("zheer")
		CheckError(tp.Execute(file, articles[i]))
		fmt.Println("nali")
	}

}
