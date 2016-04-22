
package main

import (
	"fmt"
	"statistical/fpTree"
	"statistical/textCpu"
)

const minsupport = 1

func main() {

	colors := map[uint64]string{
		1:   "面包牛奶",
		2:       "面包尿布啤酒鸡蛋",
		3:    "牛奶尿布啤酒可乐",
		4: "面包牛奶尿布啤酒",
		21:   "面包牛奶尿布可乐",

	}

	allitemsets := textCpu.WbToPyData(colors)
	fmt.Println("WbToPyData finish")


	fpTree.HeaderTable = make(map[string]*fpTree.HeadStruct)
	tree, head := fpTree.CreateFPTree(allitemsets, fpTree.HeaderTable, minsupport)


	var preFix []string
	freqItemList := make([][]string, 1)
	fpTree.MineTree(tree, head, minsupport, preFix, freqItemList)
	fmt.Println("MineTree finish")

}


