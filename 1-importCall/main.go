package main

import (
	"usage/utils"

	"github.com/goplus/gogen"
)

func main() {
	pkg := gogen.NewPackage("", "temp", nil)
	c := pkg.Import("github.com/goplus/llgo/c")
	pkg.NewFunc(nil, "Fn", nil, nil, false).BodyStart(pkg).
		Val(c.Ref("Printf")).
		/**/ Val(c.Ref("Str")).
		/****/ Val("hello").
		/****/ Call(1).
		Call(1).
		EndStmt().
		End()
	utils.ToTempFolder(pkg)
}
