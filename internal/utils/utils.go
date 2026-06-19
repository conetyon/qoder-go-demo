// Package utils provides general-purpose utility functions.
package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// PadLeft pads a string on the left to the given length with the specified character.
func PadLeft(s string, length int, pad byte) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(pad), length-len(s)) + s
}

// PadRight pads a string on the right to the given length with the specified character.
func PadRight(s string, length int, pad byte) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(string(pad), length-len(s))
}

// Truncate truncates a string to maxLen and appends "..." if truncated.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 4 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// RandomHex generates a random hex string of n bytes (2n characters).
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

// RandomInt returns a random integer in [0, max).
func RandomInt(max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// Table formats data as a simple aligned text table.
func Table(headers []string, rows [][]string) string {
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	var sb strings.Builder
	// Header
	for i, h := range headers {
		sb.WriteString(PadRight(h, widths[i], ' '))
		if i < len(headers)-1 {
			sb.WriteString(" │ ")
		}
	}
	sb.WriteString("\n")
	// Separator
	for i, w := range widths {
		sb.WriteString(strings.Repeat("─", w))
		if i < len(widths)-1 {
			sb.WriteString("─┼─")
		}
	}
	sb.WriteString("\n")
	// Rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				sb.WriteString(PadRight(cell, widths[i], ' '))
				if i < len(row)-1 {
					sb.WriteString(" │ ")
				}
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
