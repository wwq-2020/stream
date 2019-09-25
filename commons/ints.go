package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// IntSlice int的Slice
type IntSlice struct {
	value []int
}

// ToIntSlice int列表转为IntSlice
func ToIntSlice(value []int) *IntSlice {
	return &IntSlice{value: value}
}

// Concat 拼接
func (s *IntSlice) Concat(given []int) *IntSlice {
	value := make([]int, len(s.value)+len(given))
	copy(value, s.value)
	copy(value[len(s.value):], given)
	s.value = value
	return s
}

// Drop 丢弃前n个
func (s *IntSlice) Drop(n int) *IntSlice {
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
func (s *IntSlice) Filter(fn func(int, int) bool) *IntSlice {
	value := make([]int, 0, len(s.value))
	for i, each := range s.value {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	s.value = value
	return s
}

// First 获取第一个元素
func (s *IntSlice) First(value *int) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	} 
	*value = s.value[0]
	return nil
}

// Last 获取最后一个元素
func (s *IntSlice) Last(value *int) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	*value = s.value[len(s.value)-1]
	return nil
}

// Map 对每个元素进行操作
func (s *IntSlice) Map(fn func(int, int) int) *IntSlice {
	value := make([]int, len(s.value))
	for i, each := range s.value {
		value[i] = fn(i, each)
	}
	s.value = value
	return s
}

// Reduce reduce
func (s *IntSlice) Reduce(fn func(int, int, int) int, initial int) int {
	final := initial
	for i, each := range s.value {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s *IntSlice) Reverse() *IntSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] > s.value[j]
	})
	return s 
}

// Unique 唯一
func (s *IntSlice) Unique() *IntSlice {
	value := make([]int, 0, len(s.value))
	seen := make(map[int]struct{})
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
func (s *IntSlice) Append(given int) *IntSlice {
	s.value = append(s.value, given)
	return s
}

// Len 获取长度
func (s *IntSlice) Len() int {
	return len(s.value)
}

// IsEmpty 判断是否为空
func (s *IntSlice) IsEmpty() bool {
	return len(s.value) == 0
}

// IsEmpty 判断是否非空
func (s *IntSlice) IsNotEmpty() bool {
	return len(s.value) != 0
}

// Sort 排序
func (s *IntSlice) Sort() *IntSlice {
	sort.Slice(s.value, func(i, j int) bool {
		return s.value[i] < s.value[j]
	})
	return s 
}

// All 是否所有元素满足条件
func (s *IntSlice) All(fn func(int, int) bool) bool {
	for i, each := range s.value {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s *IntSlice) Any(fn func(int, int) bool) bool {
	for i, each := range s.value {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s *IntSlice) Paginate(size int) [][]int {
	if size <= 0 {
		size = 1
	}
	var pages [][]int
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
func (s *IntSlice) Preappend(given int) *IntSlice {
	value := make([]int, 0, len(s.value)+1)
	value = append(value, given)
	s.value = append(value, s.value...)
	return s
}

// Max 获取最大元素
func (s *IntSlice) Max(value *int) error {
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
func (s *IntSlice) Min(value *int) error {
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
func (s *IntSlice) Random(value *int) error {
	if len(s.value) <= 0 {
		return errors.New("empty")
	}
	n := rand.Intn(len(s.value))
	*value = s.value[n]
	return nil
}

// Shuffle 打乱列表
func (s *IntSlice) Shuffle() *IntSlice {
	if len(s.value) <= 0 {
		return s
	}

	rand.Shuffle(len(s.value), func(i, j int) {
		s.value[i], s.value[j] = s.value[j], s.value[i] 
	})
	
	return s
}

// Collect 获取列表
func (s *IntSlice) Collect() []int {
	return s.value
}
