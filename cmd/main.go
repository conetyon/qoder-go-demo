// Qoder Go Demo — Main entry point
//
// A demonstration project showcasing Go features with Qoder theming.
// Includes generics, interfaces, concurrency, and CLI interaction.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/yty-dev/qoder-go-demo/internal/utils"
	"github.com/yty-dev/qoder-go-demo/pkg/qoder"
)

func main() {
	name := flag.String("name", "Developer", "Your name for the greeting")
	style := flag.String("style", "casual", "Greeting style: casual or formal")
	workers := flag.Int("workers", 4, "Number of concurrent workers for the demo")
	showTable := flag.Bool("table", false, "Show feature comparison table")
	flag.Parse()

	// 1. Banner
	fmt.Println(qoder.Banner())
	fmt.Println("  " + qoder.Version())
	fmt.Printf("  Go %s | %s/%s\n\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	// 2. Greeting (interface demo)
	greeter := qoder.NewGreeter(*style)
	fmt.Println("💬 " + greeter.Greet(*name))
	fmt.Println()

	// 3. Generics demo: Map & Filter
	fmt.Println("📦 Generics Demo — Map & Filter")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("   Input:   %v\n", numbers)

	doubled := qoder.Map(numbers, func(n int) int { return n * 2 })
	fmt.Printf("   Doubled: %v\n", doubled)

	evens := qoder.Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("   Evens:   %v\n", evens)

	names := []string{"qoder", "golang", "cli", "demo"}
	upper := qoder.Map(names, strings.ToUpper)
	fmt.Printf("   Upper:   %v\n", upper)
	fmt.Println()

	// 4. Concurrency demo
	fmt.Printf("⚡ Concurrent Processing Demo (%d workers)\n", *workers)
	tasks := []string{
		"analyze code",
		"generate docs",
		"run tests",
		"format code",
		"build project",
		"deploy app",
	}

	results := qoder.ConcurrentProcess(tasks, *workers, func(task string) string {
		return fmt.Sprintf("✅ %s completed", task)
	})
	for _, r := range results {
		fmt.Printf("   %s\n", r)
	}
	fmt.Println()

	// 5. Random session ID
	hex, _ := utils.RandomHex(8)
	fmt.Printf("🔑 Session ID: %s\n\n", hex)

	// 6. Optional table
	if *showTable {
		fmt.Println("📋 Qoder Feature Matrix:")
		fmt.Println()
		headers := []string{"Feature", "Description"}
		rows := make([][]string, 0)
		for _, f := range qoder.Features() {
			rows = append(rows, []string{f.Name, f.Description})
		}
		fmt.Print(utils.Table(headers, rows))
		fmt.Println()
	}

	// 7. Utility demos
	fmt.Println("🛠️  Utility Functions:")
	fmt.Printf("   PadLeft(\"go\", 10, '.'):  %q\n", utils.PadLeft("go", 10, '.'))
	fmt.Printf("   Truncate(\"Hello Qoder World\", 12): %q\n", utils.Truncate("Hello Qoder World", 12))
	fmt.Println()

	fmt.Println("🙏 Thanks for trying the Qoder Go Demo!")
	fmt.Println("   https://github.com/yty-dev/qoder-go-demo")
	os.Exit(0)
}
