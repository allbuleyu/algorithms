package main

import (
	"fmt"
	"time"

	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/net/html"
	"net/http"
	"path"
	"os"
	"io"
	"runtime"
	"github.com/allbuleyu/algorithms/introductionToAlgorighms/ch15"
	"github.com/allbuleyu/algorithms/introductionToAlgorighms/ch2"
)

type myStruct struct {
	S string
	I int
	q string
}

func main() {
	a := []int{5, 2, 4, 6, 1, 3}
	a = []int{6,5,4,3,2,1}
	ch2.InsertSort(a)
	fmt.Println(a)
	return

	fib := ch15.DynamicFibo(10)
	fmt.Println(fib)
	return

	defer printStack()
	f(3)

	return

	add1 := func(r rune) rune { return r + 1 }
	fmt.Println(strings.Map(add1, "hyl"))

	fmt.Printf("%*.f", 2, 2.22)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)

	f(x-1)
}

func printStack ()  {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func findWeek(d time.Time) string {
	year, week := d.ISOWeek()
	res := fmt.Sprintf("%d-%d", year,week)
	return res
}

func findStartAndEnd(d time.Time) string {
	var start, end time.Time
	switch d.Weekday() {
	case time.Monday:
		start = d.Add(0 * time.Hour * 24)
		end = d.Add(6 * time.Hour * 24)
	case time.Tuesday:
		start = d.Add(-1 * time.Hour * 24)
		end = d.Add(5 * time.Hour * 24)
	case time.Wednesday:
		start = d.Add(-2 * time.Hour * 24)
		end = d.Add(4 * time.Hour * 24)
	case time.Thursday:
		start = d.Add(-3 * time.Hour * 24)
		end = d.Add(3 * time.Hour * 24)
	case time.Friday:
		start = d.Add(-4 * time.Hour * 24)
		end = d.Add(2 * time.Hour * 24)
	case time.Saturday:
		start = d.Add(-5 * time.Hour * 24)
		end = d.Add(1 * time.Hour * 24)
	case time.Sunday:
		start = d.Add(-6 * time.Hour * 24)
		end = d.Add(0 * time.Hour * 24)
	}
	res := fmt.Sprintf("%s~%s", start.Format("01-02"), end.Format("01-02"))

	return res
}

type MongoData struct {
	Type int64
	Time time.Time
	TimeStamp int64
	PlatForm string
	TotalFee float64
}

func mongo() {
	url := "mongodb://dbuser:91z9nm29bv1u5z8k1j@192.168.1.90:27017/phpLog"
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("phpLog")

	c := db.C("user_pay")

	data := make([]MongoData, 0)
	c.Find(bson.M{"type": 1019}).All(&data)

	fmt.Println(data)
}


func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)

	return
}

func countWordsAndImages(n *html.Node) (words, images int) {

	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		words += len([]rune(data))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	var ww, ii int
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		ww, ii = countWordsAndImages(c)
		words, images = words+ww, images+ii
	}

	return
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	defer f.Close()
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	//if closerErr := f.Close(); err == nil {
	//	err = closerErr
	//}

	if  err != nil {
		return "", 0, err
	}

	return local, n, err
}