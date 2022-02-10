package examplejson

import (
	"bufio"
	"fmt"
	"os"
)

func InitExampleFileWrite() {
	data := []byte("hello golang!\n")

	err := os.WriteFile("./example_json/text1", data, 0644) //to write new file
	checkError(err)
	// we get the check error access coz on same package see example.go file for check error
	f, err := os.Create("./example_json/text2")
	checkError(err)
	defer f.Close()

	//data2 := []byte{11, 12, 13, 14, 45, 64, 77, 97, 111, 109}
	data2 := "hello here\n"
	n, err := f.Write([]byte(data2))
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n)

	n2, err := f.WriteString("string here\n")
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)
	f.Sync()

	// new way to write the file
	w := bufio.NewWriter(f) // f is from line 17
	n3, err := w.WriteString("witing using bufio")
	checkError(err)
	fmt.Printf("wrote using bufio with %d bytes", n3)
	w.Flush()

}
