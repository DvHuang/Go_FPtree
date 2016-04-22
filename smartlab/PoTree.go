package fpTree

import (
	"fmt"
	"sort"
	"statistical/PublicStruct"
)

const (
	minSup=10
)

var(
	HeaderTable map[string]*TreeNode
	RetTree TreeNode
)

/*type STtable struct  {
	itemcount int
	nodelink *treeNode
}*/
type TreeNode struct {
	itemname string
	count    int
	parent   *TreeNode
	child 	 map[string]*TreeNode
	brother  *TreeNode
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
	if list[i].count < list[j].count {
		return true
	} else if list[i].count > list[j].count {
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


//事务集合为一个map，键位一条事务的字符串，对应值为事务中的项，及其对应频率

func CreateFPTree (allitemsets map[string]PublicStruct.MapPrefixPath, minSup int,)( TreeNode, map[string]*TreeNode){


	/*if HeaderTable==nil {
		HeaderTable = make(map[string]*TreeNode)
	}*/
	HeaderTable=make(map[string]*TreeNode)
	var bukenengba TreeNode
	bukenengba.itemname="CreateFPTree"
	HeaderTable["CreateFPTree"]=&bukenengba
	fmt.Println("fast test",HeaderTable["CreateFPTree"])

	//HeaderTable2:=make(map[string]*TreeNode)
	//数据结构解释：allitemset->事务集合 itemsets--> 事务（项集合）
	// first pass counts frequency of occurance 遍历事务集合（遍历一个map），slicevalue（map的一个元素）为一条事务
	for _,slicevalue :=range allitemsets {
		//fmt.Println("slicevalueslicevalueslicevalueslicevalue",slicevalue)
		//对一条事务（map结构）遍历，返回值key为项名称，value为该项出现次数。把每个项的出现次数放入map-->headertable中
		for key,value := range slicevalue.Path {
			//fmt.Println(key,value)

			if _,ok:=HeaderTable[key];ok{
				//如果存在该键值
				//HeaderTable[key].count=HeaderTable[key].count + value
				var sttable TreeNode
				sttable.count=value+HeaderTable[key].count
				sttable.itemname=HeaderTable[key].itemname
				HeaderTable[key]=&sttable
			}else{
				var sttable TreeNode
				sttable.count=value
				sttable.itemname=key
				HeaderTable[key]=&sttable
				//HeaderTable[key].count=value

			}

		}
	}
	//fmt.Println("1111111fmt.Println(HeaderTable,HeaderTable)",HeaderTable)
	for key,value :=range HeaderTable{
		//fmt.Println(HeaderTable[key])  输出&{1<nil>}
		//fmt.Println(HeaderTable[key].itemcount) 输出整数
		if value.count < minSup {
			delete(HeaderTable,key)
		}
	}
	//fmt.Println("2222222222fmt.Println(HeaderTable,HeaderTable)",HeaderTable)
	if len(HeaderTable) == 0{
		//fmt.Println("all item have no enough support")
		panic("all item have no enough support,defer Exit!!!!")
	}  //#if no items meet min support -->get out


	var RetTree TreeNode//注意之类的大括号
	//var RetTree *treeNode=new(treeNode)

	//RetTree.child["root"] = new(treeNode)

	//var RetTree *treeNode 进行打印   		fmt.Println("----------",RetTree.itemname)：：：   这一句将会报错
	//var RetTree *treeNode=new(treeNode)   fmt.Println("----------",RetTree.itemname)：：：   打印出来是------ 空
	//var intree  treeNode打印结果          fmt.Println("intree",intree）             ：：：   intree { 0 <nil> map[] <nil>}

	//new返回了一个指针，该指针的内容为空
	/*fmt.Println("retree and &retree",RetTree,&RetTree)

	var intree  TreeNode
	fmt.Println("intree",intree)
	fmt.Println("intree.itemname",intree.itemname)
	fmt.Println("intree.itemname",intree.child)
	fmt.Println("intree.itemname",intree.brother)
	fmt.Println("intree.itemname",intree.count)
	fmt.Println("intree.itemname",intree.parent)

	fmt.Println("----------",RetTree)*/

	//对每条事务中的项进行排序  //go through dataset 2nd time //#put transaction items in order

	for _,slicevalue :=range allitemsets {   // first pass counts frequency of occurance 遍历事务集合，slicevalue为一条事务
		//fmt.Println("一条事务",slicevalue)
		var pList2 PersonList2
		for key, _ := range slicevalue.Path {   //对一条事务（map结构[string]int）遍历，
			//对事务中的项进行赋值
			if _, ok := HeaderTable[key]; ok {
				//如果改建在headertable中也存在的话，那么判断结构体中有没有该键
				var itemGd Person
				itemGd.age = HeaderTable[key].count
				itemGd.name = key
				pList2 = append(pList2, itemGd)//一条plist是一条事务

			}

		}//进行完这一步循环之后，事务中的每个项都有一个正确的频率值

		//对一条事务中的项，按照频率进行排序
		sort.Sort(pList2)

		//按照频率顺序，将项依次放入一个字符串数组之中
		//var orderedItems []string
		orderedItems := make([]string,3,5)
		//这里要注意数组与切片之间的却别，上面的定义之后就是数组，下面是切片，


		for _, v := range pList2 {
			orderedItems = append(orderedItems, v.name)
		}
		//fmt.Println(orderedItems)

		//一个orderitems 就是一条排好序的项目集--->即一条事务-->在数据结构上看是一个字符串数组
		//headertable 的数据结构为map[string]STtable，STtable中包含count频率，和该项在fp中的第一个位置的地址
		//fmt.Println("orderedItems----------",orderedItems)
		//RetTree.child=make(map[string]*TreeNode)
		//HeaderTable2=
		updateTree(orderedItems, &RetTree, HeaderTable, 10)//#populate tree with ordered freq itemset
	}

	return RetTree, HeaderTable //#return tree and header table
}
//items 降序项集-->事务，intree 树根，链表头
func updateTree(orderedItems []string, inTree  *TreeNode, HeaderTable map[string]*TreeNode, count int) {//map[string]*TreeNode{

	HeaderTabletemp :=make(map[string]*TreeNode)
	HeaderTabletemp=HeaderTable
	//check if orderedItems[0] in RetTree.children
	//inTree.child 是一个map，直接看这个map中是否存在items【0】这个键值，存在频率加一，不存在

	//HeaderTable=make(map[string]*TreeNode)

	var bukenengba TreeNode
	bukenengba.itemname="updateTree"
	HeaderTable["updatetree"]=&bukenengba

	if inTree.child==nil {
		inTree.child=make(map[string]*TreeNode)
	}

	//fmt.Println("HeaderTable----------",HeaderTable)
	//fmt.Println("orderedItems[0]----------",orderedItems[0])

	//HeaderTable=make(map[string]*STtable)
	//fmt.Println("----------",inTree.child)
	//items=make([]string,len(items))
	//fmt.Println("----------",items)

	/*var sttable treeNode
	sttable.count+=1
	sttable.itemname=items[0]
	inTree.child["nihao"]=&sttable	//incrament count*/
	if _, ok := inTree.child[orderedItems[0]]; ok{
		//fmt.Println("查看inTree.child这个map中是否存在items[0]这个键值，存在频率加一",items[0])

		var sttable TreeNode

		//map 只能整体赋值
		sttable.count=1+inTree.child[orderedItems[0]].count//节点频率
		sttable.itemname=orderedItems[0]
		sttable.brother=inTree.child[orderedItems[0]].brother
		sttable.parent=inTree.child[orderedItems[0]].parent

		inTree.child[orderedItems[0]]=&sttable	//incrament count


	}else {
		//child中不存在items【0】这个键值，就添加一个节点，该节点的父节点parent为intree，
		// 节点信息包括name，count，child照样指向一个map
		//fmt.Println("child map中不存在items[0]这个键值，就添加一个节点",items[0])
		var sttable TreeNode
		sttable.count=1
		sttable.itemname=orderedItems[0]
		//sttable.brother=inTree 新添加的节点是没有brother的，更新brother只能在下面 upheader
		sttable.parent=inTree
		//fmt.Println("sttable",sttable)
		//fmt.Println("intree.child",inTree.child[items[0]])
		inTree.child[orderedItems[0]]=&sttable


		/*var sttablehead TreeNode
		fmt.Println("HeaderTabletemp",HeaderTabletemp[orderedItems[0]].count)
		sttablehead.count=HeaderTabletemp[orderedItems[0]].count//链表头总体频率
		sttablehead.brother=&sttable//新添加节点的地址要放到链表头指针中
		HeaderTable[orderedItems[0]]=&sttablehead*/

		/*fmt.Println("updateTree sttable brother",sttable.brother)
		fmt.Println("updateTree sttable parent",sttable.parent)
		fmt.Println("updateTree HeaderTable[orderedItems[0]] brother",HeaderTable[orderedItems[0]].brother)
		fmt.Println("updateTree HeaderTable[orderedItems[0]] parent",HeaderTable[orderedItems[0]].parent)*/
		//child【items【0】】这个键对应的值是一个节点的地址

		//询问头表中相应的键值是否存在-->即链表头是否存在该项，
		// 存在，则添加到横向链表末尾节点的brother
		if _,ok:=HeaderTable[orderedItems[0]];ok{  //update header table
			/*if value.parent==nil {
				//纵向链表中不存在该节点
				//fmt.Println("链表头中存在该项，添加到最后",inTree.child[orderedItems[0]])
				var sttable TreeNode
				sttable.count=HeaderTable[orderedItems[0]].count
				sttable.parent=inTree.child[orderedItems[0]]
				HeaderTable[orderedItems[0]].parent=&sttable

			}else {
				//纵向链表中存在该节点，跟新到最后
				updateHeader(HeaderTable[orderedItems[0]], inTree.child[orderedItems[0]])

				//如果定义为TreeNode，上面打印出来的是{ 1 0xc0cfe3c540 map[] <nil>}
				//如果定义为*TreeNode 上面打印出来的是&{ 1 0xc0cfe3c540 map[] <nil>}

				//这里体现了go的人性化，他告诉你这是个地址，还有地址中的值是什么，而没有直接输出一个毫无意义的地址
				//HeaderTable[orderedItems[0]]=inTree.child[orderedItems[0]]
				//HeaderTable[orderedItems[0]].brother = inTree.child[orderedItems[0]]//理想上，此值代表一个项的地址（即一个节点的地址）
			}*/
			updateHeader(HeaderTable[orderedItems[0]], inTree.child[orderedItems[0]])
		}else{//不存在，则纵向添加新节点
			//updateHeader(HeaderTable[orderedItems[0]], inTree.child[orderedItems[0]])


			var sttablehead TreeNode
			if _,ok:=HeaderTabletemp[orderedItems[0]];ok{
				sttablehead.count=HeaderTabletemp[orderedItems[0]].count//链表头总体频率
			}else{
				sttablehead.count=1+inTree.child[orderedItems[0]].count
			}
			sttablehead.brother=inTree.child[orderedItems[0]]//新添加节点的地址要放到链表头指针中
			HeaderTable[orderedItems[0]]=&sttablehead

			//HeaderTable[orderedItems[0]]=inTree.child[orderedItems[0]]
			//HeaderTable[orderedItems[0]].brother = inTree.child[orderedItems[0]]//理想上，此值代表一个项的地址（即一个节点的地址）

		}
	}
	//fmt.Println("&inTree.child[items[0]]======",inTree.child["nihao"])
	if len(orderedItems) > 1 {//call updateTree() with remaining ordered items
		//orderedItems=orderedItems[1:]
		updateTree(orderedItems[1:], inTree.child[orderedItems[0]], HeaderTable, 10)
	}
	//return  HeaderTable

}
(HeaderTable[orderedItems[0]], inTree.child[orderedItems[0]])
func updateHeader(nodeToTest *TreeNode, targetNode *TreeNode) {//this version does not use recursion
	//fmt.Println("链表头中存在该项，添加到最后",nodeToTest)
	//链表更新，
	//上一步判断在链表头中存在该元素，那就遍历整个横向链表节点的兄弟节点，
	// 直到为空时说明到了末尾，添加这个节点到横向链表末尾节点的brother
	for nodeToTest.brother != nil {//Do not use recursion to traverse a linked list!
		nodeToTest.brother = targetNode.brother  //！=nil说明headertable中已经存在该项的头指针，不需要跟新
		fmt.Println(".")
	}
	//nodeToTest.brother = targetNode

}

func ascendTree(leafNode *TreeNode, prefixPath []string) []string{// #ascends from leaf node to root
	fmt.Println("ascendTree TreeNode",leafNode)
	if leafNode.parent != nil {//上朔，只要父节点存在，就将父节点的名字放到字符串数组里
		prefixPath = append(prefixPath, leafNode.itemname)
		ascendTree(leafNode.parent, prefixPath)
	}
	fmt.Println("ascendTree prefixPath",prefixPath)
	return prefixPath
}

/*type  StrPrefixPath struct {
	path	 []string
	count  	 int
}*/

func FindPrefixPath(basePat string, treeNode  *TreeNode ) map[string]PublicStruct.MapPrefixPath {//treeNode comes from header table
	//basePat 为链表头排序之后的一个节点-->项名称，上朔，找集合
	//basePet表示输入的频繁项，treeNode为当前FP树中对应的第一个节点
	condPats:=make(map[string]PublicStruct.MapPrefixPath,10)
	//函数返回值即为条件模式基condPats，用一个字典表示，键为前缀路径合成的字符串，值为前缀路径结构体，包括项集和对应频率。

//	fmt.Println("FindPrefixPath treeNode.brother",treeNode.brother)

	//prefixPath:=make([]string,10)
	//prefixPath=ascendTree(treeNode, prefixPath)//把要追朔的节点地址，即要放入路径的字符串切片给ascend函数

	for treeNode!=nil{ //brother 为headtable中所有相同元素项的地址
		prefixPath:=make([]string,10)
		prefixPath=ascendTree(treeNode, prefixPath)//把要追朔的节点地址，即要放入路径的字符串切片给ascend函数
		prefixPath=append(prefixPath,ascendTree(treeNode, prefixPath)...)//把要追朔的节点地址，即要放入路径的字符串切片给ascend函数

		if len(prefixPath) > 1{  //前缀路径中存在一个以上的元素，对前缀路径进行切片，以及该支持度处理
			var allvalue string
			mapfixPath:=make(map[string]int)
			for _,value:=range prefixPath[1:]{
				allvalue+=value//前缀路径元素组成一个字符串，作为键值
				mapfixPath[value]=treeNode.count//切片改map赋值给下面的路径
			}
			fmt.Println("FindPrefixPath mapfixPath ",mapfixPath)
			fmt.Println("allvalue+=value ",allvalue)

			//condPats[allvalue].count = treeNode.count//将前缀路径上的支持度改为底部的支持度

			//allitemsets[stringMapPrefixPath].Path[key] =value
			condPatsPath, _ := condPats[allvalue]

			if condPatsPath.Path ==nil{
				condPatsPath.Path=make(map[string]int)
			}

			for key,value:=range mapfixPath {
				//condPats[allvalue].count[key] = treeNode.count//将前缀路径上的支持度改为底部的支持度
				//condPats[allvalue].Path[key] = value //
				condPatsPath.Path[key]=value
			}
			condPats[allvalue] = condPatsPath
		}
		treeNode = treeNode.brother
	}
	fmt.Println("FindPrefixPath condPats ",condPats)
	return condPats
}

func MineTree(inTree TreeNode, HeaderTable  map[string]*TreeNode, minSup int, preFix []string, freqItemList [][]string) {

	//freqItemList 频繁项集
	//preFix 作为前缀路径
	//输入为根节点和链表头
	//fmt.Println("header test", HeaderTable["CreateFPTree"].itemname)
	//fmt.Println("header test", HeaderTable["updatetree"].itemname)

	if _,ok:= HeaderTable["CreateFPTree"];ok {
		fmt.Println("header test", HeaderTable["CreateFPTree"].itemname)
	}
	if _,ok:= HeaderTable["updatetree"];ok {
		fmt.Println("header test", HeaderTable["updatetree"].itemname)
	}
	//先对链表头进行排序，然后按照顺序进行递归发现频繁项集newfreset，将每一个newfreset，放入freItemList中作为最后的输出
	var  bigL PersonList3
	var itemGd ListHeaderTable
	for key, value := range HeaderTable {
		/*fmt.Println("MineTree HeaderTable itemname",key,value.itemname)
		fmt.Println("MineTree HeaderTable brother",key,value.brother)
		fmt.Println("MineTree HeaderTable count",key,value.count)
		fmt.Println("MineTree HeaderTable child",key,value.child)
		fmt.Println("MineTree HeaderTable parent",key,value.parent)*/



		itemGd.count= value.count
		itemGd.itemname = key
		itemGd.brother=value.brother
		//fmt.Println("itemGd.brother=value.brother",itemGd.brother)
		//itemGd.child=value.child[key]
		/*if _,ok:=HeaderTable[key];ok {
			itemGd.self = HeaderTable[key]
		}*/

	}
	bigL = append(bigL, itemGd)

	sort.Sort(bigL)   //(sort header table)

	//按照顺序对链表头进行遍历

	for _,value:=range bigL {
		//start from bottom of header table 从指针头表的底端开始遍历
		//newFreqSet := append(preFix, value.name)//newFreqSet = preFix.copy()
		//newFreqSet.add(basePat)
		//#print 'finalFrequent Item: ', newFreqSet    #append to set
		//var newFreqSet []string//一个频繁项集，其大小不会超过所有项集合的长度

		newFreqSet := make([]string,3,5)
		newFreqSet=append(newFreqSet,preFix...)
		newFreqSet=append(newFreqSet,value.itemname)

		freqItemList=append(freqItemList,newFreqSet)

		fmt.Println("for _,value:=range bigL---one bigL ",value.brother)
		condPattBases := FindPrefixPath(value.itemname,value.brother)
		fmt.Println("condPattBases",condPattBases)

		//#print 'condPattBases :', basePat, condPattBases
		//#2. construct cond FP - tree from cond.pattern base
		//var newHeaderTable map[string]*TreeNode
		fmt.Println("频繁项集",freqItemList)

		/*if condPattBases==nil {
			panic("condPattBases")
		}*/
		defer fmt.Println("频繁项集",freqItemList)
		myCondTree, myHead := CreateFPTree(condPattBases, minSup)

		//#print 'head from conditional tree: ', myHead
		if myHead != nil {
			//3. mine cond.FP - tree
			//#print 'conditional tree for: ', newFreqSet
			//#myCondTree.disp(1)
			MineTree(myCondTree, myHead, minSup, newFreqSet, freqItemList)
		}else{
			fmt.Println("频繁项集",freqItemList)
			fmt.Println("myHead,为零了，exit",myHead)
		}
	}
}



