/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 15:08
 *  @Description：矩阵旋转问题解决(这个是之前在2048游戏里面)
 */

package problem

import "fmt"

const (
	SIZE = 4
)

func MatrixRotationTest() {
	//矩阵旋转.....
	e := [SIZE][SIZE]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	MatrixRotation(e)
	MatrixRotationAnti(e)
	MatrixRotationHalfCircle(e)
	MatrixRotationAntiHalfCircle(e)
}

//矩阵旋转-逆时针旋转90度
func MatrixRotationAnti(matrix [SIZE][SIZE]int) [SIZE][SIZE]int {
	var temp [SIZE][SIZE]int
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			temp[SIZE-j-1][i] = matrix[i][j]
		}
	}
	fmt.Println(temp)
	return temp
}

//顺时针旋转90度
func MatrixRotation(matrix [SIZE][SIZE]int) [SIZE][SIZE]int {
	var temp [SIZE][SIZE]int
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			temp[j][SIZE-1-i] = matrix[i][j]
		}
	}
	fmt.Println(temp)
	return temp
}

//矩阵逆时针旋转180度 --顺逆时针结果一致
func MatrixRotationAntiHalfCircle(matrix [SIZE][SIZE]int) [SIZE][SIZE]int {
	var temp [SIZE][SIZE]int
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			temp[SIZE-1-i][SIZE-1-j] = matrix[i][j]
		}
	}
	fmt.Println(temp)
	return temp
}

//矩阵顺时针旋转180度
func MatrixRotationHalfCircle(matrix [SIZE][SIZE]int) [SIZE][SIZE]int {
	var temp [SIZE][SIZE]int
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			temp[SIZE-1-i][SIZE-1-j] = matrix[i][j]
		}
	}
	fmt.Println(temp)
	return temp
}
