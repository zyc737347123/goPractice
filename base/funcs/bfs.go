package funcs

import "time"

// BFS is breadth First serach for each item in worklist
// f is called at most once for each item
func BFS(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for len(worklist) > 0 && time.Now().Before(deadline) {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
