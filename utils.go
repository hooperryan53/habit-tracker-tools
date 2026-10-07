package main

// Chunk splits a slice into groups of size n.
func Chunk(items []int, n int) [][]int {
	var out [][]int
	for i := 0; i < len(items); i += n {
		end := i + n
		if end > len(items) {
			end = len(items)
		}
		out = append(out, items[i:end])
	}
	return out
}
