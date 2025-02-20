package slice

import "fmt"

func PrintSlice(slice []int, label string) {
	fmt.Println(label)
	for _, val := range slice {
		fmt.Printf("\t%d\t", val)
	}
	fmt.Println()

}

func PrintSlice2D(slice [][]int, label string) {
	fmt.Println(label)
	for _, rows := range slice {
		for _, val := range rows {
			fmt.Printf("\t%d\t", val)
		}
		fmt.Println()
	}
	fmt.Println()

}
