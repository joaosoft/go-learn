package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func main(){
	LinebyLineScan()
}
func LinebyLineScan() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}