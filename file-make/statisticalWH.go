package main

import (
	"fmt"
	"github.com/davylin/statistical/segment"
	"github.com/davylin/statistical/wordtest"
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 判断文件是否存在  存在返回 true 不存在返回false

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func main() {

	// 以下三个变量用来写入文件
	var filename = "./output1.txt"
	var f *os.File
	var err1 error

	var baiduMainUrl string = "http://github.com/cznic/mathutil"
	fmt.Printf("baiduMainUrl=%s\n", baiduMainUrl)
	resp, err := http.Get(baiduMainUrl)
	if err != nil {
		fmt.Printf("http get response errror=%s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//   fmt.Printf("body=%s\n", body)

	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	fmt.Printf("f is %v\n", f)
	err2 := ioutil.WriteFile(filename, body, 0666) //写入文件(字符串)
	check(err2)
	// fmt.Printf("写入 %d 个字节n", n);
	inputFilePath := filename
	outputFilePath := "./output.txt"

	wordtest.CountTestBase(inputFilePath, outputFilePath)

	var segmenter sego.Segmenter
	segmenter.LoadDictionary("github.com/davylin/statistical/segment/dictionary.txt")

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToString(segments, false))

}
