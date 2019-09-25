package main

const builtinTplStr = `package {{.Pkg}}
	
import (
	"errors"
	"math/rand"
	"sort"
)

// {{.TitleName}}Slice {{.Name}}的Slice
type {{.TitleName}}Slice []{{.Name}}

// Concat 拼接
func (s {{.TitleName}}Slice) Concat(given []{{.Name}}) {{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return {{.TitleName}}Slice(value)
}

// Drop 丢弃前n个
func (s {{.TitleName}}Slice) Drop(n int) {{.TitleName}}Slice {
	if n {{.Lt}} 0 {
		n = 0
	}
	l := len(s) - n
	if l {{.Lt}} 0 {
		n = len(s)
	}
	return {{.TitleName}}Slice(s[n:])
}

// Filter 过滤
func (s {{.TitleName}}Slice) Filter(fn func(int, {{.Name}}) bool) {{.TitleName}}Slice {
	value := make([]{{.Name}}, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return {{.TitleName}}Slice(value)
}

// First 获取第一个元素
func (s {{.TitleName}}Slice) First() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s {{.TitleName}}Slice) Last() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s {{.TitleName}}Slice) Map(fn func(int, {{.Name}}) {{.Name}}) {{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return {{.TitleName}}Slice(value)
}

// Reduce reduce
func (s {{.TitleName}}Slice) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}}, initial {{.Name}}) {{.Name}} {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s {{.TitleName}}Slice) Reverse() {{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return {{.TitleName}}Slice(value)
}

// Unique 唯一
func (s {{.TitleName}}Slice) Unique() {{.TitleName}}Slice {
	value := make([]{{.Name}}, 0, len(s))
	seen := make(map[{{.Name}}]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return {{.TitleName}}Slice(value)
}

// Append 在尾部添加
func (s {{.TitleName}}Slice) Append(given {{.Name}}) {{.TitleName}}Slice {
	return append(s, given)
}

// Len 获取长度
func (s {{.TitleName}}Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s {{.TitleName}}Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s {{.TitleName}}Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s {{.TitleName}}Slice) Sort() {{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] {{.Lt}} value[j]
	})
	return {{.TitleName}}Slice(value)
}

// All 是否所有元素满足条件
func (s {{.TitleName}}Slice) All(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s {{.TitleName}}Slice) Any(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s {{.TitleName}}Slice) Paginate(size int) [][]{{.Name}} {
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
func (s {{.TitleName}}Slice) Preappend(given {{.Name}}) {{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s {{.TitleName}}Slice) Max() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if max {{.Lt}} each {
			max = each
		}
	}
	return max, nil 
}

// Min 获取最小元素
func (s {{.TitleName}}Slice) Min() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if each {{.Lt}} min {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s {{.TitleName}}Slice) Random() ({{.Name}}, error) {
	if len(s) {{.Lt}}= 0 {
		var defaultReturn {{.Name}}
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]{{.Name}}
func (s {{.TitleName}}Slice) Shuffle() {{.TitleName}}Slice {
	if len(s) {{.Lt}}= 0 {
		return s
	}
	value := make([]{{.Name}}, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return {{.TitleName}}Slice(value)
}

// Collect 获取[]{{.Name}}
func (s {{.TitleName}}Slice) Collect() []{{.Name}} {
	return s
}
`
