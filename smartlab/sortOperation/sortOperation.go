package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type PersonList []*Person

func (list PersonList) Len() int {
	return len(list)
}

func (list PersonList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	} else {
		return list[i].name < list[j].name
	}
}

func (list PersonList) Swap(i, j int) {
	var temp *Person = list[i]
	list[i] = list[j]
	list[j] = temp
}

type PersonList2 []Person

func (list PersonList2) Len() int {
	return len(list)
}
func (list PersonList2) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	} else {
		return list[i].name < list[j].name
	}
}

func (list PersonList2) Swap(i, j int) {
	var temp Person = list[i]
	list[i] = list[j]
	list[j] = temp
}
func main() {
	type ordOneItem map[string]int
	rdOneItem := ordOneItem{
		"AliceBlue":   2,
		"Coral":       3,
		"DarkGray":    5,
		"ForestGreen": 7,
	}

	var pList2 PersonList2
	var itemGd Person
	for key, value := range rdOneItem { //对一条事务（map结构[string]int）遍历，
		//对事务中的项进行赋值
		itemGd.age = value
		itemGd.name = key
		pList2 = append(pList2, itemGd)
	} //进行完这一步循环之后，事务中的每个项都有一个正确的频率值
	sort.Sort(pList2)
	fmt.Println(pList2)

	var orderedItems []string
	for _, v := range pList2 {
		orderedItems = append(orderedItems, v.name)
	}
	fmt.Println(orderedItems)

	type treeNode struct {
		itemname string
		count    int
		parent   *treeNode
		child    map[string]*treeNode
		brother  *treeNode
	}
	var retTree map[string]*treeNode
	fmt.Println(&retTree)

	var add *int
	fmt.Println(&add)

	ab := treeNode{itemname: "nihao", count: 11, child: retTree} //  #add items[0] to inTree.children
	fmt.Println("ab", ab)
	fmt.Println("ab", &ab)

}
