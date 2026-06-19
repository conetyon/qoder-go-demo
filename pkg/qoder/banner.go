// Package qoder provides Qoder-themed utilities and services.
// Built with Go 1.22, showcasing generics, interfaces, and concurrency.
package qoder

import "fmt"

// Banner returns a Qoder ASCII art banner.
func Banner() string {
	return `
   ___          _            
  / _ \  _   _ | |__    ___  _ __  
 | | | || | | || '_ \  / _ \| '__| 
 | |_| || |_| || |_) ||  __/| |    
  \__\_\ \__,_||_.__/  \___||_|    
                                    
  ⚡ Powered by AI Coding Assistant ⚡
`
}

// Version returns the current version of the demo.
func Version() string {
	return "v1.0.0 — Built with Qoder + Go 1.22"
}

// Feature represents a Qoder capability.
type Feature struct {
	Name        string
	Description string
}

// Features returns a list of Qoder features.
func Features() []Feature {
	return []Feature{
		{Name: "Code Generation", Description: "AI-powered code generation across languages"},
		{Name: "Smart Completion", Description: "Context-aware code suggestions"},
		{Name: "Code Review", Description: "Automated code quality analysis"},
		{Name: "Refactoring", Description: "Intelligent code restructuring"},
		{Name: "Debugging", Description: "AI-assisted bug detection and fixing"},
		{Name: "Multi-file Editing", Description: "Seamless cross-file modifications"},
	}
}

// PrintFeatures prints all features to stdout.
func PrintFeatures() {
	fmt.Println("🚀 Qoder Features:")
	fmt.Println("─────────────────────────────────────")
	for i, f := range Features() {
		fmt.Printf("  %d. %-20s %s\n", i+1, f.Name, f.Description)
	}
	fmt.Println("─────────────────────────────────────")
}
