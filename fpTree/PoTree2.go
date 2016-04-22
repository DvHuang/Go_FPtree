package fpTree

import (
	"fmt"
	"sort"
	"statistical/PublicStruct"
)

var (
	HeaderTable map[string]*HeadStruct
	Updatecount     int
	TreeNumberC 	int
	EndOutFor		int

)

type TreeNode struct {
	itemname string
	count    int
	parent   *TreeNode
	child    map[string]*TreeNode
	brother  *TreeNode
}
type HeadStruct struct {
	headcount    int
	NodeLink  *TreeNode
}
type Person struct {
	name string
	age  int
}
type PersonList []*Person
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

type ListHeaderTable struct {
	itemname string
	count    int
	child    *TreeNode
	brother  *TreeNode
	self     *TreeNode
}
type PersonList3 []ListHeaderTable

func (list PersonList3) Len() int {
	return len(list)
}
func (list PersonList3) Less(i, j int) bool {
	if list[i].count > list[j].count {
		return true
	} else if list[i].count < list[j].count {
		return false
	} else {
		return list[i].itemname < list[j].itemname
	}
}
func (list PersonList3) Swap(i, j int) {
	var temp ListHeaderTable = list[i]
	list[i] = list[j]
	list[j] = temp
}

func CreateFPTree(allitemsets map[string]PublicStruct.MapPrefixPath, headCF map[string]*HeadStruct, minSup int) (TreeNode, map[string]*HeadStruct) {

	headCF = make(map[string]*HeadStruct)
	TreeNumberC+=1
	/*fmt.Printf("第------%v------棵树开始\n",TreeNumberC)
	if TreeNumberC>1 {
		fmt.Printf("第--%v--个集合，长度=%v\n",TreeNumberC,len(allitemsets))
		//fmt.Println("allitemsets finish===",allitemsets)
	}else if TreeNumberC==1{
		fmt.Printf("第--%v--个集合，长度=%v\n",TreeNumberC,len(allitemsets))
		//fmt.Println("allitemsets finish===",allitemsets)
	}*/


	//遍历集合
	for _, slicevalue := range allitemsets {
		for key, value := range slicevalue.Path {
			if _, ok := headCF[key]; ok {
				//该键值已存在
				var sttable HeadStruct
				sttable.headcount = value + headCF[key].headcount

				headCF[key] = &sttable
			} else {
				var sttable HeadStruct
				sttable.headcount = value

				headCF[key] = &sttable
			}
		}
	}
	//支持度
	for key, value := range headCF {
		if value.headcount < minSup {
			delete(headCF, key)
		}
	}

	var CreateFPTreeRet TreeNode
	var pCreateFPTreeRet *TreeNode

	if len(headCF) == 0 {
		//fmt.Printf("第--%v--个项集已完，\n退回到上一层循环，\n寻找另外一个底基项\n",TreeNumberC-1)
		return  CreateFPTreeRet,nil
	}



	HeadUD := make(map[string]*HeadStruct)
	//排序
	for _, slicevalue := range allitemsets {
		var pList2 PersonList2
		for key, _ := range slicevalue.Path {
			if _, ok := headCF[key]; ok {
				var itemGd Person
				itemGd.age = headCF[key].headcount
				itemGd.name = key
				pList2 = append(pList2, itemGd)
			}
		}
		if len(pList2)<2{
			continue
			fmt.Println("空 pList2 0000000,pass！！！！！",pList2)
		}
		sort.Sort(pList2)
		orderedItems := make([]string, 1)

		for _, v := range pList2 {
			if v.name==""{
				break
			}
			orderedItems = append(orderedItems, v.name)
		}
		pCreateFPTreeRet, HeadUD = updateTree(orderedItems, HeadUD, &CreateFPTreeRet)
		CreateFPTreeRet = *pCreateFPTreeRet
	}

	//fmt.Printf("第------%v------棵树结束\n",TreeNumberC)
	return CreateFPTreeRet, HeadUD
}

//建立itemheadtable
func updateTree(orderedItems []string, HeadUD map[string]*HeadStruct, RetUd *TreeNode) (*TreeNode, map[string]*HeadStruct) {

	Updatecount+=1

		//fmt.Printf("len(orderedItems)=%v,\nUpdatecount=%v\n ", len(orderedItems), Updatecount)
		if RetUd.child == nil {
			RetUd.child = make(map[string]*TreeNode)
		}
		//ret & head
		var OldChCoP TreeNode                            //old child count plus
		var NeChAdTM TreeNode                        //new child add to map
		var OlHeNoCoPl HeadStruct                        //old head node count plus
		var NeHeNoAdBr HeadStruct                            //New Head Node Add Brother

		if value, ok := RetUd.child[orderedItems[0]]; ok {
			OldChCoP.itemname = orderedItems[0]
			OldChCoP.count = RetUd.child[orderedItems[0]].count + 1
			OldChCoP.brother = RetUd.child[orderedItems[0]].brother
			OldChCoP.parent = RetUd.child[orderedItems[0]].parent
			RetUd.child[orderedItems[0]] = &OldChCoP

			//child存在，head肯定存在，+1即可
			OlHeNoCoPl.headcount = value.count + 1
			OlHeNoCoPl.NodeLink = value.brother
			HeadUD[orderedItems[0]] = &OlHeNoCoPl

		} else {//横向无，添加新节点，必更新brother（head or ret 中）
			NeChAdTM.itemname = orderedItems[0]
			NeChAdTM.count = 1
			NeChAdTM.brother = nil
			NeChAdTM.parent = RetUd
			RetUd.child[orderedItems[0]] = &NeChAdTM

			//横向child中没有，但纵向head中有，说明该项在其他节点中出现过，找到这个节点，给他的brother赋值
			if value, ok := HeadUD[orderedItems[0]]; ok {
				OlHeNoCoPl.headcount = value.headcount + 1
				OlHeNoCoPl.NodeLink = value.NodeLink

				//通过head和brother递归，找到横向链表最后一个节点
					findbrother:=HeadUD[orderedItems[0]].NodeLink
					for findbrother.brother != nil {
						findbrother = findbrother.brother
						//fmt.Println("...",findbrother)
					}
					findbrother.brother= &NeChAdTM    //这个findbrother.brother 指向了最后一个没有兄弟节点的TreeNode

				HeadUD[orderedItems[0]] = &OlHeNoCoPl
			} else {
				//横向没有，纵向也没有，上面添加了ret新节点，这里给head更新brother
				NeHeNoAdBr.headcount = 1
				NeHeNoAdBr.NodeLink = &NeChAdTM
				HeadUD[orderedItems[0]] = &NeHeNoAdBr
			}
		}
	if len(orderedItems) > 1 {

			updateTree(orderedItems[1:], HeadUD, RetUd.child[orderedItems[0]], )

	}

	return RetUd, HeadUD
}
func ascendTree2(PbmNode *TreeNode, prefixPath []string) []string { // #ascends from leaf node to root
	if PbmNode.parent != nil {
		prefixPath = append(prefixPath, PbmNode.itemname)


		fmt.Println(" PbmNode.itemname  ",PbmNode.itemname)
		fmt.Println(" len prefixPath ",len(prefixPath))
		fmt.Println(" prefixPath ",prefixPath)

		ascendTree(PbmNode.parent, prefixPath)
	}
	return prefixPath
}
//这两个函数比较告诉我们，必须要有返回值才可以，没有返回值迭代进去之后改变的值，
// 不会被带到外边，你最终外边返回的数值，会是当初第一次执行最外围函数时的值
func ascendTree(PbmNode *TreeNode, prefixPath []string) []string { // #ascends from leaf node to root

	if  PbmNode.parent != nil {
		prefixPath = append(prefixPath, PbmNode.itemname)
		prefixPath=ascendTree(PbmNode.parent, prefixPath)
	}
	return prefixPath
}

func FindPrefixPath(basePat string, PbmNode *TreeNode) map[string]PublicStruct.MapPrefixPath {

	conddata := make(map[string]PublicStruct.MapPrefixPath)

	for PbmNode != nil {
		var prefixPath []string
		//prefixPath := make([]string,1)
		var asscendPath []string
		//fmt.Println(" 进入之前asscendPath==",asscendPath)
		asscendPath=ascendTree(PbmNode, prefixPath)
		//fmt.Println(" out asscendPath--------------",asscendPath)
		prefixPath = append(prefixPath, asscendPath...)

		//fmt.Println(" prefixPathprefixPath+++++++++++",prefixPath)
		if len(prefixPath) > 1 {
			var allvalue string
			mapfixPath := make(map[string]int)
			for _, value := range prefixPath[1:] {//最底端元素去掉
				if value==""{
					break
				}
				allvalue += value

				mapfixPath[value] = PbmNode.count  //每一个项，作为mapfixPath的键
			}
			condPatsPath, _ := conddata[allvalue]
			if condPatsPath.Path == nil {
				condPatsPath.Path = make(map[string]int)
			}
			for key, value := range mapfixPath {
				if key==""{
					break
				}
				condPatsPath.Path[key] = value     //每一个项，作为condPatsPath的键
			}

			conddata[allvalue] = condPatsPath		//做成一条事务
			//fmt.Println(" ++++++一条事务+++++",conddata[allvalue])
		}
		PbmNode = PbmNode.brother
	}
	//fmt.Println("所有事务condPats ",conddata)
	return conddata
}

func MineTree(mtRet TreeNode, MineTreehead map[string]*HeadStruct, minSup int, preFix []string, freqItemList [][]string)[][]string {


	//fmt.Printf("第------%v------个Head\n",TreeNumberC)
	//fmt.Println("Head:::::", MineTreehead)
	var Listhead PersonList3
	var itemGd ListHeaderTable
	for key, value := range MineTreehead {
		itemGd.count = value.headcount
		itemGd.itemname = key
		itemGd.brother = value.NodeLink
		Listhead = append(Listhead, itemGd)
		//fmt.Println(key,value)
	}
	sort.Sort(Listhead)
	//fmt.Printf("第------%v------个Listhead\n",TreeNumberC)
	//fmt.Println("bigL:::::", bigL)
	//bigL=bigL[1:]
	if TreeNumberC ==1{
		EndOutFor=len(Listhead)
		fmt.Printf("listhead=%v\nEnd=%v",len(Listhead),EndOutFor)
	}


	for _, value := range Listhead {
		var outsideLC,insideLC int
		var newFreqSet []string

		//fmt.Println(key,value)
		//newFreqSet := make([]string, 1, 5)
		newFreqSet = append(newFreqSet, preFix...)
		newFreqSet = append(newFreqSet, value.itemname)
		freqItemList = append(freqItemList, newFreqSet)

		//fmt.Printf("第------%v------阶项集\n",TreeNumberC)
		//fmt.Println("频繁项集", freqItemList)

		//fmt.Printf("前缀路径需求数据：name：：%v\nbrother::%v\n ",value.itemname, value.brother)
		nextdata := FindPrefixPath(value.itemname, value.brother)
		//fmt.Printf("第---%v---条前缀路径 ",TreeNumberC)
		//fmt.Println("nextdata---------::",nextdata)


		nextTree, nextHead := CreateFPTree(nextdata, MineTreehead, minSup)

		if nextHead != nil {
			outsideLC+=1
			freqItemList=MineTree(nextTree, nextHead, minSup, newFreqSet, freqItemList)
			//fmt.Printf(".....nextHead !!!!!!= nil....number=%v\n",outsideLC)
		}else{
			insideLC+=1
			EndOutFor=-1
			//fmt.Printf(".....nextHead ======= nil....number=%v\n",insideLC)
			//fmt.Printf("第......%v.......阶项集\n开始%v结束",TreeNumberC,freqItemList)
		}

	}
	if EndOutFor <3{

		//fmt.Printf("最终集合EndOutFor=%v\n>>>>>开始%v结束<<<<<<<",EndOutFor,freqItemList)

	}

	return freqItemList


}

