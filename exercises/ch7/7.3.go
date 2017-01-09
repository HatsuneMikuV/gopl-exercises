package main

import (
	"fmt"
	"strconv"
)

/*
   练习 7.3： 为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值序列。
*/

type tree struct {
	value       int
	left, right *tree
}

func (tre *tree) String() string {
	return "\ntree value is " + strconv.Itoa(tre.value) + ", \n← left value is " + strconv.Itoa(tre.left.value) +
		", \n→ right value is " + strconv.Itoa(tre.right.value) + "\n"

}
func main() {
	treLeft := &tree{1, nil, nil}
	treright := &tree{2, nil, nil}
	tre := &tree{0, treLeft, treright}
	fmt.Println(tre)
}
