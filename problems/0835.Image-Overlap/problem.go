package leetcode

func largestOverlap(img1 [][]int, img2 [][]int) int {
	maxOverlap := 0

	for yShift := 0; yShift < len(img1); yShift++ {
		for xShift := 0; xShift < len(img1); xShift++ {
			maxOverlap = max(maxOverlap, shiftAndCount(xShift, yShift, img1, img2))
			maxOverlap = max(maxOverlap, shiftAndCount(xShift, yShift, img2, img1))
		}
	}

	return maxOverlap
}

func shiftAndCount(xShift, yShift int, M, R [][]int) int {
	lShiftCount, rShiftCount, rRow := 0, 0, 0

	for mRow := yShift; mRow < len(M); mRow++ {
		rCol := 0
		for mCol := xShift; mCol < len(M); mCol++ {
			if M[mRow][mCol] == 1 && M[mRow][mCol] == R[rRow][rCol] {
				lShiftCount++
			}
			if M[mRow][rCol] == 1 && M[mRow][rCol] == R[rRow][mCol] {
				rShiftCount++
			}
			rCol++
		}
		rRow++
	}
	return max(lShiftCount, rShiftCount)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
