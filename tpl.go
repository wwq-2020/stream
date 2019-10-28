package main

const structTplStr = `

// {{.Name}}Slice {{.Name}}的Slice
type {{.Name}}Slice []{{.Name}}

// Concat 拼接
func (s {{.Name}}Slice) Concat(given []{{.Name}}) {{.Name}}Slice {
	value := make([]{{.Name}}, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return {{.Name}}Slice(value)
}

// TakeN 取前n个
func (s {{.Name}}Slice) TakeN(n int) {{.Name}}Slice {
	value := make([]{{.Name}}, 0, len(s))
	for idx, each := range s {
		if idx {{.Lt}} n {
			value = append(value, each)
		}
	}
	return {{.Name}}Slice(value)
}

// DropN 丢弃前n个
func (s {{.Name}}Slice) DropN(n int) {{.Name}}Slice {
	if n {{.Lt}} 0 {
		n = 0
	}
	l := len(s) - n
	if l {{.Lt}} 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s {{.Name}}Slice) Filter(fn func({{.Name}}) bool) {{.Name}}Slice {
	value := make([]{{.Name}}, 0, len(s))
	for _, each := range s {
		if fn(each) {
			value = append(value, each)
		}
	}
	return {{.Name}}Slice(value)
}

{{range $idx,$each := .Fields}}
// FilterBy{{$each.Name}} 通过过滤器过滤
func (s {{$.Name}}Slice) FilterBy{{$each.Name}}(fn func({{$each.Type}}) bool) {{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s))
	for _, each := range s {
		if fn(each.{{$each.Name}}) {
			value = append(value, each)
		}
	}
	return {{$.Name}}Slice(value)
}
{{end}}

// First 获取第一个元素
func (s {{.Name}}Slice) First() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	} 
	return s[0], nil
}

// Last 获取最后一个元素
func (s {{.Name}}Slice) Last(value *{{.Name}}) ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s {{.Name}}Slice) Map(fn func({{.Name}}) {{.Name}}) {{.Name}}Slice {
	value := make([]{{.Name}}, len(s))
	for i, each := range s {
		value[i] = fn(each)
	}
	return {{.Name}}Slice(value)
}

// Reduce reduce
func (s {{.Name}}Slice) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}}, initial {{.Name}}) {{.Name}} {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s {{.Name}}Slice) Reverse() {{.Name}}Slice {
	value := make([]{{.Name}}, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return {{.Name}}Slice(value)
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s {{$.Name}}Slice) UniqueBy{{$each.Name}}() {{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s))
	seen := make(map[{{$each.Type}}]struct{})
	for _, each := range s {
		if _, dup := seen[each.{{$each.Name}}]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.{{$each.Name}}] = struct{}{}	
	}
	return {{$.Name}}Slice(value)
}
{{else}}
{{if $each.IsPointer}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s {{$.Name}}Slice) UniqueBy{{$each.Name}}(compare func({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return {{$.Name}}Slice(value)
}
{{else}}
// UniqueBy{{$each.Name}} 通过{{$each.Name}}唯一
func (s {{$.Name}}Slice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}Slice {
	value := make([]{{$.Name}}, 0, len(s))
	seen:=make(map[int]struct{})
	for i, outter := range s {
		dup:=false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s {
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
	return {{$.Name}}Slice(value)
}
{{end}}
{{end}}
{{end}}

// Append 在尾部添加元素
func (s {{.Name}}Slice) Append(given {{.Name}}) {{.Name}}Slice {
	return append(s, given)
}

// Len 获取长度
func (s {{.Name}}Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s {{.Name}}Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s {{.Name}}Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// All 是否所有元素满足添加
func (s {{.Name}}Slice) All(fn func({{.Name}}) bool) bool {
	for _, each := range s {
		if !fn(each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s {{.Name}}Slice) Any(fn func({{.Name}}) bool) bool {
	for _, each := range s {
		if fn(each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s {{.Name}}Slice) Paginate(size int) [][]{{.Name}} {
	if size {{.Lt}}= 0 {
		size = 1
	}
	var pages [][]{{.Name}}
	prev := -1
	for i := range s {
		if (i-prev) {{.Lt}} size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s {{.Name}}Slice) Preappend(given {{.Name}}) {{.Name}}Slice {
	value := make([]{{.Name}}, len(s)+1)
	value = append(value, given)
	value[0] = given
	copy(value[1:], s)
	return {{.Name}}Slice(value)
}

// MaxBy 获取最后元素
func (s {{.Name}}Slice) MaxBy(bigger func({{.Name}}, {{.Name}}) bool) ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = each
		}
	}
	return max, nil
}


// MinBy 获取最小元素
func (s {{.Name}}Slice) MinBy(less func({{.Name}}, {{.Name}}) bool) ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s {{.Name}}Slice) Random() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s {{.Name}}Slice) Shuffle() {{.Name}}Slice {
	if len(s) {{.Lt}}= 0 {
		return s
	}
	
	value := make([]{{.Name}}, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return {{.Name}}Slice(value)
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// SortBy{{$each.Name}} 根据{{$each.Name}}排序
func (s {{$.Name}}Slice) SortBy{{$each.Name}}() {{$.Name}}Slice {
	value := make([]{{$.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].{{$each.Name}} {{$.Lt}} value[j].{{$each.Name}}
	})
	return {{$.Name}}Slice(value)
}
{{else}}
// SortBy{{$each.Name}} 根据{{$each.Name}}排序
func (s {{$.Name}}Slice) SortBy{{$each.Name}}(less func({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}Slice {
	value := make([]{{$.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].{{$each.Name}}, value[j].{{$each.Name}})
	})
	return {{$.Name}}Slice(value)
}
{{end}}
{{end}}



{{range $idx,$each := .Fields}}
{{if $each.SkipFieldSlice}}
{{else}}
{{if $each.IsPointer}}
// {{$each.Name}}Slice 获取{{$each.Name}}的PSlice
func (s {{$.Name}}Slice) {{$each.Name}}PSlice() {{$each.Pkg}}{{$each.TitleType}}PSlice {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}{{$each.TitleType}}PSlice(value)
	return newSlice
}
{{else}}
// {{$each.Name}}Slice 获取{{$each.Name}}的Slice
func (s {{$.Name}}Slice) {{$each.Name}}Slice() {{$each.Pkg}}{{$each.TitleType}}Slice {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}{{$each.TitleType}}Slice(value)
	return newSlice
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// {{$each.Name}}2{{$.Name}} {{$each.Name}}到{{$.Name}}的map
func (s {{$.Name}}Slice) {{$each.Name}}2{{$.Name}}() map[{{$each.Type}}]{{$.Name}} {
	result := make(map[{{$each.Type}}]{{$.Name}}, len(s))
	for _, each := range s {
		result[each.{{$each.Name}}] = each
	}
	return result
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}s 获取{{$each.Name}}的列表
func (s {{$.Name}}Slice) {{$each.Name}}s() []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	return value

}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// Max{{$each.Name}} 获取最大的{{$each.Name}}
func (s {{$.Name}}Slice) Max{{$each.Name}}() ({{$each.Type}}, error) {
	max, err := s.MaxBy(func(one, another {{$.Name}}) bool {
		return one.{{$each.Name}} > another.{{$each.Name}}
	})
	if err != nil {
		var defaultReturn {{$each.Type}}
		return defaultReturn, err
	}
	return max.{{$each.Name}}, nil
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// Min{{$each.Name}} 获取最小的{{$each.Name}}
func (s {{$.Name}}Slice) Min{{$each.Name}}() ({{$each.Type}}, error) {
	min, err := s.MaxBy(func(one, another {{$.Name}}) bool {
		return one.{{$each.Name}} {{$.Lt}} another.{{$each.Name}}
	})
	if err != nil {
		var defaultReturn {{$each.Type}}
		return defaultReturn, err
	}
	return min.{{$each.Name}}, nil
}
{{end}}
{{end}}

// Collect 获取最终的列表
func (s {{.Name}}Slice) Collect() []{{.Name}} {
	return s
}
	
// {{.Name}}PSlice {{.Name}}的PSlice		
type {{.Name}}PSlice []*{{.Name}}

// Concat 拼接
func (s {{.Name}}PSlice) Concat(given []*{{.Name}}) {{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return {{.Name}}PSlice(value)
}

// TakeN 取前n个
func (s {{.Name}}PSlice) TakeN(n int) {{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s))
	for idx, each := range s {
		if idx {{.Lt}} n {
			value = append(value, each)
		}
	}
	return {{.Name}}PSlice(value)
}

// DropN 丢弃前n个
func (s {{.Name}}PSlice) DropN(n int) {{.Name}}PSlice {
	if n {{.Lt}} 0 {
		n = 0
	}
	l := len(s) - n
	if l {{.Lt}} 0 {
		n = len(s)
	}
	return s[n:]
}

// Filter 过滤
func (s {{.Name}}PSlice) Filter(fn func(*{{.Name}}) bool) {{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s))
	for _, each := range s {
		if fn(each) {
			value = append(value, each)
		}
	}
	return {{.Name}}PSlice(value)
}

{{range $idx,$each := .Fields}}
// FilterBy{{$each.Name}} 通过过滤器过滤
func (s {{$.Name}}PSlice) FilterBy{{$each.Name}}(fn func({{$each.Type}}) bool) {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s))
	for _, each := range s {
		if fn(each.{{$each.Name}}) {
			value = append(value, each)
		}
	}
	return {{$.Name}}PSlice(value)
}
{{end}}

// First 获取第一个元素
func (s {{.Name}}PSlice) First() (*{{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		return nil, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s {{.Name}}PSlice) Last() (*{{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		return nil, errors.New("empty")
	} 
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s {{.Name}}PSlice) Map(fn func(*{{.Name}}) *{{.Name}}) {{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s))
	for i, each := range s {
		value[i] = fn(each)
	}
	return {{.Name}}PSlice(value)
}

// Reduce reduce
func (s {{.Name}}PSlice) Reduce(fn func(*{{.Name}}, *{{.Name}}, int) *{{.Name}}, initial *{{.Name}}) *{{.Name}} {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s {{.Name}}PSlice) Reverse() {{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return {{.Name}}PSlice(value)
}

// UniqueBy 通过比较器唯一
func (s {{.Name}}PSlice) UniqueBy(compare func(*{{.Name}}, *{{.Name}})bool) {{.Name}}PSlice {
	value := make([]*{{.Name}}, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return {{.Name}}PSlice(value)
}

// Append 在尾部添加
func (s {{.Name}}PSlice) Append(given *{{.Name}}) {{.Name}}PSlice {
	return append(s, given)
}

// Len 获取长度
func (s {{.Name}}PSlice) Len() int {
	return len(s)
}

// IsEmpty 是否为空
func (s {{.Name}}PSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 是否非空
func (s {{.Name}}PSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// SortBy 根据比较器排序
func (s {{.Name}}PSlice) SortBy(less func(*{{.Name}}, *{{.Name}}) bool) {{.Name}}PSlice {
	value := make([]*{{$.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i], value[j])
	})
	
	return {{.Name}}PSlice(value)
}

// All 是否所有元素满足条件
func (s {{.Name}}PSlice) All(fn func(*{{.Name}}) bool) bool {
	for _, each := range s {
		if !fn(each) {
			return false
		}
	}
	return true
}

{{range $idx,$each := .Fields}}
// AllBy{{$each.Name}} 是否所有元素的{{$each.Name}}满足条件
func (s {{$.Name}}PSlice) AllBy{{$each.Name}}(fn func({{$each.Type}}) bool) bool {
	for _, each := range s {
		if !fn(each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}


{{range $idx,$each := .Fields}}
// AllBy{{$each.Name}} 是否所有元素的{{$each.Name}}满足条件
func (s {{$.Name}}Slice) AllBy{{$each.Name}}(fn func({{$each.Type}}) bool) bool {
	for _, each := range s {
		if !fn(each.{{$each.Name}}){
			return false
		}
	}
	return true
}
{{end}}

// Any 是否有元素满足条件
func (s {{.Name}}PSlice) Any(fn func(*{{.Name}}) bool) bool {
	for _, each := range s {
		if fn(each) {
			return true
		}
	}
	return false
}


{{range $idx,$each := .Fields}}
// AnyBy{{$each.Name}} 是否有元素的{{$each.Name}}满足条件
func (s {{$.Name}}PSlice) AnyBy{{$each.Name}}(fn func({{$each.Type}}) bool) bool {
	for _, each := range s {
		if fn(each.{{$each.Name}}) {
			return true
		}
	}
	return false
}
{{end}}

{{range $idx,$each := .Fields}}
// AnyBy{{$each.Name}} 是否有元素的{{$each.Name}}满足条件
func (s {{$.Name}}Slice) AnyBy{{$each.Name}}(fn func({{$each.Type}}) bool) bool {
	for _, each := range s {
		if fn(each.{{$each.Name}}) {
			return true
		}
	}
	return false
}
{{end}}

// Paginate 分页
func (s {{.Name}}PSlice) Paginate(size int) [][]*{{.Name}} {
	if size {{.Lt}}= 0 {
		size = 1
	}
	var pages [][]*{{.Name}}
	prev := -1
	for i := range s {
		if (i-prev) {{.Lt}} size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s {{.Name}}PSlice) Preappend(given *{{.Name}}) {{.Name}}PSlice {
	value := make([]*{{.Name}}, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return {{.Name}}PSlice(value)
}

// MaxBy 获取最大元素
func (s {{.Name}}PSlice) MaxBy(bigger func(*{{.Name}}, *{{.Name}}) bool) (*{{$.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		return nil, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if bigger(each, max) {
			max = each
		}
	}
	return max, nil
}

// MinBy 获取最小元素
func (s {{.Name}}PSlice) MinBy(less func(*{{.Name}}, *{{.Name}}) bool) (*{{$.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		return nil, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取元素
func (s {{.Name}}PSlice) Random() (*{{$.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		return nil, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱列表
func (s {{.Name}}PSlice) Shuffle() {{.Name}}PSlice {
	if len(s) {{.Lt}}= 0 {
		return s
	}
	value := make([]*{{$.Name}}, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	
	return {{$.Name}}PSlice(value)
}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// SortBy{{$each.Name}} 根据元素的{{$each.Name}}排序
func (s {{$.Name}}PSlice) SortBy{{$each.Name}}() {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i].{{$each.Name}} {{$.Lt}} value[j].{{$each.Name}}
	})
	return {{$.Name}}PSlice(value)
}
{{else}}
// SortBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器排序
func (s {{$.Name}}PSlice) SortBy{{$each.Name}}(less func({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return less(value[i].{{$each.Name}}, value[j].{{$each.Name}})
	})
	return {{$.Name}}PSlice(value)
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// Max{{$each.Name}} 获取最大的{{$each.Name}}
func (s {{$.Name}}PSlice) Max{{$each.Name}}() ({{$each.Type}}, error) {
	max, err := s.MaxBy(func(one, another *{{$.Name}}) bool {
		return one.{{$each.Name}} > another.{{$each.Name}}
	})
	if err != nil {
		var defaultReturn {{$each.Type}}
		return defaultReturn, err
	}
	return max.{{$each.Name}}, nil
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// Min{{$each.Name}} 获取最小的{{$each.Name}}
func (s {{$.Name}}PSlice) Min{{$each.Name}}() ({{$each.Type}}, error) {
	min, err := s.MinBy(func(one, another *{{$.Name}}) bool {
		return one.{{$each.Name}} {{$.Lt}} another.{{$each.Name}}
	})
	if err != nil {
		var defaultReturn {{$each.Type}}
		return defaultReturn, err
	}
	return min.{{$each.Name}}, nil
}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}唯一
func (s {{$.Name}}PSlice) UniqueBy{{$each.Name}}() {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s))
	seen := make(map[{{$each.Type}}]struct{})
	for _, each := range s {
		if _, dup := seen[each.{{$each.Name}}]; dup {
			continue
		}
		value = append(value, each)
		
		seen[each.{{$each.Name}}] = struct{}{}	
	}
	return {{$.Name}}PSlice(value)
}
{{else}}
{{if $each.IsPointer}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器唯一
func (s {{$.Name}}PSlice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j,inner :=range s {
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
	return {{$.Name}}PSlice(value)
}
{{else}}
// UniqueBy{{$each.Name}} 根据元素的{{$each.Name}}和比较器唯一
func (s {{$.Name}}PSlice) UniqueBy{{$each.Name}}(compare func ({{$each.Type}}, {{$each.Type}}) bool) {{$.Name}}PSlice {
	value := make([]*{{$.Name}}, 0, len(s))
	seen := make(map[int]struct{})
	for i, outter := range s {
		dup := false
		if _, exist := seen[i]; exist {
			continue
		}		
		for j, inner := range s {
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
	return {{$.Name}}PSlice(value)
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.SkipFieldSlice}}
{{else}}
{{if $each.IsPointer}}
// {{$each.Name}}PSlice 获取{{$each.Name}}的PSlice
func (s {{$.Name}}PSlice) {{$each.Name}}PSlice() {{$each.Pkg}}{{$each.TitleType}}PSlice {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}{{$each.TitleType}}PSlice(value)
	return newSlice
}
{{else}}
// {{$each.Name}}Slice 获取{{$each.Name}}的Slice
func (s {{$.Name}}PSlice) {{$each.Name}}Slice() {{$each.Pkg}}{{$each.TitleType}}Slice {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	newSlice := {{$each.Pkg}}{{$each.TitleType}}Slice(value)
	return newSlice
}
{{end}}
{{end}}
{{end}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}s 获取{{$each.Name}}列表
func (s {{$.Name}}PSlice) {{$each.Name}}s() []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	return value

}
{{end}}

{{range $idx,$each := .Fields}}
{{if $each.IsBuiltin}}
// {{$each.Name}}2{{$.Name}} {{$each.Name}}到{{$.Name}}的map
func (s {{$.Name}}PSlice) {{$each.Name}}2{{$.Name}}() map[{{$each.Type}}]*{{$.Name}} {
	result := make(map[{{$each.Type}}]*{{$.Name}}, len(s))
	for _, each := range s {
		result[each.{{$each.Name}}] = each
	}
	return result
}
{{end}}
{{end}}

// Collect 获取列表
func (s {{.Name}}PSlice) Collect() []*{{.Name}} {
	return s
}
`
