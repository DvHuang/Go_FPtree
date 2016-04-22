package textCpu

import (
	"bufio"
	"log"
	"os"
	"statistical/PublicStruct"
	"statistical/segment"
	"strconv"
	"strings"
)

const (
	searchMode = false
	//wbsCount   = 100000
)

var (
	wbs       = map[uint64]Weibo{}
	segmenter sego.Segmenter
)

type Weibo struct {
	Id           uint64 `json:"id"`
	Timestamp    uint64 `json:"timestamp"`
	UserName     string `json:"user_name"`
	RepostsCount uint64 `json:"reposts_count"`
	Text         string `json:"text"`
}

func InitWbData() map[uint64]Weibo {

	// 读入微博数据
	file, err := os.Open("D:/weibo_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "||||")
		if len(data) != 10 {
			continue
		}
		wb := Weibo{}
		wb.Id, _ = strconv.ParseUint(data[0], 10, 64)
		wb.Timestamp, _ = strconv.ParseUint(data[1], 10, 64)
		wb.UserName = data[3]
		wb.RepostsCount, _ = strconv.ParseUint(data[4], 10, 64)
		wb.Text = data[9]
		wbs[wb.Id] = wb
		/*if len(wbs) > wbsCount {
			break
		}*/
	}

	return wbs

}

//生成potree所需的通用数据结构,一个Itemsets存放一条事务

//itemsets的切片，即全部事务的集合
func WbToPyData(wbs map[uint64]Weibo) map[string]PublicStruct.MapPrefixPath { //正经用
	//func WbToPyData(wbs map[uint64]string) map[string]PublicStruct.MapPrefixPath {//test 用

	segmenter.LoadDictionary("D:/dictionary.txt")

	allitemsets := make(map[string]PublicStruct.MapPrefixPath) //指定长度之后，后面的append会在这个长度之后进行追加，也就是切片前面会有很多的空map

	for _, weibo := range wbs {
		//每一条微博文本进行分词
		var text []byte

		text = []byte(weibo.Text)  //正经用

		//text = []byte(weibo)  //test 用

		//fmt.Printf(weibo.Text)
		segments := segmenter.Segment(text)
		//将每一条分词后事务的数据结构改为简单的map[string]int
		//fmt.Println(segments)
		item, stringMapPrefixPath := sego.DavysegmentsToString(segments)
		//将事务归并到集合
		//fmt.Println("WbToPyData",stringMapPrefixPath,item)

		//allitemsets[stringMapPrefixPath].Path[key] =value
		allitemsetsvalue, _ := allitemsets[stringMapPrefixPath]
		allitemsetsvalue.Path = make(map[string]int)

		for key, value := range item {
			allitemsetsvalue.Path[key] = value
		}
		allitemsets[stringMapPrefixPath] = allitemsetsvalue
	}
	return allitemsets
}
func StringData(strmap map[string]string) map[string]PublicStruct.MapPrefixPath {

	segmenter.LoadDictionary("D:/dictionary.txt")

	allitemsets := make(map[string]PublicStruct.MapPrefixPath) //指定长度之后，后面的append会在这个长度之后进行追加，也就是切片前面会有很多的空map

	for _, value := range strmap {
		//每一条微博文本进行分词
		var text []byte

		text = []byte(value)
		//fmt.Printf(weibo.Text)
		segments := segmenter.Segment(text)
		//将每一条分词后事务的数据结构改为简单的map[string]int
		//fmt.Println(segments)
		item, stringMapPrefixPath := sego.DavysegmentsToString(segments)
		//将事务归并到集合
		//fmt.Println("WbToPyData",stringMapPrefixPath,item)

		//allitemsets[stringMapPrefixPath].Path[key] =value
		allitemsetsvalue, _ := allitemsets[stringMapPrefixPath]
		allitemsetsvalue.Path = make(map[string]int)

		for key, value := range item {
			allitemsetsvalue.Path[key] = value
		}
		allitemsets[stringMapPrefixPath] = allitemsetsvalue
	}
	return allitemsets
}

