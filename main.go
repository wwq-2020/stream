package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	suffix     = "_gen.go"
	builtinPkg = "commons"
)

var (
	dir             string
	fieldStream     bool
	curPkg          string
	curStruct       string
	curTplStr       string
	curEmpty        string
	curTitleName    string
	commonStreamDir string
	curSorts        []SortInfo
	curUniques      []UniqueInfo
	curFields       []FieldInfo
	curHasBuiltin   bool
	builtin         bool
)

type SortInfo struct {
	Name string
	Type string
}

type UniqueInfo struct {
	Name string
	Type string
}

type FieldInfo struct {
	Type            string
	Name            string
	StreamType      string
	SkipFieldStream bool
	Pkg             string
	IsPointer       bool
	TitleType       string
	IsBuiltin       bool
}

type tpl struct {
	Pkg              string
	Name             string
	NeedCommonStream bool
	commonStreamDir  string
	Lt               template.HTML
	TitleName        string
	Sorts            []SortInfo
	Uniques          []UniqueInfo
	Builtin          bool
	Fields           []FieldInfo
	Empty            interface{}
}

func init() {
	flag.BoolVar(&builtin, "builtin", false, "-builtin=true")
	flag.BoolVar(&fieldStream, "fs", false, "-fs=true")
	flag.StringVar(&dir, "dir", ".", "-dir=.")
	flag.StringVar(&commonStreamDir, "common", "github.com/wwq1988/stream/commons", "-commons=github.com/wwq1988/stream/commons")
	flag.Parse()
}

func main() {

	if dir == "" {
		flag.PrintDefaults()
		return
	}
	if builtin {
		genBuiltin()
		return
	}
	genStruct()

}

func genStruct() {
	curTplStr = structTplStr

	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if strings.HasSuffix(name, suffix) || !strings.HasSuffix(name, ".go") {
			return nil
		}
		p, err := build.ImportDir(dir, 0)
		if err != nil {
			return err
		}
		curImport := strings.Join(p.Imports, "\n")
		curHasBuiltin = false
		curPkg = p.Name
		baseDir := filepath.Dir(path)
		dst := filepath.Join(baseDir, strings.Replace(name, ".go", suffix, -1))
		buf := bytes.NewBuffer(nil)

		if err := generate(path, buf); err != nil {
			return err
		}
		if buf.Len() != 0 {
			var importStr string = fmt.Sprintf(`package %s
			import (
				"sort"
				"math/rand"`, p.Name)
			if curHasBuiltin {
				importStr = fmt.Sprintf(`%s
						commons "%s"						
					`, importStr, commonStreamDir)

			}
			if curImport != "" {
				importStr = fmt.Sprintf(`%s
					"%s"						
				`, importStr, curImport)

			}
			importStr = fmt.Sprintf(`%s
				)`, importStr)
			rd := io.MultiReader(strings.NewReader(importStr), buf)
			bytes, _ := ioutil.ReadAll(rd)
			if err := ioutil.WriteFile(dst, bytes, 0644); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func generate(path string, buf io.Writer) error {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		if err := walkGd(gd, buf); err != nil {
			return err
		}

	}

	return nil
}

func walkGd(gd *ast.GenDecl, buf io.Writer) error {
	for _, spec := range gd.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		curStruct = ts.Name.Name
		setTagInfo(st.Fields)
		if err := execTpl(buf); err != nil {
			return err
		}
	}
	return nil
}

func setTagInfo(fields *ast.FieldList) {
	curSorts = nil
	curUniques = nil
	curFields = nil
	for _, field := range fields.List {

		ts, ok := field.Names[0].Obj.Decl.(*ast.Field)
		if !ok {
			continue
		}
		var typ string
		expr := ts.Type
		pointerStr := ""
		titleType := ""
		pkg := ""
		outter := false
	loop:
		for {
			switch t := expr.(type) {
			case *ast.Ident:
				if titleType == "" {
					titleType = strings.Title(t.Name)
				} else {
					pkg = t.Name + "."
				}
				if isBuiltIn(t.Name) {
					pkg = "commons."
				}
				typ = pointerStr + t.Name + typ
				break loop
			case *ast.StarExpr:
				expr = t.X

				pointerStr = "*"
			case *ast.SelectorExpr:
				outter = true
				expr = t.X
				typ = "." + t.Sel.Name
				titleType = strings.Title(t.Sel.Name)
			}
		}

		isBuiltIn := isBuiltIn(typ)
		fi := FieldInfo{Name: field.Names[0].Name, Type: typ, TitleType: titleType, SkipFieldStream: !fieldStream && outter, Pkg: pkg, IsPointer: pointerStr == "*", IsBuiltin: isBuiltIn && pointerStr == ""}
		curFields = append(curFields, fi)

		if isBuiltIn && !curHasBuiltin {
			curHasBuiltin = true
		}
	}

}

func isBuiltIn(typ string) bool {
	switch typ {
	case "string", "int", "int8", "int32", "int64", "uint", "uint32", "uint64", "uint8", "float32", "float64":
		return true
	case "*string", "*int", "*int8", "*int32", "*int64", "*uint", "*uint32", "*uint64", "*uint8", "*float32", "*float64":
		return true
	default:
		return false
	}
}
func execTpl(buf io.Writer) error {
	tpl := tpl{Name: curStruct, Pkg: curPkg, Lt: template.HTML("<"), Empty: template.HTML(curEmpty), TitleName: curTitleName, Fields: curFields}
	t, err := template.New("stream").Parse(curTplStr)
	if err != nil {
		return err
	}
	if err := t.Execute(buf, tpl); err != nil {
		return err
	}
	return nil
}

func genBuiltin() {
	curTplStr = builtinTplStr
	curPkg = builtinPkg
	buf := bytes.NewBuffer(nil)
	for _, each := range []string{"string", "int", "int8", "int32", "int64", "uint", "uint32", "uint64", "uint8", "float32", "float64"} {
		curTitleName = strings.Title(each)
		setEmpty(each)
		path := filepath.Join(dir, each+"s.go")
		curStruct = each
		if err := execTpl(buf); err != nil {
			fmt.Println(err)
			return
		}

		if len(buf.Bytes()) != 0 {
			if err := ioutil.WriteFile(path, buf.Bytes(), 0644); err != nil {
				fmt.Println(err)
				return
			}
		}
		buf.Reset()
	}

}

func setEmpty(typ string) {
	switch typ {
	case "string":
		curEmpty = `""`
	case "int", "int8", "int32", "int64", "uint", "uint32", "uint64", "uint8":
		curEmpty = "0"
	case "float32", "float64":
		curEmpty = "0.0"
	}
}
