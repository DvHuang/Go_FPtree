/*
package textCpu

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"
)

//简单的词频统计任务
func CountTestBase(inputFilePath string, outputFilePath string) {
	//时间开始点
	start := time.Now().UnixNano() / 1e6

	//读取文件
	fileData, err := ioutil.ReadFile(inputFilePath)
	CheckError(err, "read file")
	var fileText string = string(fileData)

	//根据CPU核数新开协程
	newRountineCount := runtime.NumCPU()*2 - 1
	runtime.GOMAXPROCS(newRountineCount + 1)
	//切分文件
	parts := splitFileText(fileText, newRountineCount)

	var ch chan map[string]int = make(chan map[string]int, newRountineCount)
	for i := 0; i < newRountineCount; i++ {
		go countTest(parts[i], ch)
	}

	//主线程接收数据
	var totalWordsMap map[string]int = make(map[string]int, 0)

	completeCount := 0
	for {
		receiveData := <-ch
		for k, v := range receiveData {
			totalWordsMap[strings.ToLower(k)] += v
			//fmt.Printf("for totalwordsmap is = (%v)\n",totalWordsMap)
			//fmt.Printf("for k is = (%v)\n",k)
			//fmt.Printf("for v is = (%v)\n",v)
		}

		completeCount++

		if newRountineCount == completeCount {
			break
		}
	}

	//添加进slice，并排序
	list := make(WordCountBeanList, 0)
	//fmt.Printf("totalwordsmap is = (%v)\n",totalWordsMap)
	for k, v := range totalWordsMap {
		list = append(list, NewWordCountBean(k, v))
	}
	sort.Sort(list)
	//时间结束点
	end := time.Now().UnixNano() / 1e6
	fmt.Printf("time consume:%dms\n", end-start)

	//输出
	wordsCount := list.totalCount()
	var data bytes.Buffer
	data.WriteString(fmt.Sprintf("程序执行：%dms\n", end-start))
	data.WriteString(fmt.Sprintf("文章总单词数：%d\n\n", wordsCount))
	for _, v := range list {
		var percent float64 = 100.0 * float64(v.count) / float64(wordsCount)
		_, err := data.WriteString(fmt.Sprintf("%s: %d, %3.2f%%\n", v.word, v.count, percent))
		CheckError(err, "bytes.Buffer, WriteString")
	}

	err = ioutil.WriteFile(outputFilePath, []byte(data.String()), os.ModePerm)
	CheckError(err, "ioutil.WriteFile")
}

func countTest(text string, ch chan map[string]int) {
	var wordMap map[string]int = make(map[string]int, 0)

	//按字母读取，除26个字母（大小写）之外的所有字符均认为是分隔符
	startIndex := 0
	letterStart := false
	for i, v := range text {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			if !letterStart {
				letterStart = true
				startIndex = i
			}
		} else {
			if letterStart {
				wordMap[text[startIndex:i]]++
				letterStart = false
			}
		}
	}

	//最后一个单词
	if letterStart {
		wordMap[text[startIndex:]]++
	}
	//fmt.Printf("wordsmap is = (%v)\n",wordMap)

	ch <- wordMap
}

//将全文分成n段
func splitFileText(fileText string, n int) []string {
	length := len(fileText)
	parts := make([]string, n)

	lastPostion := 0
	for i := 0; i < n-1; i++ {
		position := length / n * (i + 1)
		for string(fileText[position]) != " " {
			position++
		}

		parts[i] = fileText[lastPostion:position]
		lastPostion = position
	}

	//最后一段
	parts[n-1] = fileText[lastPostion:]
	return parts
}

func CheckError(err error, msg string) {
	if err != nil {
		panic(msg + "," + err.Error())
	}
}
type WordCountBean struct {
	word  string
	count int
}

func NewWordCountBean(word string, count int) *WordCountBean {
	return &WordCountBean{word, count}
}

type WordCountBeanList []*WordCountBean

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

func (list WordCountBeanList) totalCount() int {
	totalCount := 0
	for _, v := range list {
		totalCount += v.count
	}

	return totalCount
}
*/
