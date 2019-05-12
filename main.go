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
	suffix       = "_gen.go"
	builtinPkg   = "commons"
	structTplStr = `
	type {{.Name}}Stream struct{
		value	[]{{.Name}}
		defaultReturn {{.Name}}
	}
	
	func StreamOf{{.Name}}(value []{{.Name}}) *{{.Name}}Stream {
		return &{{.Name}}Stream{value:value, defaultReturn:{{.Name}}{}}
	}

	func(s *{{.Name}}Stream) OrElse(defaultReturn {{.Name}})  *{{.Name}}Stream {
		s.defaultReturn = defaultReturn
		return s
	}	

	func(s *{{.Name}}Stream) Concate(given []{{.Name}})  *{{.Name}}Stream {
		value := make([]{{.Name}}, len(s.value)+len(given))
		copy(value,s.value)
		copy(value[len(s.value):],given)
		s.value = value
		return s
	}
	
	func(s *{{.Name}}Stream) Drop(n int)  *{{.Name}}Stream {
		l := len(s.value) - n
		if l {{.Lt}} 0 {
			l = 0
		}
		s.value = s.value[len(s.value)-l:]
		return s
	}
	
	func(s *{{.Name}}Stream) Filter(fn func(int, {{.Name}})bool)  *{{.Name}}Stream {
		value := make([]{{.Name}}, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each){
				value = append(value,each)
			}
		}
		s.value = value
		return s
	}
	
	func(s *{{.Name}}Stream) First() {{.Name}} {
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		} 
		return s.value[0]
	}
	
	func(s *{{.Name}}Stream) Last() {{.Name}} {
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		} 
		return s.value[len(s.value)-1]
	}
	
	func(s *{{.Name}}Stream) Map(fn func(int, {{.Name}})) *{{.Name}}Stream {
		for i, each := range s.value {
			fn(i,each)
		}
		return s
	}
	
	func(s *{{.Name}}Stream) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}},initial {{.Name}}) {{.Name}}   {
		final := initial
		for i, each := range s.value {
			final = fn(final,each,i)
		}
		return final
	}
	
	func(s *{{.Name}}Stream) Reverse()  *{{.Name}}Stream {
		value := make([]{{.Name}}, len(s.value))
		for i, each := range s.value {
			value[len(s.value)-1-i] = each
		}
		s.value = value
		return s
	}
	
	func(s *{{.Name}}Stream) UniqueBy(compare func({{.Name}},{{.Name}})bool)  *{{.Name}}Stream{
		value := make([]{{.Name}}, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		return s
	}
	
	func(s *{{.Name}}Stream) Append(given {{.Name}}) *{{.Name}}Stream {
		s.value=append(s.value,given)
		return s
	}
	
	func(s *{{.Name}}Stream) Len() int {
		return len(s.value)
	}
	
	func(s *{{.Name}}Stream) IsEmpty() bool {
		return len(s.value) == 0
	}
	
	func(s *{{.Name}}Stream) IsNotEmpty() bool {
		return len(s.value) != 0
	}
	
	func(s *{{.Name}}Stream)  SortBy(less func({{.Name}},{{.Name}})bool)  *{{.Name}}Stream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i],s.value[j])
		})
		return s 
	}
	
	func(s *{{.Name}}Stream) All(fn func(int, {{.Name}})bool)  bool {
		for i, each := range s.value {
			if !fn(i,each){
				return false
			}
		}
		return true
	}
	
	func(s *{{.Name}}Stream) Any(fn func(int, {{.Name}})bool)  bool {
		for i, each := range s.value {
			if fn(i,each){
				return true
			}
		}
		return false
	}
	
	func(s *{{.Name}}Stream) Paginate(size int)  [][]{{.Name}} {
		var pages  [][]{{.Name}}
		prev := -1
		for i := range s.value {
			if (i-prev) {{.Lt}} size-1 && i != (len(s.value)-1) {
				continue
			}
			pages=append(pages,s.value[prev+1:i+1])
			prev=i
		}
		return pages
	}
	
	func(s *{{.Name}}Stream) Pop() {{.Name}}{
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		}
		lastIdx := len(s.value)-1
		val:=s.value[lastIdx]
		s.value[lastIdx]=s.defaultReturn
		s.value=s.value[:lastIdx]
		return val
	}
	
	func(s *{{.Name}}Stream) Prepend(given {{.Name}}) *{{.Name}}Stream {
		s.value = append([]{{.Name}}{given},s.value...)
		return s
	}
	
	func(s *{{.Name}}Stream) Max(bigger func({{.Name}},{{.Name}})bool) {{.Name}}{
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		}
		var max {{.Name}} = s.value[0]
		for _,each := range s.value {
			if bigger(each, max) {
				max = each
			}
		}
		return max
	}
	
	
	func(s *{{.Name}}Stream) Min(less func({{.Name}},{{.Name}})bool) {{.Name}}{
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		}
		var min {{.Name}} = s.value[0]
		for _,each := range s.value {
			if less(each, min) {
				min = each
			}
		}
		return min
	}
	
	func(s *{{.Name}}Stream) Random() {{.Name}}{
		if len(s.value) {{.Lt}}= 0 {
			return s.defaultReturn
		}
		n := rand.Intn(len(s.value))
		return s.value[n]
	}
	
	func(s *{{.Name}}Stream) Shuffle() *{{.Name}}Stream {
		if len(s.value) {{.Lt}}= 0 {
			return s
		}
		indexes := make([]int, len(s.value))
		for i := range s.value {
			indexes[i] = i
		}
		
		rand.Shuffle(len(s.value), func(i, j int) {
			s.value[i], s.value[j] = 	s.value[j], s.value[i] 
		})
		
		return s
	}
	
	{{range $idx,$each := .Sorts}}
	func(s *{{$.Name}}Stream)  SortBy{{$each.Name}}(less func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}Stream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i].{{$each.Name}},s.value[j].{{$each.Name}})
		})
		return s 
	}
	{{end}}
	
	{{range $idx,$each := .Uniques}}
	func(s *{{$.Name}}Stream)  UniqueBy{{$each.Name}}(compare func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}Stream {
		value := make([]{{$.Name}}, 0, len(s.value))
		seen:=make(map[int]struct{})
		for i, outter := range s.value {
			dup:=false
			if _,exist:=seen[i];exist{
				continue
			}		
			for j,inner :=range s.value {
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
		s.value = value
		
		return s
	}
	{{end}}
	
	{{range $idx,$each := .Fields}}
	{{if $each.SkipFieldStream}}
	{{else}}
	{{if $each.IsPointer}}
	func(s *{{$.Name}}Stream)  {{$each.Name}}PStream()  *{{$each.Pkg}}{{$each.TitleType}}PStream {	
		value := make([]{{$each.Type}}, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.{{$each.Name}})
		}
		newStream := {{$each.Pkg}}PStreamOf{{$each.TitleType}}(value)
		return newStream
	}
	{{else}}
	func(s *{{$.Name}}Stream)  {{$each.Name}}Stream()  *{{$each.Pkg}}{{$each.TitleType}}Stream {	
		value := make([]{{$each.Type}}, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.{{$each.Name}})
		}
		newStream := {{$each.Pkg}}StreamOf{{$each.TitleType}}(value)
		return newStream
	}
	{{end}}
	{{end}}
	{{end}}
	
	{{range $idx,$each := .Fields}}
	func(s *{{$.Name}}Stream)  {{$each.Name}}s()  []{{$each.Type}} {	
		value := make([]{{$each.Type}}, 0, len(s.value))	
		for _, each := range s.value {
			value = append(value, each.{{$each.Name}})
		}
		return value
	}
	{{end}}
	
	func(s *{{.Name}}Stream) Collect() []{{.Name}}{
		return s.value
	}

type {{.Name}}PStream struct{
	value	[]*{{.Name}}
	defaultReturn *{{.Name}}
}

func PStreamOf{{.Name}}(value []*{{.Name}}) *{{.Name}}PStream {
	return &{{.Name}}PStream{value:value,defaultReturn:nil}
}
func(s *{{.Name}}PStream) OrElse(defaultReturn *{{.Name}})  *{{.Name}}PStream {
	s.defaultReturn = defaultReturn
	return s
}

func(s *{{.Name}}PStream) Concate(given []*{{.Name}})  *{{.Name}}PStream {
	value := make([]*{{.Name}}, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *{{.Name}}PStream) Drop(n int)  *{{.Name}}PStream {
	l := len(s.value) - n
	if l {{.Lt}} 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *{{.Name}}PStream) Filter(fn func(int, *{{.Name}})bool)  *{{.Name}}PStream {
	value := make([]*{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *{{.Name}}PStream) First() *{{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn 
	} 
	return s.value[0]
}

func(s *{{.Name}}PStream) Last() *{{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn 
	} 
	return s.value[len(s.value)-1]
}

func(s *{{.Name}}PStream) Map(fn func(int, *{{.Name}})) *{{.Name}}PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *{{.Name}}PStream) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}},initial *{{.Name}}) *{{.Name}}   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *{{.Name}}PStream) Reverse()  *{{.Name}}PStream {
	value := make([]*{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *{{.Name}}PStream) UniqueBy(compare func(*{{.Name}},*{{.Name}})bool)  *{{.Name}}PStream{
	value := make([]*{{.Name}}, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	return s
}

func(s *{{.Name}}PStream) Append(given *{{.Name}}) *{{.Name}}PStream {
	s.value=append(s.value,given)
	return s
}

func(s *{{.Name}}PStream) Len() int {
	return len(s.value)
}

func(s *{{.Name}}PStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *{{.Name}}PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *{{.Name}}PStream)  SortBy(less func(*{{.Name}},*{{.Name}})bool)  *{{.Name}}PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}

func(s *{{.Name}}PStream) All(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *{{.Name}}PStream) Any(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *{{.Name}}PStream) Paginate(size int)  [][]*{{.Name}} {
	var pages  [][]*{{.Name}}
	prev := -1
	for i := range s.value {
		if (i-prev) {{.Lt}} size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *{{.Name}}PStream) Pop() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value[lastIdx]=s.defaultReturn
	s.value=s.value[:lastIdx]
	return val
}

func(s *{{.Name}}PStream) Prepend(given *{{.Name}}) *{{.Name}}PStream {
	s.value = append([]*{{.Name}}{given},s.value...)
	return s
}

func(s *{{.Name}}PStream) Max(bigger func(*{{.Name}},*{{.Name}})bool) *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var max *{{.Name}}  = s.value[0]
	for _,each := range s.value {
		if bigger(each, max) {
			max = each
		}
	}
	return max
}


func(s *{{.Name}}PStream) Min(less func(*{{.Name}},*{{.Name}})bool) *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var min *{{.Name}} = s.value[0]
	for _,each := range s.value {
		if less(each, min) {
			min = each
		}
	}
	return min
}

func(s *{{.Name}}PStream) Random() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *{{.Name}}PStream) Shuffle() *{{.Name}}PStream {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}

{{range $idx,$each := .Sorts}}
func(s *{{$.Name}}PStream)  SortBy{{$each.Name}}(less func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i].{{$each.Name}},s.value[j].{{$each.Name}})
	})
	return s 
}
{{end}}

{{range $idx,$each := .Uniques}}
func(s *{{$.Name}}PStream)  UniqueBy{{$each.Name}}(compare func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}PStream {
	value := make([]*{{$.Name}}, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _,exist:=seen[i];exist{
			continue
		}		
		for j,inner :=range s.value {
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
	s.value = value
	
	return s
}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.SkipFieldStream}}
{{else}}
{{if $each.IsPointer}}
func(s *{{$.Name}}PStream)  {{$each.Name}}PStream()  *{{$each.Pkg}}{{$each.TitleType}}PStream {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newStream := {{$each.Pkg}}PStreamOf{{$each.TitleType}}(value)
	return newStream
}
{{else}}
func(s *{{$.Name}}PStream)  {{$each.Name}}Stream()  *{{$each.Pkg}}{{$each.TitleType}}Stream {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newStream := {{$each.Pkg}}StreamOf{{$each.TitleType}}(value)
	return newStream
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
func(s *{{$.Name}}PStream)  {{$each.Name}}s()  []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	return value
}
{{end}}

func(s *{{.Name}}PStream) Collect() []*{{.Name}}{
	return s.value
}
`

	builtinTplStr = `
package {{.Pkg}}

import (
	"sort"
	"math/rand"
)

type {{.TitleName}}Stream struct{
	value	[]{{.Name}}
	defaultReturn {{.Name}}
}

func StreamOf{{.TitleName}}(value []{{.Name}}) *{{.TitleName}}Stream {
	return &{{.TitleName}}Stream{value:value,defaultReturn:{{.Empty}}}
}

func(s *{{.TitleName}}Stream) OrElase(defaultReturn {{.Name}})  *{{.TitleName}}Stream {
	s.defaultReturn = defaultReturn
	return s
}


func(s *{{.TitleName}}Stream) Concate(given []{{.Name}})  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *{{.TitleName}}Stream) Drop(n int)  *{{.TitleName}}Stream {
	l := len(s.value) - n
	if l {{.Lt}} 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *{{.TitleName}}Stream) Filter(fn func(int, {{.Name}})bool)  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}Stream) First() {{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}

func(s *{{.TitleName}}Stream) Last() {{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}

func(s *{{.TitleName}}Stream) Map(fn func(int, {{.Name}})) *{{.TitleName}}Stream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *{{.TitleName}}Stream) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}},initial {{.Name}}) {{.Name}}   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *{{.TitleName}}Stream) Reverse()  *{{.TitleName}}Stream {
	value := make([]{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}Stream) Unique()  *{{.TitleName}}Stream{
	value := make([]{{.Name}}, 0, len(s.value))
	seen:=make(map[{{.Name}}]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}Stream) Append(given {{.Name}}) *{{.TitleName}}Stream {
	s.value=append(s.value,given)
	return s
}

func(s *{{.TitleName}}Stream) Len() int {
	return len(s.value)
}

func(s *{{.TitleName}}Stream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *{{.TitleName}}Stream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *{{.TitleName}}Stream)  Sort()  *{{.TitleName}}Stream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i] {{.Lt}} s.value[j]
	})
	return s 
}

func(s *{{.TitleName}}Stream) All(fn func(int, {{.Name}})bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *{{.TitleName}}Stream) Any(fn func(int, {{.Name}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *{{.TitleName}}Stream) Paginate(size int)  [][]{{.Name}} {
	var pages  [][]{{.Name}}
	prev := -1
	for i := range s.value {
		if (i-prev) {{.Lt}} size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *{{.TitleName}}Stream) Pop() {{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}

func(s *{{.TitleName}}Stream) Prepend(given {{.Name}}) *{{.TitleName}}Stream {
	s.value = append([]{{.Name}}{given},s.value...)
	return s
}

func(s *{{.TitleName}}Stream) Max() {{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var max {{.Name}} = s.value[0]
	for _,each := range s.value {
		if max {{.Lt}} each {
			max = each
		}
	}
	return max
}


func(s *{{.TitleName}}Stream) Min() {{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var min {{.Name}} = s.value[0]
	for _,each := range s.value {
		if each  {{.Lt}} min {
			min = each
		}
	}
	return min
}

func(s *{{.TitleName}}Stream) Random() {{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *{{.TitleName}}Stream) Shuffle() *{{.TitleName}}Stream {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}

func(s *{{.TitleName}}Stream) Collect() []{{.Name}}{
	return s.value
}


type {{.TitleName}}PStream struct{
	value	[]*{{.Name}}
	defaultReturn *{{.Name}}
}

func PStreamOf{{.TitleName}}(value []*{{.Name}}) *{{.TitleName}}PStream {
	return &{{.TitleName}}PStream{value:value,defaultReturn:nil}
}

func(s *{{.TitleName}}PStream) OrElse(defaultReturn *{{.Name}})  *{{.TitleName}}PStream {
	s.defaultReturn = defaultReturn
	return s
}

func(s *{{.TitleName}}PStream) Concate(given []*{{.Name}})  *{{.TitleName}}PStream {
	value := make([]*{{.Name}}, len(s.value)+len(given))
	copy(value,s.value)
	copy(value[len(s.value):],given)
	s.value = value
	return s
}

func(s *{{.TitleName}}PStream) Drop(n int)  *{{.TitleName}}PStream {
	l := len(s.value) - n
	if l {{.Lt}} 0 {
		l = 0
	}
	s.value = s.value[len(s.value)-l:]
	return s
}

func(s *{{.TitleName}}PStream) Filter(fn func(int, *{{.Name}})bool)  *{{.TitleName}}PStream {
	value := make([]*{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}PStream) First() *{{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	} 
	return s.value[0]
}

func(s *{{.TitleName}}PStream) Last() *{{.Name}} {
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	} 
	return s.value[len(s.value)-1]
}

func(s *{{.TitleName}}PStream) Map(fn func(int, *{{.Name}})) *{{.TitleName}}PStream {
	for i, each := range s.value {
		fn(i,each)
	}
	return s
}

func(s *{{.TitleName}}PStream) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}},initial *{{.Name}}) *{{.Name}}   {
	final := initial
	for i, each := range s.value {
		final = fn(final,each,i)
	}
	return final
}

func(s *{{.TitleName}}PStream) Reverse()  *{{.TitleName}}PStream {
	value := make([]*{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}PStream) Unique()  *{{.TitleName}}PStream{
	value := make([]*{{.Name}}, 0, len(s.value))
	seen:=make(map[*{{.Name}}]struct{})
	for _, each := range s.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	s.value = value
	return s
}

func(s *{{.TitleName}}PStream) Append(given *{{.Name}}) *{{.TitleName}}PStream {
	s.value=append(s.value,given)
	return s
}

func(s *{{.TitleName}}PStream) Len() int {
	return len(s.value)
}

func(s *{{.TitleName}}PStream) IsEmpty() bool {
	return len(s.value) == 0
}

func(s *{{.TitleName}}PStream) IsNotEmpty() bool {
	return len(s.value) != 0
}

func(s *{{.TitleName}}PStream)  Sort(less func(*{{.Name}},*{{.Name}}) bool )  *{{.TitleName}}PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i],s.value[j])
	})
	return s 
}

func(s *{{.TitleName}}PStream) All(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(s *{{.TitleName}}PStream) Any(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(s *{{.TitleName}}PStream) Paginate(size int)  [][]*{{.Name}} {
	var pages  [][]*{{.Name}}
	prev := -1
	for i := range s.value {
		if (i-prev) {{.Lt}} size-1 && i != (len(s.value)-1) {
			continue
		}
		pages=append(pages,s.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(s *{{.TitleName}}PStream) Pop() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	lastIdx := len(s.value)-1
	val:=s.value[lastIdx]
	s.value=s.value[:lastIdx]
	return val
}

func(s *{{.TitleName}}PStream) Prepend(given *{{.Name}}) *{{.TitleName}}PStream {
	s.value = append([]*{{.Name}}{given},s.value...)
	return s
}

func(s *{{.TitleName}}PStream) Max() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var max *{{.Name}} = s.value[0]
	for _,each := range s.value {
		if max == nil{
			max = each
			continue
		}
		if each != nil && *max {{.Lt}}= *each {
			max = each
		}
	}
	return max
}


func(s *{{.TitleName}}PStream) Min() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	var min *{{.Name}} = s.value[0]
	for _,each := range s.value {
		if min == nil{
			min = each
			continue
		}
		if  each != nil && *each  {{.Lt}}= *min {
			min = each
		}
	}
	return min
}

func(s *{{.TitleName}}PStream) Random() *{{.Name}}{
	if len(s.value) {{.Lt}}= 0 {
		return s.defaultReturn
	}
	n := rand.Intn(len(s.value))
	return s.value[n]
}

func(s *{{.TitleName}}PStream) Shuffle() *{{.TitleName}}PStream {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}
	indexes := make([]int, len(s.value))
	for i := range s.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = 	s.value[j], s.value[i] 
	})
	
	return s
}

func(s *{{.TitleName}}PStream) Collect() []*{{.Name}}{
	return s.value
}
`
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
		fi := FieldInfo{Name: field.Names[0].Name, Type: typ, TitleType: titleType, SkipFieldStream: !fieldStream && outter, Pkg: pkg, IsPointer: pointerStr == "*"}
		curFields = append(curFields, fi)

		if isBuiltIn && !curHasBuiltin {
			curHasBuiltin = true
		}

		if field.Tag == nil {
			continue
		}
		allTags := strings.TrimSuffix(strings.TrimPrefix(field.Tag.Value, "`"), "`")
		streamTag := reflect.StructTag(allTags).Get("stream")
		if strings.Contains(streamTag, "sort") {
			curSorts = append(curSorts, SortInfo{Name: field.Names[0].Name, Type: typ})
		}
		if strings.Contains(streamTag, "unique") {
			curUniques = append(curUniques, UniqueInfo{Name: field.Names[0].Name, Type: typ})
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
	tpl := tpl{Name: curStruct, Pkg: curPkg, Lt: template.HTML("<"), Empty: template.HTML(curEmpty), TitleName: curTitleName, Sorts: curSorts, Uniques: curUniques, Fields: curFields}
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
