// testes

package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_funcs()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func funcs_compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func _funcs() {
	fmt.Println("_funcs()")
	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(1, 1))

	fmt.Println(funcs_compute(hypot))
	fmt.Println(funcs_compute(math.Pow))
}
