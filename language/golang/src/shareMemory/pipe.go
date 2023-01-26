package main

import (
	"fmt"
	"os"
)

func main() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("error: couldn't create the named pipe %v \n", err)
	}
	//Read与Write会在另一端还未就绪时对进程进行阻塞，所以二者需要并发运行
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf(reader.Stat())
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
}
