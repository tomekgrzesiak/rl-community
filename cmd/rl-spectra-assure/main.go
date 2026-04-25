package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rl-community/rl-spectra-assure/internal/adapters/primary/cli"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Show help on first run or when no args
	if len(os.Args) < 2 {
		fmt.Println("\n--- Commands ---")
		fmt.Println("  help       - Show this help message")
		fmt.Println("  version    - Show version information")
		fmt.Println("  analyze    - Run spectral analysis")
		fmt.Println("  validate   - Validate input data")
		fmt.Println("  convert    - Convert between formats")
	}

	if err := cli.NewRootCmd().ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
