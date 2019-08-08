package M_PrisonCellsAfterNDays_957

/*
There are 8 prison cells in a row, and each cell is either occupied or vacant.

Each day, whether the cell is occupied or vacant changes according to the following rules:

If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
Otherwise, it becomes vacant.
(Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)

We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)



Example 1:

Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation:
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

Example 2:

Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]

*/

func prisonAfterNDays(cells []int, N int) []int {
	n := cells2int(cells)

	N = (N-1)%14 + 1

	for i := 0; i < N; i++ {
		n = (^((n << 1) ^ (n >> 1))) & 0x7e
	}

	return int2cells(n)
}

func cells2int(cells []int) int {
	res := 0
	t := uint(7)
	for _, c := range cells {
		res += c << t
		t--
	}
	return res
}

func int2cells(n int) []int {
	cells := make([]int, 8)
	t := uint(7)
	for i := range cells {
		cells[i] = (n >> t) & 1
		t--
	}
	return cells
}
