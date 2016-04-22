package main
import (
	"fmt"
	"net/http/fcgi"
)


func main() {

	var fcgi  [][]string
	fcgi={"{"你好","不好","是吗"}","{"你好","不好","是吗"}","{"你好","不好","是吗"}"}
	for no,value :=range fcgi{

		fmt.Println(no,value)
	}

}
