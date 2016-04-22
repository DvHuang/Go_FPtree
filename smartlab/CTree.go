
package fpTree

import "fmt"


type Fpnode struct {
	itemname string
	count    int
	parent   *Fpnode
	child    *Fpnode
	brother  *Fpnode
}

func (fpnode *Fpnode) makefpnode(name string) *Fpnode {
	fpnode.itemname = name
	fpnode.count = 0
	fpnode.parent = nil
	fpnode.child = nil
	fpnode.brother = nil
	return *fpnode
}


func Insertfpnode (item string,itemsets []string,root *Fpnode) *Fpnode {

	True := 0
	midnode := root.child

	if midnode == nil {
		True = 2
	}

	for midnode != nil {

		if (*midnode).itemname == item {
			True = 1
			break
		} else {
			if midnode.brother != nil {
				midnode = midnode.brother
			} else {
				fmt.Printf("运行到这里说明这一行节点及其兄弟节点都没有我们需要的，即true=0")
				break
			}
		}
	}

	if True == 1 {
		(*midnode).count++
	}

	if True == 2 {
		newnode := (*Fpnode).makefpnode(item)
		(*newnode).count++
		newnode.parent = root
		root.child = newnode
		midnode = newnode
	}

	if True == 0 {
		newnode := (*Fpnode).makefpnode(item)
		newnode.count++
		newnode.parent = root
		midnode.brother = newnode
		midnode = newnode
	}

	if itemsets[1] !=0 {
		a := itemsets[1]
		itemsets=itemsets[1:]
		Insertfpnode(string(a),itemsets, midnode)
	}

	return *root

}


type Tablenode struct {
	frequence int
	next      *Fpnode
}

func (tablenode *Tablenode) functablenode(name []byte) *Tablenode {
	tablenode.frequence = 0
	tablenode.next = nil
	return  tablenode
}

func Creattablelist(item []byte, i int, head *Tablenode) map[string]Tablenode {

	j := 0
	node := head

	if node.next == nil {
		for j,_ =range item {


			newnode := (*Tablenode).functablenode(item[j])
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
					newnode := (*Tablenode).functablenode(item[j])
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



type FrequentItemsets  struct {


}

func Fp_Growth(support int,root *Fpnode,tablelist *Tablenode) FrequentItemsets {





	return frequentItemsets

}

//ascends from leaf node to root
func ascendTree (leafNode *Fpnode, prefixPath []Fpnode)  {

	if leafNode.parent != nil {
		prefixPath.append(leafNode.itemname)
	}
	ascendTree(leafNode.parent, prefixPath)
}

//treeNode comes from header table
func  findPrefixPath (basePat Tablenode, treeNode *Fpnode) {

	condPats = {}
	for range  treeNode != nil {
		var prefixPath = []Fpnode
		ascendTree(treeNode, prefixPath)
		if len(prefixPath) > 1 {
			condPats[frozenset(prefixPath[1:])] = treeNode.count
		}
		treeNode = treeNode.brother
	}
	return condPats

}

