package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/Functions"
)

func main() {
	banner := "standard"
	Arguments := os.Args[1:]
	// checking if the N of args are between 1 (only the name) and 3 ARGs
	if len(Arguments) < 1 || len(Arguments) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		return
	}
	if len(Arguments) == 1 {
		sentence := Arguments[0]
		if sentence == "" {
			return
		}
		result := asciiart.PrintChar(asciiart.Split(sentence), asciiart.GetTheSymboles(banner))
		fmt.Println(result)
	}
	if len(Arguments) == 2 {
		if strings.HasPrefix(Arguments[0], "--output=") {
			if len(Arguments[0]) > 9 {
				fileName := Arguments[0][9:]
				sentence := Arguments[1]
				if sentence == "" {
					os.WriteFile(fileName, []byte{}, 0o777)
					return
				}
				symboles := asciiart.GetTheSymboles(banner)
				result := asciiart.PrintChar(asciiart.Split(sentence), symboles)
				os.WriteFile(fileName, []byte(result), 0o777)
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
				return
			}
		} else {
			if Arguments[0] == "" {
				return
			} else {
				sentence := Arguments[0]
				if Arguments[1] != "standard" && Arguments[1] != "shadow" && Arguments[1] != "thinkertoy" {
					fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
					return
				} else {
					banner = Arguments[1]
					symboles := asciiart.GetTheSymboles(banner)
					result := asciiart.PrintChar(asciiart.Split(sentence), symboles)
					fmt.Println(result)
				}
			}
		}
	}
	if len(Arguments) == 3 {
		if strings.HasPrefix(Arguments[0], "--output=") {
			if len(Arguments[0]) > 9 {
				fileName := Arguments[0][9:]
				sentence := Arguments[1]
				if sentence == "" {
					os.WriteFile(fileName, []byte{}, 0o777)
					return
				}
				if Arguments[2] != "standard" && Arguments[2] != "shadow" && Arguments[2] != "thinkertoy" {
					fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
					return
				} else {
					banner = Arguments[2]
					symboles := asciiart.GetTheSymboles(banner)
					result := asciiart.PrintChar(asciiart.Split(sentence), symboles)
					os.WriteFile(fileName, []byte(result), 0o777)
					return
				}
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
			return
		}
	}
}
