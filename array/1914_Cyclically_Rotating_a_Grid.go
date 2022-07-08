package array

func rotateGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	result := make([][]int, 0)

	for i := range grid {
		row := make([]int, 0)

		for j := range grid[i] {
			d := findD(i, j, n, m)
			side := findSide(i, j, n, m, d)
			x, y := findValue(i, j, n, m, d, side, k)
			row = append(row, grid[x][y])
		}

		result = append(result, row)
	}

	return result
}

func findValue(i, j, n, m, d, side, k int) (int, int) {
	l := 2*(n-2*d) + 2*(m-2*d) - 4
	k = k % l

	var diff, step int
	for k > 0 {

		switch side {
		case 0:
			diff = i - d
			step = min(diff, k)
			i -= step
		case 1:
			diff = m - d - j - 1
			step = min(diff, k)
			j += step
		case 2:
			diff = n - d - i - 1
			step = min(diff, k)
			i += step
		default:
			diff = j - d
			step = min(diff, k)
			j -= step
		}

		k -= step
		side = (side + 1) % 4
	}

	return i, j
}

func findSide(i, j, n, m, d int) int {
	if i > d && i < n-d && j == d {
		return 0
	}

	if j >= d && j < m-d-1 && i == d {
		return 1
	}

	if i >= d && i < n-d-1 && j == m-d-1 {
		return 2
	}

	return 3
}

func findD(i, j, n, m int) int {
	return min(min(i, j), min(n-i-1, m-j-1))
}

/*func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}*/
