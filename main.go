package main

import (
	"fmt"
	"flag"
	"os/user"
	"time"
	"github.com/abadojack/whatlanggo"
	"os"
	"path/filepath"
	"strings"
	"github.com/daviddengcn/go-colortext"
	"log"
)


type Blog struct {
	Title         string
	Author        string
	CreateTime    string
	Category      string
	Summary       string
	CopyrightInfo string

	SavePath string
}

func IntelligentCategorize(title string, lang whatlanggo.Lang) string {
	if lang == whatlanggo.Eng {
		return "Default Category"
	} else if lang == whatlanggo.Cmn {
		return "默认分类"
	} else {
		return "默认分类"
	}
}

// automatically detect language
func detectLang(title string) whatlanggo.Lang {
	info := whatlanggo.Detect(title)
	lang := info.Lang
	return lang
}

func getCurrentUserName() string {
	currentUserName, err := user.Current()
	if err != nil {
		ct.Foreground(ct.Red, false)
		fmt.Println(err)
	}
	return currentUserName.Username
}
func getNowTimeString() string {
	t := time.Now()
	tFormat := t.Format("2006-01-02 15:05:05")
	return tFormat
}

// this can be change individually
func getCopyrightInfo(lang whatlanggo.Lang) string {
	copyright := ""
	if lang == whatlanggo.Cmn {
		copyright = "> " + "本文由在当地较为英俊的男子金天大神原创，版权所有，欢迎转载，" +
			"但请保留这段版权信息，多谢合作，有任何疑问欢迎通过微信联系我交流：`jintianiloveu` \n"
	} else {
		copyright = "> " + "This article was original written by Jin Tian, welcome re-post, " +
			"but please keep this copyright info, thanks, any question could be asked via wechat: `jintianiloveu` \n"
	}
	return copyright
}

// generate summary according to title
func generateSummary(title string, lang whatlanggo.Lang) string {
	pref := ""
	if lang == whatlanggo.Eng {
		pref = "Introduce something about "
	} else {
		pref = "本文介绍 "
	}
	return pref + title
}

// write blog into template
func WriteBlog(blog Blog) {
	if _, err := os.Stat(blog.SavePath); os.IsNotExist(err) {
		// does not exit create path
		os.MkdirAll(blog.SavePath, 777)
		log.Println("directory " + blog.SavePath + " doest not exist, so create it.")
	}
	// do the write thing
	// file name replace space in title, -1 means replace all
	fileName := strings.Replace(blog.Title, " ", "_", -1) + ".md"
	saveFile := filepath.Join(blog.SavePath, fileName)
	ct.Foreground(ct.Green, false)
	fmt.Println("file will be save into : " + saveFile)

	templateContent := "----\n" +
		"title: " + blog.Title + "\n" +
		"date: " + blog.CreateTime + "\n" +
		"category: " + blog.Category + "\n" +
		"---" + "\n" +
		blog.Summary + "\n" +
		"<!-- more -->" + "\n" +
		"# " + blog.Title + "\n" +
		blog.CopyrightInfo + "\n"

	f, err := os.Create(saveFile)
	if err != nil {
		fmt.Println("create file error, " + err.Error())
	} else {
		f.WriteString(templateContent)
	}
}

func main() {
	ct.Foreground(ct.Green, false)
	fmt.Println("Hello, blogo helps writing blog. generates template as you want.")
	ct.Foreground(ct.Red, false)
	fmt.Println("author: Jin Tian.")
	fmt.Println(`   ___    __     ____    _____  ____
  / _ )  / /    / __ \  / ___/ / __ \
 / _  | / /__  / /_/ / / (_ / / /_/ /
/____/ /____/  \____/  \___/  \____/
                                     `)

	title := flag.String("title", "No Title", "specific blog title, both English and Chinese are supported.")
	savePath := flag.String("path", "./", "blog file save path, default is current directory.")

	// remember to call this Parse method!!
	flag.Parse()

	lang := detectLang(*title)

	blog := Blog{
		Title:         *title,
		Author:        getCurrentUserName(),
		CreateTime:    getNowTimeString(),
		Category:      IntelligentCategorize(*title, lang),
		Summary:       generateSummary(*title, lang),
		SavePath:      *savePath,
		CopyrightInfo: getCopyrightInfo(lang),
	}

	WriteBlog(blog)
	ct.Foreground(ct.Red, false)
	fmt.Println("write blog templates success!")

}
