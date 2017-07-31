package main

import "fmt"

func main() {
	fmt.Println("FINAL:", teste(1, 2, 3, 4, 5))
}

func teste(values ...int) int {
	var final int
outer:
	for i := range values {
		fmt.Println("RUN ONE:", i)
		if i == 4 {
			break
		}

		for _, j := range values {
			fmt.Println("RUN TWO:", j)
			if j == 2 {
				fmt.Println("VOLTAR AO INICIO")
				break outer // just break the for
			} else {
				goto outer // goto FLAG
			}
		}
		final = i
		break
	}
	return final
}
