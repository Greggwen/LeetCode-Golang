package main

import "fmt"

// TODO:
func rotateMartix(matrix [][]int) {
	len := len(matrix)
	for i := 0; i < len/2; i++ {
		for j := 0; j < len/2; j++ {
			fmt.Println(i, j, len-i-1, j)
			matrix[i][j], matrix[len-i-1][j] = matrix[len-j-1][j], matrix[i][j]
			//00          30
			//01          20
			//02          10
			//03          00

			//00  03

		}

	}

	fmt.Printf("%v\n", matrix)
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	//ret := [][]int{
	//	{7, 4, 1},
	//	{8, 5, 2},
	//	{9, 6, 3},
	//}
	//fmt.Println(matrix)
	//fmt.Println(ret)

	rotateMartix(matrix)
}

//给定一个 n × n 的二维矩阵表示一个图像。
//将图像顺时针旋转 90 度。
//说明：
//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//示例 1:
//给定 matrix =
//[
//[1,2,3],
//[4,5,6],
//[7,8,9]
//],
//原地旋转输入矩阵，使其变为:
//[
//[7,4,1],
//[8,5,2],
//[9,6,3]
//]
//示例 2:
//给定 matrix =
//[
//[ 5, 1, 9,11],
//[ 2, 4, 8,10],
//[13, 3, 6, 7],
//[15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//[15,13, 2, 5],
//[14, 3, 4, 1],
//[12, 6, 8, 9],
//[16, 7,10,11]
//]
