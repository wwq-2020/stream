package main

const builtinTplStr = `package {{.Pkg}}
	
import (
	"errors"
	"math/rand"
	"sort"
)

// {{.TitleName}}Slice {{.Name}}的Slice
type {{.TitleName}}Slice struct {
	value []{{.Name}}
}

// To{{.TitleName}}Slice {{.Name}}列表转为{{.TitleName}}Slice
func To{{.TitleName}}Slice(value []{{.Name}}) *{{.TitleName}}Slice {
	return &{{.TitleName}}Slice{value: value}
}

// Concat 拼接
func (s *{{.TitleName}}Slice) Concat(given []{{.Name}}) *{{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *{{.TitleName}}Slice) Drop(n int) *{{.TitleName}}Slice {
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
func (s *{{.TitleName}}Slice) Filter(fn func(int, {{.Name}}) bool) *{{.TitleName}}Slice {
	value := make([]{{.Name}}, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *{{.TitleName}}Slice) First(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *{{.TitleName}}Slice) Last(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *{{.TitleName}}Slice) Map(fn func(int, {{.Name}}) {{.Name}}) *{{.TitleName}}Slice {
	value := make([]{{.Name}}, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *{{.TitleName}}Slice) Reduce(fn func({{.Name}}, {{.Name}}, int) {{.Name}}, initial {{.Name}}) {{.Name}} {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *{{.TitleName}}Slice) Reverse() *{{.TitleName}}Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] {{.Ht}} s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *{{.TitleName}}Slice) Unique() *{{.TitleName}}Slice {
	value := make([]{{.Name}}, 0, len(s.value))
	seen := make(map[{{.Name}}]struct{})
	for _, each := range s.value {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	s.value = value
	return s
}

// Append 在尾部添加
func (s *{{.TitleName}}Slice) Append(given {{.Name}}) *{{.TitleName}}Slice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *{{.TitleName}}Slice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *{{.TitleName}}Slice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *{{.TitleName}}Slice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *{{.TitleName}}Slice) Sort() *{{.TitleName}}Slice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] {{.Lt}} s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *{{.TitleName}}Slice) All(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *{{.TitleName}}Slice) Any(fn func(int, {{.Name}}) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *{{.TitleName}}Slice) Paginate(size int) [][]{{.Name}} {
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
func (s *{{.TitleName}}Slice) Preappend(given {{.Name}}) *{{.TitleName}}Slice {
	value := make([]{{.Name}}, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *{{.TitleName}}Slice) Max(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if *value {{.Lt}} each {
			*value  = each
		}
	}
	return nil 
}

// Min 获取最小元素
func (s *{{.TitleName}}Slice) Min(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if each {{.Lt}} *value {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *{{.TitleName}}Slice) Random(value *{{.Name}}) error {
	if len(s.value) {{.Lt}}= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *{{.TitleName}}Slice) Shuffle() *{{.TitleName}}Slice {
	if len(s.value) {{.Lt}}= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *{{.TitleName}}Slice) Collect() []{{.Name}} {
	return s.value
}
`
