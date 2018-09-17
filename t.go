package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type myStruct struct {
	S string
	I int
	q string
}

//var arr []int = []int{2,4,5,6,1,2,3,6}
var arr []int = []int{2,8,7,1, 3, 5,6, 4}

func main() {
	base := "./leetcode/"
	files, _ := ioutil.ReadDir(base)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		fname := strings.Replace(f.Name(), "-easy", "", 1)
		fname = strings.Replace(fname, "-hard", "", 1)

		filename := strings.Split(fname, "-")
		filename[0] = strings.Replace(fmt.Sprintf("%4s", filename[0]), " ", "0", 10)

		pack := "package prob" + filename[0]

		dirname := strings.Replace(strings.Join(filename, "-"), ".go", "", 1)
		dirname = base+dirname
		name := strings.Join(filename[1:], "-")

		filePath := base+f.Name()
		file, err := os.OpenFile(filePath, os.O_RDWR, 0)
		if err != nil {
			fmt.Println("err is :", err)
			break
		}

		file.WriteString(pack)
		file.Close()

		err = os.Mkdir(dirname, 0755)


		//filePath = strings.Replace(filePath, "easy", "", 1)
		//filePath = strings.Replace(filePath, "hard", "", 1)

		os.Rename(filePath, dirname+"/"+name)




	}
}




