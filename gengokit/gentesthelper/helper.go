package gentesthelper

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"reflect"
	"runtime"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

// FuncSourceCode returns a string representing the source code of the function
// provided to it.
func FuncSourceCode(val interface{}) (string, error) {
	ptr := reflect.ValueOf(val).Pointer()
	fpath, _ := runtime.FuncForPC(ptr).FileLine(ptr)

	funcName := runtime.FuncForPC(ptr).Name()
	parts := strings.Split(funcName, ".")
	funcName = parts[len(parts)-1]

	// Parse the go file into the ast
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("ERROR: go parser couldn't parse file '%v'\n", fpath)
	}

	// Search ast for function declaration with name of function passed
	var fAst *ast.FuncDecl
	for _, decs := range fileAst.Decls {
		switch decs.(type) {
		case *ast.FuncDecl:
			f := decs.(*ast.FuncDecl)
			if f.Name.String() == funcName {
				fAst = f
			}
		}
	}
	code := bytes.NewBuffer(nil)
	err = printer.Fprint(code, fset, fAst)

	if err != nil {
		return "", fmt.Errorf("couldn't print code for func '%v': %v\n", funcName, err)
	}

	return code.String(), nil
}

// DiffStrings returns the line differences of two strings. Useful for
// examining how generated code differs from expected code.
func DiffStrings(a, b string) string {
	t := difflib.UnifiedDiff{
		A:        difflib.SplitLines(a),
		B:        difflib.SplitLines(b),
		FromFile: "A",
		ToFile:   "B",
		Context:  5,
	}
	text, _ := difflib.GetUnifiedDiffString(t)
	return text
}
