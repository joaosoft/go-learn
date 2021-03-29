/////////////////////////////////
// Writing to Files using a Buffer in Memory
// Go Playground: https://play.golang.org/p/7U3g_B33aui
/////////////////////////////////

package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// Opening the file for writing
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// Creating a buffered writer from the file variable using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(file)

	// declaring a byte slice
	bs := []byte{97, 98, 99}

	// writing the byte slice to the buffer in memory
	bytesWritten, err := bufferedWriter.Write(bs)

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)
	// => 2019/10/21 16:30:59 Bytes written to buffer (not file): 3

	// checking the available buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
	// => 2019/10/21 16:30:59 Bytes available in buffer: 4093

	// writing a string (not a byte slice) to the buffer in memory
	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// checking how much data is stored in buffer, just  waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// -> 24 (3 bytes in the byte slice + 21 runes in the string, each rune is 1 byte)

	// The bytes have been written to buffer, not yet to file.
	// Writing from buffer to file.
	bufferedWriter.Flush()
}