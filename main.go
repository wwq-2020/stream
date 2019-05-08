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
	defaultSuffix = "_gen.go"
	builtinPkg    = "commons"
	structTplStr  = `
package {{.Pkg}}
type {{.Name}}Chain struct{
	value	[]*{{.Name}}
}

func New{{.Name}}Chain(value []*{{.Name}}) *{{.Name}}Chain {
	return &{{.Name}}Chain{value:value}
}

func(c *{{.Name}}Chain) Concate(given []*{{.Name}})  *{{.Name}}Chain {
	value := make([]*{{.Name}}, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *{{.Name}}Chain) Drop(n int)  *{{.Name}}Chain {
	l := len(c.value) - n
	if l {{.Le}} 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *{{.Name}}Chain) Filter(fn func(int, *{{.Name}})bool)  *{{.Name}}Chain {
	value := make([]*{{.Name}}, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *{{.Name}}Chain) First() *{{.Name}} {
	if len(c.value) {{.Le}} 0 {
		return nil
	} 
	return c.value[0]
}

func(c *{{.Name}}Chain) Last() *{{.Name}} {
	if len(c.value) {{.Le}} 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *{{.Name}}Chain) Map(fn func(int, *{{.Name}})) *{{.Name}}Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *{{.Name}}Chain) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}},initial *{{.Name}}) *{{.Name}}   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *{{.Name}}Chain) Reverse()  *{{.Name}}Chain {
	value := make([]*{{.Name}}, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *{{.Name}}Chain) Unique()  *{{.Name}}Chain{
	value := make([]*{{.Name}}, 0, len(c.value))
	seen:=make(map[int]struct{})
	for i, outter := range c.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range c.value {
			if i==j {
				continue
			}
			if inner.Compare(outter){
				seen[j]=struct{}{}				
				dup=true
			}
		}
		if dup {
			seen[i]=struct{}{}
		}
		value=append(value,outter)			
	}
	c.value = value
	return c
}

func(c *{{.Name}}Chain) Collect() []*{{.Name}}{
	return c.value
}
`

	builtinTplStr = `
package {{.Pkg}}

const Empty{{.TitleName}} {{.Name}} ={{.Empty}}

type {{.TitleName}}Chain struct{
	value	[]{{.Name}}
}

func New{{.TitleName}}Chain(value []{{.Name}}) *{{.TitleName}}Chain {
	return &{{.TitleName}}Chain{value:value}
}

func(c *{{.TitleName}}Chain) Concate(given []{{.Name}})  *{{.TitleName}}Chain {
	value := make([]{{.Name}}, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *{{.TitleName}}Chain) Drop(n int)  *{{.TitleName}}Chain {
	l := len(c.value) - n
	if l {{.Le}} 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *{{.TitleName}}Chain) Filter(fn func(int, {{.Name}})bool)  *{{.TitleName}}Chain {
	value := make([]{{.Name}}, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *{{.TitleName}}Chain) First() {{.Name}} {
	if len(c.value) {{.Le}} 0 {
		return Empty{{.TitleName}}
	} 
	return c.value[0]
}

func(c *{{.TitleName}}Chain) Last() {{.Name}} {
	if len(c.value) {{.Le}} 0 {
		return Empty{{.TitleName}}
	} 
	return c.value[len(c.value)-1]
}

func(c *{{.TitleName}}Chain) Map(fn func(int, {{.Name}})) *{{.TitleName}}Chain {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *{{.TitleName}}Chain) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}},initial {{.Name}}) {{.Name}}   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *{{.TitleName}}Chain) Reverse()  *{{.TitleName}}Chain {
	value := make([]{{.Name}}, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *{{.TitleName}}Chain) Unique()  *{{.TitleName}}Chain{
	value := make([]{{.Name}}, 0, len(c.value))
	seen:=make(map[{{.Name}}]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *{{.TitleName}}Chain) Collect() []{{.Name}}{
	return c.value
}
`
)

var (
	dir          string
	suffix       string
	curPkg       string
	curStruct    string
	curTplStr    string
	curEmpty     string
	curTitleName string
	builtin      bool
)

type tpl struct {
	Pkg       string
	Name      string
	Le        template.HTML
	TitleName string
	Empty     interface{}
}

func init() {
	flag.BoolVar(&builtin, "builtin", false, "-builtin=true")
	flag.StringVar(&dir, "dir", ".", "-dir=.")
	flag.StringVar(&suffix, "suffix", "", "-suffix=_gen.go")
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
	if suffix == "" {
		suffix = defaultSuffix
	}
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
		curPkg = p.Name
		baseDir := filepath.Dir(path)
		dst := filepath.Join(baseDir, strings.Replace(name, ".go", suffix, -1))
		buf := bytes.NewBuffer(nil)

		if err := generate(path, buf); err != nil {
			return err
		}
		if len(buf.Bytes()) != 0 {
			if err := ioutil.WriteFile(dst, buf.Bytes(), 0644); err != nil {
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
		walkGd(gd, buf)

	}

	return nil
}

func walkGd(gd *ast.GenDecl, buf io.Writer) error {
	for _, spec := range gd.Specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		_, ok = ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		curStruct = ts.Name.Name
		if err := execTpl(buf); err != nil {
			return err
		}
	}
	return nil
}

func execTpl(buf io.Writer) error {
	tpl := tpl{Name: curStruct, Pkg: curPkg, Le: template.HTML("<="), Empty: template.HTML(curEmpty), TitleName: curTitleName}
	t, err := template.New("collection").Parse(curTplStr)
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
