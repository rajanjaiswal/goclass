package examplejson

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func InitExampleFile() {
	// to open file and read from open file
	f, err := os.Open("C:/Users/rkjay/go/src/github.com/rajanjaiswal/goclass/example_json/text")
	checkError(err)
	defer f.Close() //this is to close the file always put this at first

	b1 := make([]byte, 5)
	n, err := f.Read(b1) //passing bytes that is b1
	checkError(err)
	fmt.Printf("%d bytes : %s\n", n, string(b1[:n]))

	// we are using seek bar
	o2, err := f.Seek(5, 0)
	checkError(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	checkError(err)
	fmt.Printf("%d bytes @ %d %s\n", n2, o2, string(b2[:n2]))

	o3, err := f.Seek(2, 1)
	checkError(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3) // we are using io.ReadFile instead of line 29 its the same thing
	checkError(err)
	fmt.Printf("%d bytes @ %d %s\n", n3, o3, string(b3[:n3]))

	// this will take back to first file line 20
	_, err = f.Seek(0, 0)
	checkError(err)

	// new way to read the file but we still need line 17 and 18 this is buffer reader
	newReader := bufio.NewReader(f) // f is from line 17
	b4, err := newReader.Peek(5)
	checkError(err)
	fmt.Printf(" bytes : %s\n", string(b4))

}

// to  read file from text
func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	//	if err != nil {  // this is to check error but we can use function to checkerror everywhere by creating the function in line 7-9
	//		panic(err)
	checkError(err)
	fmt.Println(string(data))
	return data

}
