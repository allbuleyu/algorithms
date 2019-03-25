package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"os"
)

var (
	pkg = flag.String("p", "0", "pkg name:example -p=0001 package = prob0001")
	name = flag.String("m", "", "file name")

	// 主目录
	base = "./leetcode"
)

func main()  {
	flag.Parse()

	if *pkg == "0" {
		fmt.Println("pkg name can not be 0 or empty")
		return
	}

	if strings.Trim(*name, " ") == "" {
		fmt.Println("name can not empty")
		return
	}

	*pkg = strings.Replace(fmt.Sprintf("%4s", *pkg), " ", "0", 4)

	buf := bytes.Buffer{}
	buf.WriteString(*pkg)
	buf.WriteByte('-')
	buf.WriteString(*name)

	dname := buf.String()

	dpath := fmt.Sprintf("%s/%s", base, dname)
	err := os.Mkdir(dpath, 0766)
	//if  {
	//	fmt.Printf("%s dir is exist! \n", dpath)
	//	return
	//}

	isExist := os.IsExist(err)
	if isExist {
		panic("file or dir has been existed")
	}

	if err != nil && !isExist {
		fmt.Printf("%s \n", err)
		return
	}

	pkgStr := fmt.Sprintf("package prob%s",*pkg)
	fpath := fmt.Sprintf("%s/%s.go", dpath, *name)
	ftPath :=fmt.Sprintf("%s/%s_test.go", dpath, *name)
	f, err := os.Create(fpath)
	if err != nil  {
		fmt.Printf("mkdir err is %s \n", err)
		return
	}

	ft, err := os.Create(ftPath)
	if err != nil  {
		fmt.Printf("mkdir err is %s \n", err)
		return
	}

	f.WriteString(pkgStr)
	ft.WriteString(pkgStr)

	mdPath := fmt.Sprintf("%s/README.md", dpath)
	fmd, err := os.Create(mdPath)
	fmd.WriteString(fmt.Sprintf("# [%s-%s](https://leetcode.com/problems/%s)", *pkg, *name, *name))

	// 运行Git命令,把生成的文件加入Git中
	cmd := exec.Command("git", "add", dpath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("cmd run fail, pls check err")
		return
	}



	defer fmd.Close()
	defer ft.Close()
	defer f.Close()

	fmt.Println(*pkg, *name, dname)
}
