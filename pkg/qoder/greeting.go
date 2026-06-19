package qoder

import (
	"fmt"
	"strings"
	"sync"
)

// Greeter defines the interface for generating greetings.
type Greeter interface {
	Greet(name string) string
}

// FormalGreeter generates formal greetings.
type FormalGreeter struct{}

func (g *FormalGreeter) Greet(name string) string {
	return fmt.Sprintf("Good day, %s. Welcome to the Qoder ecosystem.", name)
}

// CasualGreeter generates casual greetings.
type CasualGreeter struct{}

func (g *CasualGreeter) Greet(name string) string {
	return fmt.Sprintf("Hey %s! Ready to code with Qoder? 🚀", name)
}

// NewGreeter creates a greeter by style name.
func NewGreeter(style string) Greeter {
	switch strings.ToLower(style) {
	case "formal":
		return &FormalGreeter{}
	default:
		return &CasualGreeter{}
	}
}

// --- Generics Showcase ---

// Map applies a function to each element of a slice and returns a new slice.
// Demonstrates Go 1.18+ generics.
func Map[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// Filter returns elements that satisfy the predicate.
func Filter[T any](items []T, fn func(T) bool) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// ConcurrentProcess demonstrates goroutine-based concurrent processing.
// It processes items in parallel using worker goroutines and collects results.
func ConcurrentProcess[T any, U any](items []T, workerCount int, fn func(T) U) []U {
	if workerCount <= 0 {
		workerCount = 1
	}

	type indexed struct {
		index  int
		result U
	}

	jobs := make(chan int, len(items))
	results := make(chan indexed, len(items))
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				results <- indexed{index: idx, result: fn(items[idx])}
			}
		}()
	}

	// Send jobs
	for i := range items {
		jobs <- i
	}
	close(jobs)

	// Wait and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results in order
	output := make([]U, len(items))
	for r := range results {
		output[r.index] = r.result
	}
	return output
}
