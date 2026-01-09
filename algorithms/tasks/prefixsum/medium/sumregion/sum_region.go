package sumregion

type NumMatrix struct {
	sums [][]int
}

/*
[[[
	[3,0,1,4,2], 10
	[5,6,3,2,1], 17
	[1,*2,0,1,5], 9
	[4,1,0,1,7], 13
	[1,0,3,*0,5], 9
]],

sums
[[[
	[3,		3,		4,		8,		10],
	[8,		14,		18,		24,		27],
	[9,		*17,	21,		28,		36],
	[13,	22,		26,		34,		49],
	[14,	23,		30,		*38,	58],
]],


[[[[-1]]],[0,0,0,0]]
Expected -1

[[[
	[-4,-5]
]],

[0,0,0,0], -4
[0,0,0,1], -9
[0,1,0,1]] -5

Output
[null,0,0,0]
Expected
[null,-4,-9,-5]

*/

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	cols := len(matrix[0])

	sums := make([][]int, rows)
	for i := 0; i < rows; i++ {
		sums[i] = make([]int, cols)
	}

	sums[0][0] = matrix[0][0]
	for i := 1; i < cols; i++ {
		sums[0][i] = matrix[0][i] + sums[0][i-1]
	}
	for i := 1; i < rows; i++ {
		sums[i][0] = matrix[i][0] + sums[i-1][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			sums[i][j] = matrix[i][j] + sums[i][j-1] + sums[i-1][j] - sums[i-1][j-1]
		}
	}

	return NumMatrix{
		sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	topSum := 0
	if row1 > 0 {
		topSum = this.sums[row1-1][col2]
	}
	leftSum := 0
	if col1 > 0 {
		leftSum = this.sums[row2][col1-1]
	}

	diagonalSumBefore := 0
	if row1 > 0 && col1 > 0 {
		diagonalSumBefore = this.sums[row1-1][col1-1]
	}

	return this.sums[row2][col2] - topSum - leftSum + diagonalSumBefore
}

/*
[[[
[-4,-5]
]],

sums:
 [-4, -9]

[0,0,0,0],
[0,0,0,1],
[0,1,0,1]]

Use Testcase
Output
[null,-4,-9,-9]
Expected
[null,-4,-9,-5]
*/

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
