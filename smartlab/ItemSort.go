/*
package fpTree
import (
	"fmt"
	"sort"
	"bytes"
	"io/ioutil"

	"os"
"statistical/segment"

)


type WordCountBean struct {
	word  string
	count int
}
func NewWordCountBean(word string, count int) *WordCountBean {
	return &WordCountBean{word, count}
}

type WordCountBeanList []*WordCountBean

func CheckError(err error, msg string) {
	if err != nil {
		panic(msg + "," + err.Error())
	}
}
func (list WordCountBeanList) totalCount() int {
	totalCount := 0
	for _, v := range list {
		totalCount += v.count
	}

	return totalCount
}

func (list WordCountBeanList) Len() int {
	return len(list)
}

func (list WordCountBeanList) Less(i, j int) bool {
	if list[i].count > list[j].count {
		return true
	} else if list[i].count < list[j].count {
		return false
	} else {
		return list[i].word < list[j].word
	}
}

func (list WordCountBeanList) Swap(i, j int) {
	var temp *WordCountBean = list[i]
	list[i] = list[j]
	list[j] = temp
}



func ItemSetsSort( wordMap map[string]int)  map[string]int {


	list := make(WordCountBeanList, 0)

	for k, v := range wordMap {
		list = append(list, NewWordCountBean(k, v))
	}
	sort.Sort(list)
	//fmt.Printf(" wordMap[twotoone]++%v\n\n\n\n",list)
	wordsCount := list.totalCount()
	var data bytes.Buffer

	data.WriteString(fmt.Sprintf("文章总单词数：%d\n\n", wordsCount))
	for _, v := range list {
		var percent float64 = 100.0 * float64(v.count) / float64(wordsCount)
		_, err := data.WriteString(fmt.Sprintf("%s: %d, %3.2f%%\n", v.word, v.count, percent))
		CheckError(err, "bytes.Buffer, WriteString")
	}

	err := ioutil.WriteFile("C:/samba/go-work/src/statistical/wordcount.txt", []byte(data.String()), os.ModePerm)
	CheckError(err, "ioutil.WriteFile")

	return  wordMap

}

type Weibo struct {
	Id           uint64 `json:"id"`
	Timestamp    uint64 `json:"timestamp"`
	UserName     string `json:"user_name"`
	RepostsCount uint64 `json:"reposts_count"`
	Text         string `json:"text"`
}

var (
	segmenter sego.Segmenter
	itemwordmap map[string]int
)

func keyOvalue (itemwordmap map[string]int) map[int]string{

	var itemwordmap2 map[int]string
	for key, value := range itemwordmap {
		itemwordmap2=key
		itemwordmap2=value
	}
	return itemwordmap2

}


func ItemSort(wbs map[uint64]Weibo,wordMapSort map[string]int) []string{

	var text []byte

	for _, weibo := range wbs {
		text = []byte(weibo.Text)
		itemwordmap=segmenter.Segment(text)//事务分词，并将项目集放入itemwordmap中


		//对事务内项目进行赋值int
		for key, value := range itemwordmap {

			value=wordMapSort(key)
			itemwordmap(key)=value
		}//进行完这一步循环之后，事务中的每个项都有一个正确的频率值

		//对项频与项值调换，因为你没有办法通过值找到键，只能通过键找到值
		itemwordmap2:=keyOvalue(itemwordmap)

		//对频率进行排序
		sort.Ints(itemwordmap2)

		//按照频率顺序，将项依次放入一个字符串数组之中
        var itemsortslicestring []string
		for _, k := range itemwordmap2 {
			itemsortslicestring =itemwordmap2[k]
		}
		return itemsortslicestring

	}
}
*/



