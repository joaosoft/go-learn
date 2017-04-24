// testes

package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_loops()
	
	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions
func _loops() {
	fmt.Println("_loops()")
	// while
	for true {
		fmt.Println("for true")
		break
	}

	// infinite for
	for {
		fmt.Println("infinite for")
		break
	}

	//O range do laço for itera sobre uma slice ou map.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

}
