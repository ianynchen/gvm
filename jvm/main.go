package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ianynchen/gvm/jvm/classfile"
)

func main() {
	class := classfile.NewClass()
	f, err := os.Open("/Users/u6065224/programming/go/src/github.com/ianynchen/gvm/jvm/classfile/HelloWorld.class")

	if err != nil {
		fmt.Println(fmt.Sprintf("error encounterd while opening classfile: %v", err))
	}
	defer f.Close()
	content, err1 := ioutil.ReadAll(f)

	if err1 != nil {
		fmt.Println(fmt.Sprintf("error encounterd while reading classfile: %v", err))
	}
	err = class.Parse(content, 0)

	if err != nil {
		fmt.Println("error encounterd parsing")
	}
	class.Print(os.Stdout)
}
