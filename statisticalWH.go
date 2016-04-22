package main

import (
	"fmt"
	"statistical/fpTree"
	"statistical/textCpu"
)

const minsupport = 500

func main() {

	//读入数据得到wb结构体
	wbs := textCpu.InitWbData()
	fmt.Println("InitWbData finish")

	/*所有事务集合为一个外部map，外部map键值对应一个结构体。
	type  MapPrefixPath struct{
		Path    map[string]int   	结构体元素path--是一个map类型的项集合
		Count 	int  				结构体元素int---是该路径的支持度或者
	}wbs进入，得到需要的数据结构map[string]PublicStruct.MapPrefixPath*/
	allitemsets := textCpu.WbToPyData(wbs)
	fmt.Println("WbToPyData finish")

	//事务集合进入，得到树根节点和链表头
	fpTree.HeaderTable = make(map[string]*fpTree.HeadStruct)
	tree, head := fpTree.CreateFPTree(allitemsets, fpTree.HeaderTable, minsupport)
	fmt.Println("CreateFPTree finish")


	//递归查找频繁项集
	var preFix []string
	var  freqItemList [][]string
	fpTree.MineTree(tree, head, minsupport, preFix, freqItemList)
	fmt.Println("MineTree finish")

}
