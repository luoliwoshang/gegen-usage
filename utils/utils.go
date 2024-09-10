package utils

import (
	"os"
	"path/filepath"

	"github.com/goplus/gogen"
)

func ToTempFolder(pkg *gogen.Package) {
	resultDir := "result"
	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(resultDir, "temp.go")

	outfile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = pkg.WriteTo(outfile)
	if err != nil {
		panic(err)
	}
}
