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
	"reflect"
	"strings"
)

const (
	defaultSuffix = "_gen.go"
	builtinPkg    = "commons"
	structTplStr  = `
type {{.Name}}Stream struct{
	value	[]*{{.Name}}
}

func StreamOf{{.Name}}(value []*{{.Name}}) *{{.Name}}Stream {
	return &{{.Name}}Stream{value:value}
}

func(c *{{.Name}}Stream) Concate(given []*{{.Name}})  *{{.Name}}Stream {
	value := make([]*{{.Name}}, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *{{.Name}}Stream) Drop(n int)  *{{.Name}}Stream {
	l := len(c.value) - n
	if l {{.Lt}} 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *{{.Name}}Stream) Filter(fn func(int, *{{.Name}})bool)  *{{.Name}}Stream {
	value := make([]*{{.Name}}, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *{{.Name}}Stream) First() *{{.Name}} {
	if len(c.value) {{.Lt}} 0 {
		return nil
	} 
	return c.value[0]
}

func(c *{{.Name}}Stream) Last() *{{.Name}} {
	if len(c.value) {{.Lt}} 0 {
		return nil
	} 
	return c.value[len(c.value)-1]
}

func(c *{{.Name}}Stream) Map(fn func(int, *{{.Name}})) *{{.Name}}Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *{{.Name}}Stream) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}},initial *{{.Name}}) *{{.Name}}   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *{{.Name}}Stream) Reverse()  *{{.Name}}Stream {
	value := make([]*{{.Name}}, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *{{.Name}}Stream) UniqueBy(compare func(*{{.Name}},*{{.Name}})bool)  *{{.Name}}Stream{
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
			if compare(inner,outter) {
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

func(c *{{.Name}}Stream) Append(given *{{.Name}}) *{{.Name}}Stream {
	c.value=append(c.value,given)
	return c
}

func(c *{{.Name}}Stream) Len() int {
	return len(c.value)
}

func(c *{{.Name}}Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *{{.Name}}Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *{{.Name}}Stream)  SortBy(less func(*{{.Name}},*{{.Name}})bool)  *{{.Name}}Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *{{.Name}}Stream) All(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *{{.Name}}Stream) Any(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *{{.Name}}Stream) Paginate(size int)  [][]*{{.Name}} {
	var pages  [][]*{{.Name}}
	prev := -1
	for i := range c.value {
		if (i-prev) {{.Lt}} size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *{{.Name}}Stream) Pop() *{{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value[lastIdx]=nil
	c.value=c.value[:lastIdx]
	return val
}

func(c *{{.Name}}Stream) Prepend(given *{{.Name}}) *{{.Name}}Stream {
	c.value = append([]*{{.Name}}{given},c.value...)
	return c
}

func(c *{{.Name}}Stream) Max(bigger func(*{{.Name}},*{{.Name}})bool) *{{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	var max *{{.Name}}
	for _,each := range c.value {
		if max==nil{
			max=each
			continue
		}
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(c *{{.Name}}Stream) Min(less func(*{{.Name}},*{{.Name}})bool) *{{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	var min *{{.Name}}
	for _,each := range c.value {
		if min==nil{
			min=each
			continue
		}
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(c *{{.Name}}Stream) Random() *{{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *{{.Name}}Stream) Shuffle() *{{.Name}}Stream {
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

{{range $idx,$each := .Sorts}}
func(c *{{$.Name}}Stream)  SortBy{{$each.Name}}(less func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i].{{$each.Name}},c.value[j].{{$each.Name}})
	})
	return c 
}
{{end}}

{{range $idx,$each := .Uniques}}
func(c *{{$.Name}}Stream)  UniqueBy{{$each.Name}}(compare func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}Stream {
	value := make([]*{{$.Name}}, 0, len(c.value))
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
			if compare(inner.{{.Name}},outter.{{.Name}}) {
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
{{end}}



func(c *{{.Name}}Stream) Collect() []*{{.Name}}{
	return c.value
}
`

	builtinTplStr = `
package {{.Pkg}}

import (
	"sort"
	"math/rand"
)

const Empty{{.TitleName}} {{.Name}} ={{.Empty}}

type {{.TitleName}}Stream struct{
	value	[]{{.Name}}
}

func StreamOf{{.TitleName}}(value []{{.Name}}) *{{.TitleName}}Stream {
	return &{{.TitleName}}Stream{value:value}
}

func(c *{{.TitleName}}Stream) Concate(given []{{.Name}})  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *{{.TitleName}}Stream) Drop(n int)  *{{.TitleName}}Stream {
	l := len(c.value) - n
	if l {{.Lt}} 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *{{.TitleName}}Stream) Filter(fn func(int, {{.Name}})bool)  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *{{.TitleName}}Stream) First() {{.Name}} {
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}}
	} 
	return c.value[0]
}

func(c *{{.TitleName}}Stream) Last() {{.Name}} {
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}}
	} 
	return c.value[len(c.value)-1]
}

func(c *{{.TitleName}}Stream) Map(fn func(int, {{.Name}})) *{{.TitleName}}Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *{{.TitleName}}Stream) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}},initial {{.Name}}) {{.Name}}   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *{{.TitleName}}Stream) Reverse()  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *{{.TitleName}}Stream) Unique()  *{{.TitleName}}Stream{
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

func(c *{{.TitleName}}Stream) Append(given {{.Name}}) *{{.TitleName}}Stream {
	c.value=append(c.value,given)
	return c
}

func(c *{{.TitleName}}Stream) Len() int {
	return len(c.value)
}

func(c *{{.TitleName}}Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *{{.TitleName}}Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *{{.TitleName}}Stream)  SortBy(less func({{.Name}},{{.Name}}) bool )  *{{.TitleName}}Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *{{.TitleName}}Stream) All(fn func(int, {{.Name}})bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *{{.TitleName}}Stream) Any(fn func(int, {{.Name}})bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *{{.TitleName}}Stream) Paginate(size int)  [][]{{.Name}} {
	var pages  [][]{{.Name}}
	prev := -1
	for i := range c.value {
		if (i-prev) {{.Lt}} size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *{{.TitleName}}Stream) Pop() {{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}} 
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *{{.TitleName}}Stream) Prepend(given {{.Name}}) *{{.TitleName}}Stream {
	c.value = append([]{{.Name}}{given},c.value...)
	return c
}

func(c *{{.TitleName}}Stream) Max() {{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}} 
	}
	var max {{.Name}}
	for idx,each := range c.value {
		if idx==0{
			max=each
			continue
		}
		if max {{.Lt}} each {
			max = each
		}
	}
	return max
}


func(c *{{.TitleName}}Stream) Min() {{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}} 
	}
	var min {{.Name}}
	for idx,each := range c.value {
		if idx==0{
			min=each
			continue
		}
		if each  {{.Lt}} min {
			min = each
		}
	}
	return min
}

func(c *{{.TitleName}}Stream) Random() {{.Name}}{
	if len(c.value) {{.Lt}} 0 {
		return Empty{{.TitleName}} 
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *{{.TitleName}}Stream) Shuffle() *{{.TitleName}}Stream {
	if len(c.value) {{.Lt}} 0 {
		return nil
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

func(c *{{.TitleName}}Stream) Collect() []{{.Name}}{
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
	curSorts     []SortInfo
	curUniques   []UniqueInfo
	builtin      bool
)

type SortInfo struct {
	Name string
	Type string
}

type UniqueInfo struct {
	Name string
	Type string
}

type tpl struct {
	Pkg       string
	Name      string
	Lt        template.HTML
	TitleName string
	Sorts     []SortInfo
	Uniques   []UniqueInfo
	Builtin   bool
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
			rd := io.MultiReader(strings.NewReader(fmt.Sprintf(`package %s
				import (
					"sort"
					"math/rand"
				)`, p.Name)), buf)
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
	for _, field := range fields.List {
		if field.Tag == nil {
			continue
		}
		ts, ok := field.Names[0].Obj.Decl.(*ast.Field)
		if !ok {
			continue
		}
		var typ string
		ident, _ := ts.Type.(*ast.Ident)
		if ident != nil {
			typ = ident.Name
		} else {
			ident, _ = ts.Type.(*ast.StarExpr).X.(*ast.Ident)
			typ = "*" + ident.Name
		}
		allTags := strings.TrimSuffix(strings.TrimPrefix(field.Tag.Value, "`"), "`")
		collectionTag := reflect.StructTag(allTags).Get("collections")
		if strings.Contains(collectionTag, "sort") {
			curSorts = append(curSorts, SortInfo{Name: field.Names[0].Name, Type: typ})
		}
		if strings.Contains(collectionTag, "unique") {
			curUniques = append(curUniques, UniqueInfo{Name: field.Names[0].Name, Type: typ})
		}
	}

}
func execTpl(buf io.Writer) error {
	tpl := tpl{Name: curStruct, Pkg: curPkg, Lt: template.HTML("<"), Empty: template.HTML(curEmpty), TitleName: curTitleName, Sorts: curSorts, Uniques: curUniques}
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
