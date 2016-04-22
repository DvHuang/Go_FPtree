package main

import (
	"fmt"
	"sort"
)

type IMessage interface {
	Print()
}

type BaseMessage struct {
	msg []byte
}

func (message BaseMessage) Print() {
	fmt.Println("baseMessage:msg", message.msg)
	for i, j := range message.msg {
		fmt.Println("SubMessage:msg", i, j)
	}
}

type SubMessage struct {
	msg []string
}

func (message SubMessage) Print() {

	fmt.Println("SubMessage:msg", message.msg)
	for i, j := range message.msg {

		fmt.Println("SubMessage:msg", i, j)

	}
}

type doublebyte struct {
	msg [][]byte
}

func (message doublebyte) Print() {

	fmt.Println("SubMessage:msg", message.msg)
	for i, j := range message.msg {

		fmt.Println("SubMessage:msg", i, j)

	}
}

func main() {

	//实验一：接口的操作
	//var message IMessage

	baseMessage := new(BaseMessage)
	baseMessage.msg = []byte{'a', 'b', 'c', 'f', 'g', ' '}
	baseMessage.Print()

	//message = baseMessage
	//message.Print()

	SubMessage := new(SubMessage)
	SubMessage.msg = []string{"abcdkuiej", "你好", "我很好啊"}
	SubMessage.Print()

	//message = SubMessage
	//message.Print()

	//实验二：map操作
	//map的声明和使用
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, _ := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, colors[key])
	}

	//map 的使用2
	var item string
	item = "Red"
	headertable := map[string]string{} //注意之类的大括号
	headertable["Red"] = "#da1337"
	fmt.Printf("???%s\n", headertable[item])

	//实验三：切片操作
	//v_IntSlice := []int{1, 2, 5,5,2,5,9,22: 10}
	v_stringSlice := []string{2: "ab", 3: "cd", 5: "ef", "iu", 7: "in"}
	//v_IntSlice[0] = 100
	//
	textSliceToString(v_stringSlice)

	//实验四，结构体的操作
	//对一个结构体赋值而不指明是哪个成员，会得到什么
	type test struct {
		name  string
		count int
		next  *test
	}

	//实验五：对map行排序
	//实验证明，无法对map进行排序，因为swap这个函数这里，无法对两个元素进行交换，或者说无法通过“位置”制定map中的一个元素

	/*	rdOneItem := ordOneItem{
		"AliceBlue":   2,
		"Coral":       3,
		"DarkGray":    5,
		"ForestGreen": 7,
	}*/

	rdOneItem := ordOneItem{
		3:  2,
		4:  3,
		5:  5,
		1:  7,
		11: 2,
		14: 3,
		15: 5,
		21: 7,
	}
	fmt.Println("len rdoneitem", len(rdOneItem))
	fmt.Println("onemap", rdOneItem)
	sort.Sort(rdOneItem)
	fmt.Println("onemap", rdOneItem, len(rdOneItem))

	for key, value := range rdOneItem {
		fmt.Println("maporder", key, value)
	}

}

type ordOneItem map[int]int

func (list ordOneItem) Len() int {
	return len(list)
}

func (list ordOneItem) Less(i, j int) bool {
	if list[i] > list[j] {
		return true
	} else if list[i] < list[j] {
		return false
	} else {
		return list[i] < list[j]
	}
}

func (list ordOneItem) Swap(i, j int) {
	var temp ordOneItem = list[i]
	list[i] = list[j]
	list[j] = temp
}

//实验二：切片也可以用rang，那切片返回什么呢

func textSliceToString(text []string) {
	//var output string
	for number, word := range text {
		//output += string(word)
		fmt.Println(word, number)
	}
	//return output
}

//实验三：对一个结构体进行赋值，会得到什么
