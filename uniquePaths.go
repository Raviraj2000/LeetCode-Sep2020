package main

import ("fmt")

func uniquePathsIII(grid [][]int) int {
  rows := len(grid)
  if rows == 0 {
    return 0
  }
  cols := len(grid[0])
  if cols == 0{
    return 0
  }
  var sr,sc,er,ec int
  var ans [][]int
  var todo int

  for i := 0; i < rows; i++{
    for j := 0; j < cols; j++{
      if grid[i][j] == 1{
        sr = i
        sc = j
        ans = append(ans, start)
      }
      if grid[i][j] == 2{
        er = i
        ec = j
        end = []int{i,j}
      }
      if grid[i][j] != -1{
        todo++
      }
    }
  }

  //var visited [rows][cols]bool
  var ans int
  findPath(grid, ans, sr, sc, er, ec, todo)
  return ans
  fmt.Println(start,end,todo)
  return 1
}

func findPath(grid [][]int, ans int, sr int, sc int, er int, ec int, todo int) {
  todo --
  if (todo < 0) {
    return
  }
  if (sr == er && sc == ec) {
    if (todo == 0){
      ans ++
    }
    return
  }
   R := len(grid)
   C := len(grid[0])

   grid[sr][sc] = 3

   for k := 0; k < len(4); k++ {
     
   }



}



func main () {
  grid := [][]int{{1,0,0,0},{0,0,0,0},{0,0,2,-1}}

  fmt.Println(uniquePathsIII(grid))

}
