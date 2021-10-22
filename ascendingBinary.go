package main

import ("fmt"
        "errors"
        "sort")

var NodeList []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) ([]int){
  sortedArray1 := Sort(root1)
  NodeList = nil
  sortedArray2 := Sort(root2)
  NodeList = append(sortedArray1, sortedArray2...)
  fmt.Println(NodeList)
  sort.Ints(NodeList)
  return NodeList
}

func Sort(n *TreeNode) []int{

  if (n == nil){
    return nil
  } else {
		Sort(n.Left)
    NodeList = append(NodeList, n.Val)
		Sort(n.Right)
	}

  return NodeList
  }




type TreeNode struct {
    Val int
    Left  *TreeNode
    Right *TreeNode

}

func (t *TreeNode) Insert(value int) error {
  if t == nil {
    return errors.New("Tree is nil")
  }
  if t.Val == value {
    return errors.New("Node already exists")
  }

  if t.Val > value {
    if t.Left == nil {
      t.Left = &TreeNode{Val:value}
      return nil
    }
    return t.Left.Insert(value)
  }

  if t.Val < value {
    if t.Right == nil {
      t.Right = &TreeNode{Val:value}
      return nil
    }
    return t.Right.Insert(value)
  }
  return nil
}



func main() {
  root1 := TreeNode{Val : 0}
  root1.Insert(-10)
  root1.Insert(10)
  root2 := TreeNode{Val : 5}
  root2.Insert(1)
  root2.Insert(7)
  root2.Insert(0)
  root2.Insert(2)
  fmt.Println(getAllElements(&root1, &root2))

}
