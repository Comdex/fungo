package process

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"os"
)

type Site struct {
	Title, Introduce, Url, Author, Keywords, Email string
	Latest, Summary                                int
}

func ParseSiteConf(fileName string) *Site {
	var site *Site = new(Site)
	c, err := goconfig.LoadConfigFile(fileName)
	checkParseError(err)
	var err2, err3, err4, err5, err6, err7, err8, err9 error
	site.Title, err2 = c.GetValue("global", "title")
	checkParseError(err2)
	site.Introduce, err3 = c.GetValue("global", "introduce")
	checkParseError(err3)
	site.Url, err4 = c.GetValue("global", "url")
	checkParseError(err4)
	site.Author, err5 = c.GetValue("global", "author")
	checkParseError(err5)
	site.Keywords, err6 = c.GetValue("global", "keywords")
	checkParseError(err6)
	site.Email, err7 = c.GetValue("global", "email")
	checkParseError(err7)
	site.Latest, err8 = c.Int("personal", "latest")
	checkParseError(err8)
	site.Summary, err9 = c.Int("personal", "summary")
	checkParseError(err9)
	return site
}

func checkParseError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
