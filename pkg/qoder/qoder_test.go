package qoder

import (
	"strings"
	"testing"
)

// --- Banner Tests ---

func TestBanner_ContainsQoder(t *testing.T) {
	banner := Banner()
	if !strings.Contains(banner, "Qoder") && !strings.Contains(banner, "qoder") && !strings.Contains(banner, "Q O D E R") && !strings.Contains(banner, "_") {
		t.Error("Banner should contain Qoder branding")
	}
}

func TestVersion_ContainsGo(t *testing.T) {
	v := Version()
	if !strings.Contains(v, "Go") {
		t.Errorf("Version should mention Go, got: %s", v)
	}
}

// --- Greeter Tests ---

func TestNewGreeter_Casual(t *testing.T) {
	g := NewGreeter("casual")
	greeting := g.Greet("Alice")
	if !strings.Contains(greeting, "Alice") {
		t.Errorf("Greeting should contain the name, got: %s", greeting)
	}
	if !strings.Contains(greeting, "Qoder") {
		t.Errorf("Greeting should mention Qoder, got: %s", greeting)
	}
}

func TestNewGreeter_Formal(t *testing.T) {
	g := NewGreeter("formal")
	greeting := g.Greet("Bob")
	if !strings.Contains(greeting, "Bob") {
		t.Errorf("Greeting should contain the name, got: %s", greeting)
	}
	if !strings.Contains(greeting, "Good day") {
		t.Errorf("Formal greeting should say 'Good day', got: %s", greeting)
	}
}

func TestNewGreeter_Default(t *testing.T) {
	g := NewGreeter("unknown_style")
	greeting := g.Greet("Eve")
	if !strings.Contains(greeting, "Eve") {
		t.Errorf("Default greeting should still work, got: %s", greeting)
	}
}

// --- Generics Tests ---

func TestMap_IntToString(t *testing.T) {
	nums := []int{1, 2, 3}
	strs := Map(nums, func(n int) string {
		return strings.Repeat("*", n)
	})
	expected := []string{"*", "**", "***"}
	for i, s := range strs {
		if s != expected[i] {
			t.Errorf("Map[%d] = %q, want %q", i, s, expected[i])
		}
	}
}

func TestMap_StringToInt(t *testing.T) {
	words := []string{"go", "qoder", "demo"}
	lengths := Map(words, func(s string) int { return len(s) })
	expected := []int{2, 5, 4}
	for i, l := range lengths {
		if l != expected[i] {
			t.Errorf("Map[%d] = %d, want %d", i, l, expected[i])
		}
	}
}

func TestFilter_Ints(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	if len(evens) != 3 {
		t.Fatalf("Filter should return 3 evens, got %d", len(evens))
	}
	for _, n := range evens {
		if n%2 != 0 {
			t.Errorf("Filter returned odd number: %d", n)
		}
	}
}

func TestFilter_Empty(t *testing.T) {
	nums := []int{1, 3, 5}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	if len(evens) != 0 {
		t.Errorf("Filter should return empty slice, got %v", evens)
	}
}

// --- Concurrency Tests ---

func TestConcurrentProcess_Basic(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	results := ConcurrentProcess(items, 3, func(n int) int {
		return n * 10
	})
	expected := []int{10, 20, 30, 40, 50}
	for i, r := range results {
		if r != expected[i] {
			t.Errorf("ConcurrentProcess[%d] = %d, want %d", i, r, expected[i])
		}
	}
}

func TestConcurrentProcess_Strings(t *testing.T) {
	items := []string{"hello", "world"}
	results := ConcurrentProcess(items, 2, strings.ToUpper)
	expected := []string{"HELLO", "WORLD"}
	for i, r := range results {
		if r != expected[i] {
			t.Errorf("ConcurrentProcess[%d] = %q, want %q", i, r, expected[i])
		}
	}
}

func TestConcurrentProcess_SingleWorker(t *testing.T) {
	items := []int{10, 20}
	results := ConcurrentProcess(items, 1, func(n int) int { return n + 1 })
	if results[0] != 11 || results[1] != 21 {
		t.Errorf("Single worker failed: got %v", results)
	}
}

// --- Features Tests ---

func TestFeatures_NotEmpty(t *testing.T) {
	feats := Features()
	if len(feats) == 0 {
		t.Error("Features() should return at least one feature")
	}
	for _, f := range feats {
		if f.Name == "" {
			t.Error("Feature name should not be empty")
		}
		if f.Description == "" {
			t.Error("Feature description should not be empty")
		}
	}
}
