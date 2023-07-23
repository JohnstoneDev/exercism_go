package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for _, checkOccupied := range cb[file] {
		if checkOccupied {
			count ++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

	// check if the rank is valid
	if rank < 1 || rank > 8 {
		return 0
	}

	for file := range cb {
		if cb[file][rank - 1] {
			count ++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count ++
	}
	// count all ranks, multiply horizontally
	count *= count
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for file := range cb {
		count += CountInFile(cb, file)
	}

	return count
}
