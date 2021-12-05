package libs

import "fmt"

func PrintStringArray(arr []string) {
	for _, line := range arr {
		fmt.Println(line)
	}
}
