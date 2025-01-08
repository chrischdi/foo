package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	if bi, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("Main.Version: %s\n", bi.Main.Version)
		fmt.Printf("Main.Sum: %s\n", bi.Main.Sum)
		fmt.Printf("Main.Path: %s\n", bi.Main.Path)
		for _, v := range bi.Settings {
			fmt.Printf("%s: %s\n", v.Key, v.Value)
		}
	}
}
