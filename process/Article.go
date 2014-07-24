package process

import (
	"fmt"
	//"io/ioutil"
	"os"
	"strings"
)

//文章结构体
type Article struct {
	Title, Tag, Article_url, Content, Summary string
}

//解析文章,返回Article结构体
func ParseArticle(content string) *Article {
	var art = new(Article)
	//获取文章标题
	fmt.Println(content)
	title_pos := strings.Index(content, "title:")
	fmt.Println(title_pos)
	checkError(title_pos, "title")
	p1 := string(content[title_pos:])
	d1 := strings.Split(p1, "\n")
	titleString := d1[0]
	art.Title = string(titleString[6:])

	//获取文章标签
	tag_pos := strings.Index(content, "tag:")
	checkError(tag_pos, "tag")
	p2 := string(content[tag_pos:])
	d2 := strings.Split(p2, "\n")
	tagString := d2[0]
	art.Tag = string(tagString[4:])

	//获取文章url
	article_url_pos := strings.Index(content, "article-url:")
	checkError(article_url_pos, "url")
	p3 := string(content[article_url_pos:])
	d3 := strings.Split(p3, "\n")
	article := d3[0]
	art.Article_url = string(article[12:])

	//获取文章内容
	content_pos := strings.Index(content, "content:")
	checkError(content_pos, "content")
	art.Content = string(content[content_pos+9:])
	return art

}

//错误检测,传入-1代表解析出错
func checkError(code int, err string) {
	if code == -1 {
		fmt.Println("文章解析错误，请检查所提交的文章格式是否正确！" + err)
		os.Exit(1)
	}
}
