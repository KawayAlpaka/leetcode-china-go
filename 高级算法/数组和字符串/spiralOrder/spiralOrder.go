package spiralorder

// Point ..
type point struct {
	i, j int
}

func spiralOrder(matrix [][]int) []int {
	rl := len(matrix)
	if rl == 0 {
		return []int{}
	}
	cl := len(matrix[0])
	if cl == 0 {
		return []int{}
	}
	var fa [][]int
	for i := 0; i < rl; i++ {
		fa = append(fa, make([]int, cl))
	}
	// fmt.Printf("%d \n", fa)
	directions := []point{
		point{i: 0, j: 1},
		point{i: 1, j: 0},
		point{i: 0, j: -1},
		point{i: -1, j: 0},
	}
	direIndex := 0

	isValid := func(i, j int) bool {
		if i < 0 || i >= rl || j < 0 || j >= cl || fa[i][j] == 1 {
			return false
		}
		return true
	}

	getNextPoint := func(i int, j int) point {
		direction := directions[direIndex]
		_i := i + direction.i
		_j := j + direction.j
		if !isValid(_i, _j) {
			direIndex++
			if direIndex > 3 {
				direIndex = 0
			}
			direction = directions[direIndex]
			_i = i + direction.i
			_j = j + direction.j
		} else {
			return point{_i, _j}
		}
		return point{_i, _j}
	}

	var result []int

	for i, j := 0, 0; true; {
		// fmt.Printf("%v \n", result)
		// fmt.Printf("i:%d,j:%d \n", i, j)
		if !isValid(i, j) {
			break
		}
		result = append(result, matrix[i][j])
		fa[i][j] = 1
		p := getNextPoint(i, j)
		i = p.i
		j = p.j
	}
	// fmt.Printf("%v \n", fa)
	return result
}
