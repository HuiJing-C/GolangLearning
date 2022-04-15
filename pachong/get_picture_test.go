package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var jpgRegex, _ = regexp.Compile(`src="/\wpload\wile/[0-9]{6}/[0-9]*/[0-9a-zA-z]*.jpg" /><br />`)
var totalPageRegex, _ = regexp.Compile(`[0-9]*</a><a href="/[0-9a-zA-Z_/]*.html">下页</a>*`)
var titleRegex, _ = regexp.Compile(`<h1>.*</h1>`)
var wg = &sync.WaitGroup{}
var rootDir = "J://xrmn//tmp"

type SS struct {
	maxPage int
	url     string
}

var all = make(map[string]*SS)

func init() {
	// all["Uxing"] = &SS{2, "https://www.xrmn5.cc/Uxing/"}
	// all["MyGirl"] = &SS{22, "https://www.xrmn5.cc/MyGirl/"}
	// all["YouWu"] = &SS{6, "https://www.xrmn5.cc/YouWu/"}
	// all["LeYuan"] = &SS{2, "https://www.xrmn5.cc/LeYuan/"}
	// all["MintYe"] = &SS{1, "https://www.xrmn5.cc/MintYe/"}
	// all["MTMeng"] = &SS{1, "https://www.xrmn5.cc/MTMeng/"}
	// all["FeiLin"] = &SS{12, "https://www.xrmn5.cc/FeiLin/"}
	// all["XingYan"] = &SS{5, "https://www.xrmn5.cc/XingYan/"}
	// all["MFStar"] = &SS{19, "https://www.xrmn5.cc/MFStar/"}
	// all["WingS"] = &SS{1, "https://www.xrmn5.cc/WingS/"}
	// all["HuaYan"] = &SS{3, "https://www.xrmn5.cc/HuaYan/"}
	// all["YouMi"] = &SS{26, "https://www.xrmn5.cc/YouMi/"}
	// all["Imiss"] = &SS{23, "https://www.xrmn5.cc/Imiss/"}
	// all["MiStar"] = &SS{11, "https://www.xrmn5.cc/MiStar/"}
	// all["Candy"] = &SS{3, "https://www.xrmn5.cc/Candy/"}
	// all["XiuRen"] = &SS{160, "https://www.xrmn5.cc/XiuRen/"}
	// all["Micat"] = &SS{3, "https://www.xrmn5.cc/Micat/"}
	// all["DKGirl"] = &SS{4, "https://www.xrmn5.cc/DKGirl/"}
	// all["HuaYang"] = &SS{17, "https://www.xrmn5.cc/HuaYang/"}
	// all["XiaoYu"] = &SS{25, "https://www.xrmn5.cc/XiaoYu/"}
	// all["Taste"] = &SS{1, "https://www.xrmn5.cc/Taste/"}
	all["BoLoli"] = &SS{5, "https://www.xrmn5.cc/BoLoli/"}
	all["MiiTao"] = &SS{5, "https://www.xrmn5.cc/MiiTao/"}
}

func TestAll(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for prefix, ss := range all {
		for i := 1; i <= ss.maxPage; i++ {
			link := getPageAllLink(i, prefix, ss.url)
			wg.Add(len(link))
			for j := range link {
				// l = "https://www.xrmn5.cc/Uxing/2021/20213084.html"
				go func(l string, w *sync.WaitGroup) {
					text := getHtmlText(l)
					maxPage, title := getTotalPageAndTitle(text)
					split := strings.Split(title, "]")
					allUrls := getAllUrls(l, maxPage)
					for k := range allUrls {
						downloadJpgs(allUrls[k], split[0][1:], split[1])
					}
					w.Done()
				}(link[j], wg)
			}
			wg.Wait()
			log.Printf("%s %d 下载完毕", ss.url, i)
			sec := rand.Intn(5) + 5
			time.Sleep(time.Duration(sec * 1e9))
		}
	}
}

func TestOthers(t *testing.T) {
	links := []string{
		"https://www.xrmn5.cc/XiaoYu/2022/202210449.html", "https://www.xrmn5.cc/XiuRen/2022/202210458.html",
		"https://www.xrmn5.cc/XiuRen/2022/202210457.html", "https://www.xrmn5.cc/XiuRen/2022/202210456.html",
		"https://www.xrmn5.cc/XiuRen/2022/202210454.html", "https://www.xrmn5.cc/XiuRen/2022/202210453.html",
		"https://www.xrmn5.cc/XiuRen/2022/202210452.html", "https://www.xrmn5.cc/XiuRen/2022/202210451.html",
		"https://www.xrmn5.cc/XiuRen/2022/202210450.html", "https://www.xrmn5.cc/XiuRen/2022/202210455.html",
	}
	wg.Add(len(links))
	for j := range links {
		// l = "https://www.xrmn5.cc/Uxing/2021/20213084.html"
		go func(l string, w *sync.WaitGroup) {
			text := getHtmlText(l)
			maxPage, title := getTotalPageAndTitle(text)
			split := strings.Split(title, "]")
			allUrls := getAllUrls(l, maxPage)
			for k := range allUrls {
				downloadJpgs(allUrls[k], split[0][1:], split[1])
			}
			w.Done()
		}(links[j], wg)
	}
	wg.Wait()
}

// 获取网页文本
func getHtmlText(url string) string {
	for {
		if text := getHtml(url); text != "" {
			return text
		}
		time.Sleep(5 * time.Second)
	}
}

func getHtml(url string) string {
	response, err2 := http.Get(url)
	if err2 != nil {
		log.Printf("getHtmlText error %s %v", url, err2)
		return ""
	}
	defer response.Body.Close()
	bufer := &bytes.Buffer{}
	io.Copy(bufer, response.Body)
	return bufer.String()
}

// 获取每套图总页数
func getTotalPageAndTitle(htmlText string) (page int, title string) {
	ss := totalPageRegex.FindAllString(htmlText, -1)
	if len(ss) != 1 {
		log.Printf("get page error %v", ss)
		page = 1
	} else {
		split := strings.Split(ss[0], "<")
		parseInt, err := strconv.ParseInt(split[0], 0, 64)
		if err != nil {
			log.Printf("parseInt error %v", err)
			page = 1
		}
		page = int(parseInt)
	}
	titles := titleRegex.FindAllString(htmlText, -1)
	if len(titles) != 1 {
		log.Printf("get title error %v", titles)
		return page, fmt.Sprintf("[Others]%d", time.Now().Unix())
	}
	a := strings.ReplaceAll(titles[0], "<h1>", "")
	title = strings.ReplaceAll(a, "</h1>", "")
	return
}

// 获取系列url
func getPageAllLink(page int, prefix, url string) (allLink []string) {
	re, _ := regexp.Compile(fmt.Sprintf(`href="/%s/[0-9]{4}/[0-9]*.html`, prefix))
	var pageUrl string
	if page < 2 {
		pageUrl = getHtmlText(fmt.Sprintf("%sindex.html", url))
	} else {
		pageUrl = getHtmlText(fmt.Sprintf("%sindex%d.html", url, page))
	}
	allStrings := re.FindAllString(pageUrl, -1)
	m := make(map[string]struct{})
	for i := range allStrings {
		suffix := strings.ReplaceAll(allStrings[i], fmt.Sprintf(`href="/%s/`, prefix), "")
		m[fmt.Sprintf("%s%s", url, suffix)] = struct{}{}
	}
	for k := range m {
		allLink = append(allLink, k)
	}
	return
}

// 获取每套图每张画的url
func getAllUrls(baseUrl string, page int) []string {
	var urls []string
	var url string
	baseUrl = strings.ReplaceAll(baseUrl, ".html", "")
	for i := 1; i <= page; i++ {
		if i == 1 {
			url = fmt.Sprintf("%s.html", baseUrl)
		} else {
			url = fmt.Sprintf("%s_%d.html", baseUrl, i-1)
		}
		urls = append(urls, url)
	}
	return urls
}

func downloadJpgs(url, unit, dir string) {
	allPics := jpgRegex.FindAllString(getHtmlText(url), -1)
	for i := range allPics {
		suffix := strings.Split(allPics[i], "\"")[1] // uploadfile/202101/16/E910558509.jpg
		suffix = strings.Split(suffix, "ile/")[1]    // 202101/16/E910558509.jpg
		imageUrl := fmt.Sprintf("https://t.xrmn5.cc/UploadFile/%s", suffix)
		for {
			success, fileName := download(imageUrl, unit, dir)
			if success {
				break
			}
			os.Remove(fileName)
		}
	}
}

// 下载图片
func download(imgUrl, unit, dir string) (bool, string) {
	split := strings.Split(imgUrl, "/")
	fileName := fmt.Sprintf("%s//%s//%s//%d_%s", rootDir, unit, dir, time.Now().Unix(), split[len(split)-1])
	if mkdirErr := os.MkdirAll(fmt.Sprintf("%s//%s//%s", rootDir, unit, dir), os.ModePerm); mkdirErr != nil {
		log.Printf("文件夹 %s 创建失败 %v", fileName, mkdirErr)
		return false, fileName
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("文件 %s 创建失败 %v", fileName, err)
		return false, fileName
	}
	defer f.Close() // 结束关闭文件

	if resp, err := http.Get(imgUrl); err != nil {
		log.Printf("http.get err %s %v ", fileName, err)
		return false, fileName
	} else {
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			log.Printf("读取数据失败 %s %v", fileName, err1)
			return false, fileName
		}
		defer resp.Body.Close() // 结束关闭
		f.Write(body)
	}
	return true, fileName
}
