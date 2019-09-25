package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// StringSlice string的Slice
type StringSlice struct {
	value []string
}

// ToStringSlice string列表转为StringSlice
func ToStringSlice(value []string) *StringSlice {
	return &StringSlice{value: value}
}

// Concat 拼接
func (s *StringSlice) Concat(given []string) *StringSlice {
	value := make([]string, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *StringSlice) Drop(n int) *StringSlice {
	if n < 0 {
		n = 0
	}
	l := len(s.value) - n
	if l < 0 {
		n = len(s.value)
	}
	s.value = s.value[n:]
	return s
}

// Filter 过滤
func (s *StringSlice) Filter(fn func(int, string) bool) *StringSlice {
	value := make([]string, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *StringSlice) First(value *string) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *StringSlice) Last(value *string) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *StringSlice) Map(fn func(int, string) string) *StringSlice {
	value := make([]string, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *StringSlice) Reduce(fn func(string, string, int) string, initial string) string {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *StringSlice) Reverse() *StringSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *StringSlice) Unique() *StringSlice {
	value := make([]string, 0, len(s.value))
	seen := make(map[string]struct{})
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
func (s *StringSlice) Append(given string) *StringSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *StringSlice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *StringSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *StringSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *StringSlice) Sort() *StringSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *StringSlice) All(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *StringSlice) Any(fn func(int, string) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *StringSlice) Paginate(size int) [][]string {
	if size <= 0 {
		size = 1
	}
	var pages [][]string
	prev := -1
	for i := range s.value {
		if (i-prev) < size && i != (len(s.value)-1) {
			continue
		}
		pages = append(pages, s.value[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s *StringSlice) Preappend(given string) *StringSlice {
	value := make([]string, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *StringSlice) Max(value *string) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if *value < each {
			*value  = each
		}
	}
	return nil 
}

// Min 获取最小元素
func (s *StringSlice) Min(value *string) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[0]
	for _, each := range s.value {
		if each < *value {
			*value = each
		}
	}
	return nil
}

// Random 随机获取一个元素
func (s *StringSlice) Random(value *string) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *StringSlice) Shuffle() *StringSlice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *StringSlice) Collect() []string {
	return s.value
}
