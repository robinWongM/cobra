# cobra
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?logo=go)](https://pkg.go.dev/github.com/robinWongM/cobra?tab=doc)

A simple "replacement" for spf13/cobra, and it just works!

# Example

```golang
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
```

# Acknowledgements

This repository contains (lots of) code from [spf13/cobra](https://github.com/spf13/cobra), which is licensed under Apache License, Version 2.0.