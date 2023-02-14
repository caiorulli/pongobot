// Just practicing go by implementing nqueens.
// Completely unrelated to the rest of the project.
package main

func Nqueens(n int) int {
	board := make([]int, n)
	valid_occurrences := 0
	for i := 0; i < n; i++ {
		board[0] = i
		valid_occurrences += nqueens_recur(n, board, 1);
	}
	return valid_occurrences
}

func nqueens_recur(n int, board []int, level int) int {
	if !is_valid(n, board, level) {
		return 0
	} else if level == n {
		return 1
	} else {
		valid_occurrences := 0
		for i := 0; i < n; i++ {
			board[level] = i
			valid_occurrences += nqueens_recur(n, board, level + 1)
		}
		return valid_occurrences
	}
}

func is_valid(n int, board []int, level int) bool {
	current_pos := level - 1
	ref := board[current_pos]
	for i := 0; i < level - 1; i++ {
		local := board[i]
		level_diff := current_pos - i
		if local == ref || local + level_diff == ref || local - level_diff == ref {
			return false
		}
	}
	return true
}
