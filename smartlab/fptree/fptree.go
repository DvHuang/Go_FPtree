/*
package c

import (
	"fmt"
)

type fpnode struct {
	itemname []byte
	count    int
	parent   *fpnode
	child    *fpnode
	brother  *fpnode
}

func (fpnode *fpnode) funcfpnode(name []byte) *fpnode {
	fpnode.itemname = name
	fpnode.count = 0
	fpnode.parent = nil
	fpnode.child = nil
	fpnode.brother = nil
	return fpnode
}

type tablenode struct {
	itemname  []byte
	frequence int
	next      *tablenode
}

func (tablenode *tablenode) functablenode(name []byte) *tablenode {

	tablenode.itemname = name
	tablenode.frequence = 0
	tablenode.next = nil
	return  tablenode
}

func creattablelist(item []byte, i int, head *tablenode) tablenode {

	j := 0
	node := head

	if node.next == nil {
		for j,_ =range item {


				newnode := (*tablenode).functablenode(item[j])
				newnode.frequence += 1
				node.next = newnode
				node = newnode
				head.frequence++ //保存链表结点数
				j++


		}
	} else {
		for j,_ =range item {
			node = head.next
			for node != nil {
				if node.itemname == item[j] {
					node.frequence++
					break
				} else if node.next != nil {
					node = node.next
				} else {
					newnode := (*tablenode).functablenode(item[j])
					head.frequence++
					newnode.frequence++
					node.next = newnode
					break
				}
			}
			j++
		}
	}
	//delete node
	return *head
}

func insert(a []byte, data []byte, i int, root *fpnode) fpnode {

	True := 0

	midnode := root.child
	if midnode == nil {
		True = 2
	}

	for midnode != nil {

		if (*midnode).itemname == a {
			True = 1
			break
		} else {
			if midnode.brother != nil {
				midnode = midnode.brother
			} else {
				break
			}
		}
	}
	if True == 1 {
		(*midnode).count++
	}

	if True == 2 {
		newnode := (*fpnode).funcfpnode(a)
		(*newnode).count++
		newnode.parent = root
		root.child = newnode

		midnode = newnode
	}

	if True == 0 {
		newnode := (*fpnode).funcfpnode(a)
		newnode.count++
		newnode.parent = root
		midnode.brother = newnode
		midnode = newnode
	}

	if data[1] !=0 {
		a := data[1]
		m := i - 1
		for j := 0; j < m; j++ {
			data[j] = data[j+1]
		}

		insert(string(a), data, m, midnode)
	}
	return *root
	//delete midnode;
}

func bubblesort(head *tablenode) tablenode {

	var (
		m       int
		a       int
		newnode *tablenode
		tnode   *tablenode //交换时的过渡指针;
	)

	if head.next == nil {
		fmt.Printf("please use a legal list")
		//exit(1)
	}

	for m = head.frequence; m >= 0; m-- {
		a = m
		newnode = head
		for newnode.next != nil && newnode.next.next != nil && a > 0 {
			if newnode.next.frequence < newnode.next.next.frequence {
				tnode = newnode.next
				newnode.next = tnode.next
				tnode.next = newnode.next.next
				newnode.next.next = tnode
				//		newnode = tnode;
			}
			newnode = newnode.next
			a--
		}
	}
	return *head
	//delete newnode;
	//delete tnode;
}

func getfrequenceitem(support int, head *tablenode) tablenode {

	var newnode *tablenode

	if head.next == nil {
		fmt.Printf("this is a wrong list\n")
		//exit(1)
	}
	newnode = head
	for newnode.next != nil {
		if newnode.next.frequence >= support {
			newnode = newnode.next
		} else {

			//tnode := newnode.next
			newnode.next = nil
			break
		}
	}
	return *head
	//delete newnode;

}

func getftransaction(data []byte, n int, head *tablenode) []byte {
	var (
		a       byte
		m       int
		i       int
		newnode *tablenode
	)
	if head.next == nil {
		fmt.Printf("this is a wrong list\n")
		//exit(1);
	}
	newnode = head.next
	for newnode != nil {
		for i = m; i < n; i++ {
			if data[i] == newnode.itemname {
				a = data[m]
				data[m] = data[i]
				data[i] = a
				m++
				break
			}
		}
		//if(newnode->next !=null)
		newnode = newnode.next
		//else break;
	}
	for i = m; i < n; i++ {
		data[i] = ' '
	}
	return data
	//delete newnode;
}

func printfptree(root *fpnode) {
	var (
		newnode, tnode, printnode *fpnode
		//fpnode *tnode
		//fpnode *newnode
	)
	if root.child == nil {

		//exit(1);
	}
	newnode = root.child
	fmt.Printf("itemname is %v", newnode.itemname)
	if newnode.brother != nil {
		tnode = newnode.brother
		printnode = newnode.brother
	}

	for printnode != nil {
		fmt.Printf("itemname is %v", printnode.itemname)

		printnode = printnode.brother
	}
	fmt.Printf("\n")
	printfptree(newnode)
	printfptree(tnode)
	//delete printnode;
	//delete newnode;
}
func m() {
	data := []byte{'a','b','c','f','g',' '}
	data1 := []byte{'n','f','g',' '}
	data2 := []byte{'m','b','h','f','h',' '}
	//test := []byte{'b','h',' '}

	root := (*fpnode).funcfpnode("")
	//先通过事务集求出项目链表,然后根据support求出频繁项目链表,
	support := 2

	head := (*tablenode).functablenode("")
	creattablelist(data, 6, head)
	creattablelist(data1, 4, head)
	creattablelist(data2, 6, head)
	bubblesort(head)
	getfrequenceitem(support, head)
	getftransaction(data, 6, head)
	getftransaction(data2, 6, head)
	getftransaction(data1, 4, head)
	for head.next != nil {
		fmt.Printf("head.next.itemname is %v", head.next.itemname)
		fmt.Printf("head.next.frequence is %v", head.next.frequence)
		fmt.Printf("\n")
		head = head.next
	}
	i := 0
	for data[i] != ' ' {
		fmt.Printf("data[i] is %v", data[i])
		i++
	}
	fmt.Printf("\n")
	i = 0
	*/
/*for data2[i]!=' '
	{
	fmt.Printf("data[i] is %v",data2[i])
	i++
	}*//*

	fmt.Printf("\n")
	i = 0
	*/
/*for data1[i]!=nil
	{
		fmt.Printf("data[i] is %v",data1[i])
		i++
	}*//*

	fmt.Printf("\n")

*/
/*	insert(data[0], data, 6, root)
	insert(data2[0], data2, 6, root)
	insert(data1[0], data1, 4, root)
	insert(test[0], test, 3, root)*//*


	fmt.Printf("\n   root\n\n")
	fmt.Printf("%v\n", root.child.itemname)
	fmt.Printf("%v\n", root.child.count)
	fmt.Printf("   ")

	fmt.Printf("%v\n", root.child.brother.itemname)
	fmt.Printf("   ")
	fmt.Printf("%v\n", root.child.brother.count)
	fmt.Printf("  \n ")

	fmt.Printf("%v\n %v\n", root.child.child.itemname, root.child.child.count)

	fmt.Printf("%v\n %v\n", root.child.child.brother.itemname, root.child.child.brother.count)

	fmt.Printf("%v\n %v\n", root.child.brother.child.itemname, root.child.brother.child.count)

	fmt.Printf("%v\n %v\n", root.child.child.child.itemname, root.child.child.child.count)

	fmt.Printf("%v\n %v\n", root.child.child.child.brother.itemname, root.child.child.child.brother.count)

}
*/
