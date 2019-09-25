package main

const structTplStr = `

// {{.Name}}Slice {{.Name}}的Slice
type {{.Name}}Slice struct {
	value []{{.Name}}
}

// To{{.Name}}Slice {{.Name}}列表转成{{.Name}}Slice
func To{{.Name}}Slice(value []{{.Name}}) *{{.Name}}Slice {
	return &{{.Name}}Slice{value: value}
}

// Concat 拼接
func (s *{{.Name}}Slice) Concat(given []{{.Name}}) *{{.Name}}Slice {
	value := make([]{{.Name}}, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *{{.Name}}Slice) Drop(n int) *{{.Name}}Slice {
	if n {{.Lt}} 0 {
		n = 0
	}
	l := len(s.value) - n
	if l {{.Lt}} 0 {
		n = len(s.value)
	}
	s.value = s.value[n:]
	return s
}

// Filter 过滤
func (s *{{.Name}}Slice) Filter(fn func(int, {{.Name}}) bool) *{{.Name}}Slice {
	value := make([]{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

{{range $idx,$each := .Fields}}
// FilterBy{{$each.Name}} 通过过滤器过滤
func (s *{{$.Name}}Slice) FilterBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) *{{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.{{$each.Name}}) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
{{end}}

// First 获取第一个元素
func (s *{{.Name}}Slice) First(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *{{.Name}}Slice) Last(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	} 
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *{{.Name}}Slice) Map(fn func(int, {{.Name}}) {{.Name}}) *{{.Name}}Slice {
	value := make([]{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *{{.Name}}Slice) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}}, initial {{.Name}}) {{.Name}} {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *{{.Name}}Slice) Reverse() *{{.Name}}Slice {
	value := make([]{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s *{{$.Name}}Slice) UniqueBy{{$each.Name}}() *{{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s.value))
	seen := make(map[{{$each.Type}}]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.{{$each.Name}}]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.{{$each.Name}}] = struct{}{}	
	}
	s.value = value
	return s
}
{{else}}
{{if $each.IsPointer}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s *{{$.Name}}Slice) UniqueBy{{$each.Name}}(compare func({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner.{{.Name}}, outter.{{.Name}}) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}
{{else}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s *{{$.Name}}Slice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s.value {
			if i == j {
				continue
			}
			if compare(inner.{{.Name}}, outter.{{.Name}}) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value,outter)			
	}
	s.value = value
	return s
}
{{end}}
{{end}}
{{end}}

// Append 在尾部添加元素
func (s *{{.Name}}Slice) Append(given {{.Name}}) *{{.Name}}Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *{{.Name}}Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *{{.Name}}Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 判断是否非空
func (s *{{.Name}}Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// All 是否所有元素满足添加
func (s *{{.Name}}Slice) All(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *{{.Name}}Slice) Any(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *{{.Name}}Slice) Paginate(size int) [][]{{.Name}} {
	if size {{.Lt}}= 0 {
		size = 1
	}
	var pages [][]{{.Name}}
	prev := -1
	for i := range s.value {
		if (i-prev) {{.Lt}} size && i != (len(s.value)-1) {
			continue
		}
		pages = append(pages, s.value[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s *{{.Name}}Slice) Preappend(given {{.Name}}) *{{.Name}}Slice {
	value := make([]{{.Name}}, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最后元素
func (s *{{.Name}}Slice) Max(bigger func({{.Name}}, {{.Name}}) bool, value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if bigger(each, *value) {
			*value = each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *{{.Name}}Slice) Min(less func({{.Name}}, {{.Name}}) bool, value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if less(each, *value) {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *{{.Name}}Slice) Random(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *{{.Name}}Slice) Shuffle() *{{.Name}}Slice {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	return s
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// SortBy{{$each.Name}} 根据{{$each.Name}}排序
func (s *{{$.Name}}Slice) SortBy{{$each.Name}}() *{{$.Name}}Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].{{$each.Name}} {{$.Lt}} s.value[j].{{$each.Name}}
	})
	return s 
}
{{else}}
// SortBy{{$each.Name}} 根据{{$each.Name}}排序
func (s *{{$.Name}}Slice) SortBy{{$each.Name}}(less func({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].{{$each.Name}}, s.value[j].{{$each.Name}})
	})
	return s 
}
{{end}}
{{end}}



{{range $idx,$each := .Fields}}
{{if $each.SkipFieldSlice}}
{{else}}
{{if $each.IsPointer}}
// {{$each.Name}}PSlice 获取{{$each.Name}}的PSlice
func (s *{{$.Name}}Slice) {{$each.Name}}PSlice() *{{$each.Pkg}}{{$each.TitleType}}PSlice {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}To{{$each.TitleType}}PSlice(value)
	return newSlice
}
{{else}}
// {{$each.Name}}PSlice 获取{{$each.Name}}的Slice
func (s *{{$.Name}}Slice) {{$each.Name}}Slice() *{{$each.Pkg}}{{$each.TitleType}}Slice {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}To{{$each.TitleType}}Slice(value)
	return newSlice
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}s 获取{{$each.Name}}的列表
func (s *{{$.Name}}Slice) {{$each.Name}}s() []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	return value
}
{{end}}

// Collect 获取最终的列表
func (s *{{.Name}}Slice) Collect() []{{.Name}} {
	return s.value
}
	
// {{.Name}}PSlice	{{.Name}}的PSlice		
type {{.Name}}PSlice struct {
	value []*{{.Name}}
}

// To{{.Name}}PSlice {{.Name}}的指针列表转成{{.Name}}PSlice 
func To{{.Name}}PSlice(value []*{{.Name}}) *{{.Name}}PSlice {
	return &{{.Name}}PSlice{value: value}
}

// Concat 拼接
func (s *{{.Name}}PSlice) Concat(given []*{{.Name}}) *{{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *{{.Name}}PSlice) Drop(n int) *{{.Name}}PSlice {
	if n {{.Lt}} 0 {
		n = 0
	}
	l := len(s.value) - n
	if l {{.Lt}} 0 {
		n = len(s.value)
	}
	s.value = s.value[n:]
	return s
}

// Filter 过滤
func (s *{{.Name}}PSlice) Filter(fn func(int, *{{.Name}}) bool) *{{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

{{range $idx,$each := .Fields}}
// FilterBy{{$each.Name}} 通过过滤器过滤
func (s *{{$.Name}}PSlice) FilterBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) *{{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each.{{$each.Name}}) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}
{{end}}

// First 获取第一个元素
func (s *{{.Name}}PSlice) First(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *{{.Name}}PSlice) Last(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	} 
	*value = *s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *{{.Name}}PSlice) Map(fn func(int, *{{.Name}}) *{{.Name}}) *{{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *{{.Name}}PSlice) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}}, initial *{{.Name}}) *{{.Name}} {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *{{.Name}}PSlice) Reverse() *{{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[len(s.value)-1-i] = each
	}
	s.value = value
	return s
}

// UniqueBy 通过比较器唯一
func (s *{{.Name}}PSlice) UniqueBy(compare func(*{{.Name}}, *{{.Name}})bool) *{{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s.value))
	seen := make(map[int]struct{})
	for i, outter := range s.value {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s.value {
			if i == j {
				continue
			}
			if compare(inner, outter) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	return s
}

// Append 在尾部添加
func (s *{{.Name}}PSlice) Append(given *{{.Name}}) *{{.Name}}PSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *{{.Name}}PSlice) Len() int {
	return len(s.value)
}

// IsEmpty 是否为空
func (s *{{.Name}}PSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsNotEmpty 是否非空
func (s *{{.Name}}PSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// SortBy 根据比较器排序
func (s *{{.Name}}PSlice) SortBy(less func(*{{.Name}}, *{{.Name}}) bool) *{{.Name}}PSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i], s.value[j])
	})
	
	return s 
}

// All 是否所有元素满足条件
func (s *{{.Name}}PSlice) All(fn func(int, *{{.Name}}) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

{{range $idx,$each := .Fields}}
// AllBy{{$each.Name}} 是否所有元素的{{$each.Name}}满足条件
func (s *{{$.Name}}PSlice) AllBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}


{{range $idx,$each := .Fields}}
// AllBy{{$each.Name}} 是否所有元素的{{$each.Name}}满足条件
func (s *{{$.Name}}Slice) AllBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) bool {
	for i, each := range s.value {
		if !fn(i, each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}

// Any 是否有元素满足条件
func (s *{{.Name}}PSlice) Any(fn func(int, *{{.Name}}) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}


{{range $idx,$each := .Fields}}
// AnyBy{{$each.Name}} 是否有元素的{{$each.Name}}满足条件
func (s *{{$.Name}}PSlice) AnyBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) bool {
	for i, each := range s.value {
		if fn(i, each.{{$each.Name}}) {
			return true
		}
	}
	return false
}
{{end}}

{{range $idx,$each := .Fields}}
// AnyBy{{$each.Name}} 是否有元素的{{$each.Name}}满足条件
func (s *{{$.Name}}Slice) AnyBy{{$each.Name}}(fn func(int, {{$each.Type}}) bool) bool {
	for i, each := range s.value {
		if fn(i, each.{{$each.Name}}) {
			return true
		}
	}
	return false
}
{{end}}

// Paginate 分页
func (s *{{.Name}}PSlice) Paginate(size int) [][]*{{.Name}} {
	if size {{.Lt}}= 0 {
		size = 1
	}
	var pages [][]*{{.Name}}
	prev := -1
	for i := range s.value {
		if (i-prev) {{.Lt}} size && i != (len(s.value)-1) {
			continue
		}
		pages = append(pages, s.value[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s *{{.Name}}PSlice) Preappend(given *{{.Name}}) *{{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *{{.Name}}PSlice) Max(bigger func(*{{.Name}}, *{{.Name}}) bool, value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if bigger(each, value) {
			*value = *each
		}
	}
	return nil
}

// Min 获取最小元素
func (s *{{.Name}}PSlice) Min(less func(*{{.Name}}, *{{.Name}}) bool, value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = *s.value[0]
	for _, each := range s.value {
		if less(each, value) {
			*value = *each
		}
	}
	return nil
}

// Random 随机获取元素
func (s *{{.Name}}PSlice) Random(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = *s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *{{.Name}}PSlice) Shuffle() *{{.Name}}PSlice {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}
	
	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// SortBy{{$each.Name}} 根据元素的{{$each.Name}}排序
func (s *{{$.Name}}PSlice) SortBy{{$each.Name}}() *{{$.Name}}PSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i].{{$each.Name}} {{$.Lt}} s.value[j].{{$each.Name}}
	})
	return s 
}
{{else}}
// SortBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器排序
func (s *{{$.Name}}PSlice) SortBy{{$each.Name}}(less func({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}PSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return less(s.value[i].{{$each.Name}}, s.value[j].{{$each.Name}})
	})
	return s 
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}唯一
func (s *{{$.Name}}PSlice) UniqueBy{{$each.Name}}() *{{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s.value))
	seen:=make(map[{{$each.Type}}]struct{})
	for _, each := range s.value {
		if _, dup := seen[each.{{$each.Name}}]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.{{$each.Name}}] = struct{}{}	
	}
	s.value = value
	return s
}
{{else}}
{{if $each.IsPointer}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器唯一
func (s *{{$.Name}}PSlice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s.value {
			if i == j {
				continue
			}
			if compare(inner.{{.Name}}, outter.{{.Name}}) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value, outter)			
	}
	s.value = value
	
	return s
}
{{else}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器唯一
func (s *{{$.Name}}PSlice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) *{{$.Name}}PSlice {
	value := make([]{{$.Name}}, 0, len(s.value))
	seen:=make(map[int]struct{})
	for i, outter := range s.value {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s.value {
			if i == j {
				continue
			}
			if compare(inner.{{.Name}}, outter.{{.Name}}) {
				seen[j] = struct{}{}				
				dup = true
			}
		}
		if dup {
			seen[i] = struct{}{}
		}
		value = append(value,outter)			
	}
	s.value = value
	
	return s
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.SkipFieldSlice}}
{{else}}
{{if $each.IsPointer}}
// {{$each.Name}}PSlice 获取{{$each.Name}}的PSlice
func (s *{{$.Name}}PSlice) {{$each.Name}}PSlice() *{{$each.Pkg}}{{$each.TitleType}}PSlice {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}To{{$each.TitleType}}PSlice(value)
	return newSlice
}
{{else}}
// {{$each.Name}}Slice 获取{{$each.Name}}的Slice
func (s *{{$.Name}}PSlice) {{$each.Name}}Slice() *{{$each.Pkg}}{{$each.TitleType}}Slice {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}To{{$each.TitleType}}Slice(value)
	return newSlice
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}s 获取{{$each.Name}}列表
func (s *{{$.Name}}PSlice) {{$each.Name}}s() []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s.value))	
	for _, each := range s.value {
		value = append(value, each.{{$each.Name}})
	}
	return value
}
{{end}}

// Collect 获取列表
func (s *{{.Name}}PSlice) Collect() []*{{.Name}} {
	return s.value
}
`
