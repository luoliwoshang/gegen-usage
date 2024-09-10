package main

import (
	"go/token"
	"go/types"
	"os"
	"usage/utils"

	"github.com/goplus/gogen"
)

func main() {
	pkg := gogen.NewPackage("", "temp", nil)
	c := pkg.Import("github.com/goplus/llgo/c")

	params := types.NewTuple(
		types.NewVar(token.NoPos, pkg.Types, "message", types.Typ[types.String]), // message string
		types.NewVar(token.NoPos, pkg.Types, "start", c.Ref("Char").Type()),      //start (int8) c.Char

		// types.NewVar(token.NoPos, pkg.Types, "start", c.Ref("Char").Type().(*types.Named))
		// panic: interface conversion: types.Type is *types.Basic, not *types.Named
	)

	simpleFunc := pkg.NewFunc(nil, "SimpleFunction", params, nil, false)
	simpleFunc.BodyStart(pkg).End()

	pkg.WriteTo(os.Stdout)
	utils.ToTempFolder(pkg)
}
