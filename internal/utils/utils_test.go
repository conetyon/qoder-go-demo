package utils

import "testing"

func TestPadLeft(t *testing.T) {
	tests := []struct {
		s      string
		length int
		pad    byte
		want   string
	}{
		{"go", 10, '.', "........go"},
		{"hello", 5, ' ', "hello"},
		{"long", 3, 'x', "long"},
		{"", 4, '-', "----"},
	}
	for _, tt := range tests {
		got := PadLeft(tt.s, tt.length, tt.pad)
		if got != tt.want {
			t.Errorf("PadLeft(%q, %d, %q) = %q, want %q", tt.s, tt.length, tt.pad, got, tt.want)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		s      string
		length int
		pad    byte
		want   string
	}{
		{"go", 10, '.', "go........"},
		{"hello", 5, ' ', "hello"},
		{"", 3, '*', "***"},
	}
	for _, tt := range tests {
		got := PadRight(tt.s, tt.length, tt.pad)
		if got != tt.want {
			t.Errorf("PadRight(%q, %d, %q) = %q, want %q", tt.s, tt.length, tt.pad, got, tt.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		s      string
		maxLen int
		want   string
	}{
		{"Hello Qoder World", 12, "Hello Qod..."},
		{"Short", 10, "Short"},
		{"abc", 3, "abc"},
		{"abcdef", 4, "a..."},
		{"abcdef", 2, "ab"},
	}
	for _, tt := range tests {
		got := Truncate(tt.s, tt.maxLen)
		if got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.maxLen, got, tt.want)
		}
	}
}

func TestRandomHex(t *testing.T) {
	hex, err := RandomHex(8)
	if err != nil {
		t.Fatalf("RandomHex error: %v", err)
	}
	if len(hex) != 16 {
		t.Errorf("RandomHex(8) length = %d, want 16", len(hex))
	}
	// Two calls should produce different results
	hex2, _ := RandomHex(8)
	if hex == hex2 {
		t.Error("Two RandomHex calls should produce different values")
	}
}

func TestRandomInt(t *testing.T) {
	for i := 0; i < 20; i++ {
		n, err := RandomInt(100)
		if err != nil {
			t.Fatalf("RandomInt error: %v", err)
		}
		if n < 0 || n >= 100 {
			t.Errorf("RandomInt(100) = %d, out of range [0, 100)", n)
		}
	}
}

func TestTable(t *testing.T) {
	headers := []string{"Name", "Value"}
	rows := [][]string{
		{"Go", "1.22"},
		{"Qoder", "latest"},
	}
	output := Table(headers, rows)
	if len(output) == 0 {
		t.Error("Table output should not be empty")
	}
	// Check it contains the data
	for _, h := range headers {
		if !contains(output, h) {
			t.Errorf("Table output missing header %q", h)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
