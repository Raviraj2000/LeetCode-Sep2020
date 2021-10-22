package main

import ("fmt")


func sumRootToLeaf(root *TreeNode) int64 {
  var path []int64
  getPath(root, path)
  return root.Right.Val
}

func getPath(node *TreeNode, path []int64) {
  if (node == nil) {
    return
  }
  path = append(path, node.Val)

  if (node.Left == nil && node.Right == nil){
    fmt.Println(path)
  }else {
    getPath(node.Left, path)
    getPath(node.Right, path)
  }
}


type TreeNode struct {
    Val int64
    Left  *TreeNode
    Right *TreeNode

}

func (n *TreeNode) Insert(value int64) {
    if n == nil {
        return
    } else if value < n.Val {
        if n.Left == nil {
            n.Left = &TreeNode{Val: value, Left: nil, Right: nil}
        } else {
            n.Left.Insert(value)
        }
    } else {
        if n.Right == nil {
            n.Right = &TreeNode{Val: value, Left: nil, Right: nil}
        } else {
            n.Right.Insert(value)
        }
    }
}


func main() {
  root := TreeNode{Val : 1}
  root.Insert(0)
  root.Insert(1)
  root.Insert(0)
  root.Insert(1)
  root.Insert(0)
  root.Insert(1)
  fmt.Println(sumRootToLeaf(&root))

}
