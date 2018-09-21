package main

import (
	"fmt"
	"os"
	"flag"
)

// 用于创建算法包文件的函数

func makefile() {

}

func main() {
	flag.Parse()
	arg := os.Args

	flag.Args()
	//args := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	fmt.Println(arg[1:], flag.Args())
}
