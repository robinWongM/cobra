package main

import (
	"fmt"

	"github.com/robinWongM/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "root", Run: func(_ *cobra.Command, _ []string) {}}

	helloCmd := &cobra.Command{
		Use: "hello",
		Run: func(_ *cobra.Command, args []string) {
			name := "World"
			if len(args) >= 1 {
				name = args[0]
			}
			fmt.Printf("Hello, %v!\n", name)
		},
	}

	rootCmd.AddCommand(helloCmd)

	rootCmd.Execute()
}
