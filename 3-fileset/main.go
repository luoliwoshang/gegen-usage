package main

import (
	"os"
	"path/filepath"

	"github.com/goplus/gogen"
)

func main() {
	// 创建一个包
	pkg := gogen.NewPackage("", "mypackage", nil)
	fmt := pkg.Import("fmt")

	// 创建第一个文件
	pkg.SetCurFile("file1", true)
	pkg.NewFunc(nil, "Func1", nil, nil, false).BodyStart(pkg).
		Val(fmt.Ref("Println")).Val("Hello from file1").Call(1).EndStmt().
		End()

	ToTempFolder(pkg, "file1")

	pkg.SetCurFile("file2", true)
	pkg.NewFunc(nil, "Func2", nil, nil, false).BodyStart(pkg).
		Val(fmt.Ref("Println")).Val("Hello from file2").Call(1).EndStmt().
		End()

	ToTempFolder(pkg, "file2")

}

func ToTempFolder(pkg *gogen.Package, fname string) {
	resultDir := "result"
	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(resultDir, fname+".go")

	outfile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = pkg.WriteTo(outfile, fname)
	if err != nil {
		panic(err)
	}
}
