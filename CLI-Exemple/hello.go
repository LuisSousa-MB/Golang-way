package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Hello() *cobra.Command {
	return &amp
	cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", args[0])
		},
	}
}
