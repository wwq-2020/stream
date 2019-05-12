package main

const (
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

	{{range $idx,$each := .Fields}}
	func(s *{{$.Name}}Stream) FilterBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  *{{$.Name}}Stream {
		value := make([]{{$.Name}}, 0, len(s.value))
		for i, each := range s.value {
			if fn(i,each.{{$each.Name}}){
				value = append(value,each)
			}
		}
		s.value = value
		return s
	}
	{{end}}
	
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
	
	{{range $idx,$each := .Fields}}
	{{if $each.IsBuiltin}}
	func(s *{{$.Name}}Stream)  UniqueBy{{$each.Name}}()  *{{$.Name}}Stream {
		value := make([]{{$.Name}}, 0, len(s.value))
		seen:=make(map[{{$each.Type}}]struct{})
		for _, each := range s.value {
			if _,dup:=seen[each.{{$each.Name}}];dup{
				continue
			}
			value = append(value, each)
			seen[each.{{$each.Name}}]=struct{}{}	
		}
		s.value = value
		return s
	}
	{{else}}
	{{if $each.IsPointer}}
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
	{{else}}
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
	{{end}}
	{{end}}
	
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
	
	{{range $idx,$each := .Fields}}
	{{if $each.IsBuiltin}}
	func(s *{{$.Name}}Stream)  SortBy{{$each.Name}}()  *{{$.Name}}Stream {
		sort.Slice(s.value, func(i,j int)bool{
			return s.value[i].{{$each.Name}} {{$.Lt}} s.value[j].{{$each.Name}}
		})
		return s 
	}
	{{else}}
	func(s *{{$.Name}}Stream)  SortBy{{$each.Name}}(less func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}Stream {
		sort.Slice(s.value, func(i,j int)bool{
			return less(s.value[i].{{$each.Name}},s.value[j].{{$each.Name}})
		})
		return s 
	}
	{{end}}
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

{{range $idx,$each := .Fields}}
func(s *{{$.Name}}PStream) FilterBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  *{{$.Name}}PStream {
	value := make([]*{{$.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i,each.{{$each.Name}}){
			value = append(value,each)
		}
	}
	s.value = value
	return s
}
{{end}}

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


{{range $idx,$each := .Fields}}
func(s *{{$.Name}}Stream) AllBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}

func(s *{{.Name}}PStream) Any(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

{{range $idx,$each := .Fields}}
func(s *{{$.Name}}Stream) AnyBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each.{{$each.Name}}){
			return true
		}
	}
	return false
}
{{end}}

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

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
func(s *{{$.Name}}PStream)  SortBy{{$each.Name}}()  *{{$.Name}}PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return s.value[i].{{$each.Name}} {{$.Lt}} s.value[j].{{$each.Name}}
	})
	return s 
}
{{else}}
func(s *{{$.Name}}PStream)  SortBy{{$each.Name}}(less func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}PStream {
	sort.Slice(s.value, func(i,j int)bool{
		return less(s.value[i].{{$each.Name}},s.value[j].{{$each.Name}})
	})
	return s 
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
func(s *{{$.Name}}PStream)  UniqueBy{{$each.Name}}()  *{{$.Name}}PStream {
	value := make([]*{{$.Name}}, 0, len(s.value))
	seen:=make(map[{{$each.Type}}]struct{})
	for _, each := range s.value {
		if _,dup:=seen[each.{{$each.Name}}];dup{
			continue
		}
		value = append(value, each)
		seen[each.{{$each.Name}}]=struct{}{}	
	}
	s.value = value
	return s
}
{{else}}
{{if $each.IsPointer}}
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
{{else}}
func(s *{{$.Name}}PStream)  UniqueBy{{$each.Name}}(compare func({{$each.Type}},{{$each.Type}})bool)  *{{$.Name}}PStream {
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
{{end}}
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

{{range $idx,$each := .Fields}}
func(s *{{$.Name}}PStream) AllBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  bool {
	for i, each := range s.value {
		if !fn(i,each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}

func(s *{{.TitleName}}PStream) Any(fn func(int, *{{.Name}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

{{range $idx,$each := .Fields}}
func(s *{{$.Name}}Stream) AnyBy{{$each.Name}}(fn func(int,{{$each.Type}})bool)  bool {
	for i, each := range s.value {
		if fn(i,each.{{$each.Name}}){
			return true
		}
	}
	return false
}
{{end}}

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
