package main

import "fmt"

func main() {
	//workDir, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("work dir:", workDir)
	//
	//executable, err := os.Executable()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//_, binName := filepath.Split(executable)
	//fmt.Println(binName)
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param grid int整型二维数组
 * @return int整型
 */
func ncov_defect(grid [][]int) int {
	// write code here
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
			}
		}
	}
	return res
}

func spiralOrder(matrix [][]int) []int {
	left, right := 0, len(matrix[0])-1
	up, down := 0, len(matrix)-1
	length := (right + 1) * (down + 1)
	var res []int
	if len(matrix) == 0 {
		return res
	}
	for {
		for i := left; i <= right && len(res) < length; i++ {
			res = append(res, matrix[up][i])
		}
		for i := up + 1; i <= down && len(res) < length; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right - 1; i >= left && len(res) < length; i-- {
			res = append(res, matrix[down][i])
		}
		for i := down - 1; i > up && len(res) < length; i-- {
			res = append(res, matrix[i][left])
		}
		if len(res) == length {
			break
		}
		left++
		right--
		up--
		down++
	}
	return res
}
