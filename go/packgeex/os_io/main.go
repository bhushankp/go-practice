package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("testfile.txt")
	// create a file named testfile in the current directory if it does not exist
	if err != nil {
		fmt.Println("can not create file")
	} else {
		fmt.Println("file is created succesfully")
	}

	writer := io.Writer(file)
	n, err := writer.Write([]byte("hii this is first file for interview preperation!!!!"))
	if err != nil {
		fmt.Println("file not writing", err.Error())
	} else {
		fmt.Println("file writed succes", n)
	}
	file, err = os.Open("testfile.txt")
	if err != nil {
		fmt.Println("can not open the file")
	} else {
		fmt.Println("file is open succesfully")
	}

	reader := io.Reader(file)
	data, _ := io.ReadAll(reader)

	fmt.Println(string(data))
}
