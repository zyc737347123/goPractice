package funcs

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for key := range items {
			if !seen[key] {
				seen[key] = true

				keys := make(map[string]bool)
				for _, i := range m[key] {
					keys[i] = true
				}
				visitAll(keys)

				order = append(order, key)
			}
		}
	}
	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)

	return order
}

// Topo is func
func Topo() {

	if hasCycle(prereqs) {
		return
	}

	order := topoSort(prereqs)
	for _, v := range order {
		fmt.Println(v)
	}
	fmt.Println()
	order = topoSort2(prereqs)
	for _, v := range order {
		fmt.Println(v)
	}
}

// Not Good algorithms
func hasCycle(org map[string][]string) bool {
	var stack []string
	degree := make(map[string]int)
	m := make(map[string][]string)

	for k, v := range org {
		m[k] = v
	}

	for key, value := range m {
		degree[key] = len(value)
		for _, v := range value {
			degree[v] += 0
		}
	}

	for k, v := range degree {
		if v == 0 {
			stack = append(stack, k)
			degree[k] = -1
		}
	}

	var top string
	for len(stack) > 0 {
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// remove top from graph
		for k, v := range m {
			for index, i := range v {
				if i == top {
					v[index] = v[len(v)-1]
					m[k] = v[:len(v)-1]
					degree[k]--
					if degree[k] == 0 {
						stack = append(stack, k)
						degree[k] = -1
					}
				}
			}
		}
	}

	for _, v := range degree {
		if v != -1 {
			fmt.Println("has cycle")
			return true
		}
	}

	return false
}
