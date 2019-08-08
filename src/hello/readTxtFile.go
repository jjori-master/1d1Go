package main

import (
	"fmt"
	"io/ioutil"
)

func readIfCase01() {
	var b []byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}
}

func readIfCase02() {

	var b []byte
	var err error

	if b, err = ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}


}

func main()  {

	readIfCase01()
	readIfCase02()
}
