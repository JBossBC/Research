package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func unname_pipe() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var output bytes.Buffer
	cmd1.Stdout = &output
	if err := cmd1.Start(); err != nil {
		fmt.Printf("error: the first command cant be startup:%v\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("error: could wait for the first command:%v\n", err)
		return
	}
	var outputSecond bytes.Buffer
	cmd2.Stdin = &output
	cmd2.Stdout = &outputSecond
	if err := cmd2.Start(); err != nil {
		fmt.Printf("error: the second command cant be start up %v\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("error: could wait for the second command:%v\n", err)
		return
	}
	fmt.Println("hello world")
	value, _ := ioutil.ReadAll(&outputSecond)
	fmt.Println(string(value))
}
