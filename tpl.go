package main

const structTplStr = `

// {{.Name}}Slice {{.Name}}的Slice
type {{.Name}}Slice []{{.Name}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}Of{{$.Name}}Map {{$each.Name}}Of{{$.Name}}Map
type {{$each.Name}}Of{{$.Name}}Map map[{{$each.Type}}]{{$.Name}}Slice

// FlatMap FlatMap
func (m {{$each.Name}}Of{{$.Name}}Map) FlatMap(fn func([]{{$.Name}})) {
	for _, list := range m {
		fn(list)
	}
}
{{end}}

// {{.Name}}Result {{.Name}}Result
type {{.Name}}Result struct{
	value {{.Name}}
	isPresent bool 
}

// IsPresent 是否存在
func (r {{.Name}}Result) IsPresent() bool {
	return r.isPresent
}

// Get 获取值
func (r {{.Name}}Result) Get() {{.Name}} {
	return r.value
}

// Concat 拼接
func (s {{.Name}}Slice) Concat(given []{{.Name}}) {{.Name}}Slice {
	result := make([]{{.Name}}, len(s)+len(given))
	copy(result, s)
	copy(result[len(s):], given)
	return result
}

// Limit 取前n个
func (s {{.Name}}Slice) Limit(n int) {{.Name}}Slice {
	result := make([]{{.Name}}, 0, len(s))
	for idx, each := range s {
		if idx {{.Lt}} n {
			result = append(result, each)
		}
	}
	return result
}

// Peek Peek
func (s {{.Name}}Slice) Peek(fn func({{.Name}})) {
	for _, each := range s {
		fn(each)
	}
}

// Skip 丢弃前n个
func (s {{.Name}}Slice) Skip(n int) {{.Name}}Slice {
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
func (s {{.Name}}Slice) Filter(filters ...func({{.Name}}) bool) {{.Name}}Slice {
	result := make([]{{.Name}}, 0, len(s))
	for _, each := range s {
		valid := true
		for _, filter := range filters {
			if !filter(each) {
				valid = false
			}
		}
		if valid {
			result = append(result, each)
		}
	}
	return result
}

{{range $idx,$each := .Fields}}
// GroupBy{{$each.Name}} 通过{{$each.Name}}分组
func (s {{$.Name}}Slice) GroupBy{{$each.Name}}(comparator func({{$each.Type}}, {{$each.Type}}) bool) {{$each.Name}}Of{{$.Name}}Map {
	result := make(map[{{$each.Type}}]{{$.Name}}Slice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _,skip:=skip[i];skip{
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.{{$each.Name}}, inner.{{$each.Name}}) {
				skip[j] = struct{}{}
			}
		}
		result[outter.{{$each.Name}}] = append(result[outter.{{$each.Name}}], outter)
	}
	return result
}
{{end}}

// First 获取第一个元素
func (s {{.Name}}Slice) First() {{.Name}}Result {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return {{.Name}}Result{value:defaultReturn, isPresent:false}
	} 
	return {{.Name}}Result{value:s[0], isPresent:true}
}

// Last 获取最后一个元素
func (s {{.Name}}Slice) Last(value *{{.Name}}) {{.Name}}Result {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return {{.Name}}Result{value:defaultReturn, isPresent:false}
	}
	return {{.Name}}Result{value:s[len(s)-1], isPresent:true}
}

// Map 对每个元素进行操作
func (s {{.Name}}Slice) Map(fn func({{.Name}}) {{.Name}}) {{.Name}}Slice {
	result := make([]{{.Name}}, len(s))
	for i, each := range s {
		result[i] = fn(each)
	}
	return result
}

// Reduce reduce
func (s {{.Name}}Slice) Reduce(fn func({{.Name}}, {{.Name}}) {{.Name}}, initial {{.Name}}) {{.Name}} {
	final := initial
	for _, each := range s {
		final = fn(final, each)
	}
	return final
}

// Reverse 逆序
func (s {{.Name}}Slice) Reverse() {{.Name}}Slice {
	result := make([]{{.Name}}, len(s))
	for i, each := range s {
		result[len(s)-1-i] = each
	}
	return result
}

// Distinct 去重
func (s {{$.Name}}Slice) Distinct(comparator func({{$.Name}}, {{$.Name}}) bool) {{$.Name}}Slice {
	result := make({{$.Name}}Slice, 0, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip:=skip[i];skip{
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter, inner) {
				skip[j] = struct{}{}
			}
		}
		result = append(result,outter)
	}
	return result
}

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

// AllMatch 是否所有元素满足添加
func (s {{.Name}}Slice) AllMatch(matchFuncs ...func({{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if !matchFunc(each) {
				return false
			}
		}
	}
	return true
}

// AnyMatch 是否有元素满足条件
func (s {{.Name}}Slice) AnyMatch(matchFuncs ...func({{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return true
			}
		}
	}
	return false
}

// NoneMatch 是否没有元素满足条件
func (s {{.Name}}Slice) NoneMatch(matchFuncs ...func({{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return false
			}
		}
	}
	return true
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
	result := make([]{{.Name}}, len(s)+1)
	result = append(result, given)
	result[0] = given
	copy(result[1:], s)
	return result
}

// Max 获取最大元素
func (s {{.Name}}Slice) Max(comparator func({{.Name}}, {{.Name}}) bool) {{.Name}}Result {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return {{.Name}}Result{value:defaultReturn, isPresent:false}
	}
	max := s[0]
	for _, each := range s {
		if comparator(each, max) {
			max = each
		}
	}
	return {{.Name}}Result{value:max, isPresent:true}
}


// Min 获取最小元素
func (s {{.Name}}Slice) Min(less func({{.Name}}, {{.Name}}) bool) {{.Name}}Result {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return {{.Name}}Result{value:defaultReturn, isPresent:true}
	}
	min := s[0]
	for _, each := range s {
		if less(each, min) {
			min = each
		}
	}
	return {{.Name}}Result{value:min, isPresent:true}
}

// Random 随机获取一个元素
func (s {{.Name}}Slice) Random() {{.Name}}Result {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return {{.Name}}Result{value:defaultReturn, isPresent:true}
	}
	n := rand.Intn(len(s))
	return {{.Name}}Result{value:s[n], isPresent:true}
}

// Shuffle 打乱列表
func (s {{.Name}}Slice) Shuffle() {{.Name}}Slice {
	if len(s) {{.Lt}}= 0 {
		return s
	}
	
	result := make([]{{.Name}}, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i] 
	})
	return result
}


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

// Collect 获取最终的列表
func (s {{.Name}}Slice) Collect() []{{.Name}} {
	return s
}

// {{.Name}}PSlice {{.Name}}的PSlice
type {{.Name}}PSlice []*{{.Name}}

{{range $idx,$each := .Fields}}
// {{$each.Name}}Of{{$.Name}}PMap {{$each.Name}}Of{{$.Name}}PMap
type {{$each.Name}}Of{{$.Name}}PMap map[{{$each.Type}}]{{$.Name}}PSlice

// FlatMap FlatMap
func (m {{$each.Name}}Of{{$.Name}}PMap) FlatMap(fn func([]*{{$.Name}})) {
	for _, list := range m {
		fn(list)
	}
}
{{end}}

// {{.Name}}PResult {{.Name}}PResult
type {{.Name}}PResult struct{
	value *{{.Name}}
	isPresent bool 
}

// IsPresent 是否存在
func (r {{.Name}}PResult) IsPresent() bool {
	return r.isPresent
}

// Get 获取值
func (r {{.Name}}PResult) Get() *{{.Name}} {
	return r.value
}

// Concat 拼接
func (s {{.Name}}PSlice) Concat(given []*{{.Name}}) {{.Name}}PSlice {
	result := make([]*{{.Name}}, len(s)+len(given))
	copy(result, s)
	copy(result[len(s):], given)
	return result
}

// Limit 取前n个
func (s {{.Name}}PSlice) Limit(n int) {{.Name}}PSlice {
	result := make([]*{{.Name}}, 0, len(s))
	for idx, each := range s {
		if idx {{.Lt}} n {
			result = append(result, each)
		}
	}
	return result
}

// Peek Peek
func (s {{.Name}}PSlice) Peek(fn func(*{{.Name}})) {
	for _, each := range s {
		fn(each)
	}
}

// Skip 丢弃前n个
func (s {{.Name}}PSlice) Skip(n int) {{.Name}}PSlice {
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
func (s {{.Name}}PSlice) Filter(filters ...func(*{{.Name}}) bool) {{.Name}}PSlice {
	result := make([]*{{.Name}}, 0, len(s))
	for _, each := range s {
		valid := true
		for _, filter := range filters {
			if !filter(each) {
				valid = false
			}
		}
		if valid {
			result = append(result, each)
		}
	}
	return result
}

{{range $idx,$each := .Fields}}
// GroupBy{{$each.Name}} 通过{{$each.Name}}分组
func (s {{$.Name}}PSlice) GroupBy{{$each.Name}}(comparator func({{$each.Type}}, {{$each.Type}}) bool) {{$each.Name}}Of{{$.Name}}PMap {
	result := make(map[{{$each.Type}}]{{$.Name}}PSlice, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip:=skip[i];skip{
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter.{{$each.Name}}, inner.{{$each.Name}}) {
				skip[j] = struct{}{}
			}
		}
		result[outter.{{$each.Name}}] = append(result[outter.{{$each.Name}}], outter)
	}
	return result
}
{{end}}

// First 获取第一个元素
func (s {{.Name}}PSlice) First() {{.Name}}PResult {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn *{{.Name}}
		return {{.Name}}PResult{value:defaultReturn, isPresent:false}
	} 
	return {{.Name}}PResult{value:s[0], isPresent:true}
}

// Last 获取最后一个元素
func (s {{.Name}}PSlice) Last(value *{{.Name}}) {{.Name}}PResult {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn *{{.Name}}
		return {{.Name}}PResult{value:defaultReturn, isPresent:false}
	}
	return {{.Name}}PResult{value:s[len(s)-1], isPresent:true}
}

// Map 对每个元素进行操作
func (s {{.Name}}PSlice) Map(fn func(*{{.Name}}) *{{.Name}}) {{.Name}}PSlice {
	result := make([]*{{.Name}}, len(s))
	for i, each := range s {
		result[i] = fn(each)
	}
	return result
}

// Reduce reduce
func (s {{.Name}}PSlice) Reduce(fn func(*{{.Name}}, *{{.Name}}) *{{.Name}}, initial *{{.Name}}) *{{.Name}} {
	final := initial
	for _, each := range s {
		final = fn(final, each)
	}
	return final
}

// Reverse 逆序
func (s {{.Name}}PSlice) Reverse() {{.Name}}PSlice {
	result := make([]*{{.Name}}, len(s))
	for i, each := range s {
		result[len(s)-1-i] = each
	}
	return result
}

// Distinct 去重
func (s {{$.Name}}PSlice) Distinct(comparator func(*{{$.Name}}, *{{$.Name}}) bool) {{$.Name}}PSlice {
	result := make([]*{{.Name}}, len(s))
	skip := make(map[int]struct{})
	for i, outter := range s {
		if _, skip:=skip[i];skip{
			continue
		}
		for j, inner := range s[i:] {
			if comparator(outter, inner) {
				skip[j] = struct{}{}
			}
		}
		result = append(result,outter)
	}
	return result
}

// Append 在尾部添加元素
func (s {{.Name}}PSlice) Append(given *{{.Name}}) {{.Name}}PSlice {
	return append(s, given)
}

// Len 获取长度
func (s {{.Name}}PSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s {{.Name}}PSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s {{.Name}}PSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// AllMatch 是否所有元素满足添加
func (s {{.Name}}PSlice) AllMatch(matchFuncs ...func(*{{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if !matchFunc(each) {
				return false
			}
		}
	}
	return true
}

// AnyMatch 是否有元素满足条件
func (s {{.Name}}PSlice) AnyMatch(matchFuncs ...func(*{{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return true
			}
		}
	}
	return false
}

// NoneMatch 是否没有元素满足条件
func (s {{.Name}}PSlice) NoneMatch(matchFuncs ...func(*{{.Name}}) bool) bool {
	for _, each := range s {
		for _, matchFunc := range matchFuncs {
			if matchFunc(each) {
				return false
			}
		}
	}
	return true
}


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
	result := make([]*{{.Name}}, len(s)+1)
	result[0] = given
	copy(result[1:], s)
	return result
}

// Max 获取最大元素
func (s {{.Name}}PSlice) Max(comparator func(*{{.Name}}, *{{.Name}}) bool) {{.Name}}PResult {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn *{{.Name}}
		return {{.Name}}PResult{value:defaultReturn, isPresent:false}
	}
	max := s[0]
	for _, each := range s {
		if comparator(each, max) {
			max = each
		}
	}
	return {{.Name}}PResult{value:max, isPresent:true}
}


// Min 获取最小元素
func (s {{.Name}}PSlice) Min(comparator func(*{{.Name}}, *{{.Name}}) bool) {{.Name}}PResult {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn *{{.Name}}
		return {{.Name}}PResult{value:defaultReturn, isPresent:false}
	}
	min := s[0]
	for _, each := range s {
		if comparator(each, min) {
			min = each
		}
	}
	return {{.Name}}PResult{value:min, isPresent:true}
}

// Random 随机获取一个元素
func (s {{.Name}}PSlice) Random() {{.Name}}PResult {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn *{{.Name}}
		return {{.Name}}PResult{value:defaultReturn, isPresent:false}
	}
	n := rand.Intn(len(s))
	return {{.Name}}PResult{value:s[n], isPresent:true}
}

// Shuffle 打乱列表
func (s {{.Name}}PSlice) Shuffle() {{.Name}}PSlice {
	if len(s) {{.Lt}}= 0 {
		return s
	}
	
	result := make([]*{{.Name}}, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i] 
	})
	return result
}


{{range $idx,$each := .Fields}}
// {{$each.Name}}s 获取{{$each.Name}}的列表
func (s {{$.Name}}PSlice) {{$each.Name}}s() []{{$each.Type}} {	
	value := make([]{{$each.Type}}, 0, len(s))	
	for _, each := range s {
		value = append(value, each.{{$each.Name}})
	}
	return value
}
{{end}}

// Collect 获取最终的列表
func (s {{.Name}}PSlice) Collect() []*{{.Name}} {
	return s
}

`
