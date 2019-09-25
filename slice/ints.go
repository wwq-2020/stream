package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// IntSlice int的Slice
type IntSlice []int

// Concat 拼接
func (s IntSlice) Concat(given []int) IntSlice {
	value := make([]int, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return IntSlice(value)
}

// Drop 丢弃前n个
func (s IntSlice) Drop(n int) IntSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return IntSlice(s[n:])
}

// Filter 过滤
func (s IntSlice) Filter(fn func(int, int) bool) IntSlice {
	value := make([]int, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return IntSlice(value)
}

// First 获取第一个元素
func (s IntSlice) First() (int, error) {
	if len(s) <= 0 {
		var defaultReturn int
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s IntSlice) Last() (int, error) {
	if len(s) <= 0 {
		var defaultReturn int
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s IntSlice) Map(fn func(int, int) int) IntSlice {
	value := make([]int, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return IntSlice(value)
}

// Reduce reduce
func (s IntSlice) Reduce(fn func(int, int, int) int, initial int) int {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s IntSlice) Reverse() IntSlice {
	value := make([]int, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return IntSlice(value)
}

// Unique 唯一
func (s IntSlice) Unique() IntSlice {
	value := make([]int, 0, len(s))
	seen := make(map[int]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return IntSlice(value)
}

// Append 在尾部添加
func (s IntSlice) Append(given int) IntSlice {
	return append(s, given)
}

// Len 获取长度
func (s IntSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s IntSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s IntSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s IntSlice) Sort() IntSlice {
	value := make([]int, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return IntSlice(value)
}

// All 是否所有元素满足条件
func (s IntSlice) All(fn func(int, int) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s IntSlice) Any(fn func(int, int) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s IntSlice) Paginate(size int) [][]int {
	if size <= 0 {
		size = 1
	}
	var pages [][]int
	prev := -1
	for i := range s {
		if (i-prev) < size && i != (len(s)-1) {
			continue
		}
		pages = append(pages, s[prev+1:i+1])
		prev = i
	}
	return pages
}

// Preappend 在首部添加元素
func (s IntSlice) Preappend(given int) IntSlice {
	value := make([]int, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return IntSlice(value)
}

// Max 获取最大元素
func (s IntSlice) Max() (int, error) {
	if len(s) <= 0 {
		var defaultReturn int
		return defaultReturn, errors.New("empty")
	}
	max := s[0]
	for _, each := range s {
		if max < each {
			max = each
		}
	}
	return max, nil 
}

// Min 获取最小元素
func (s IntSlice) Min() (int, error) {
	if len(s) <= 0 {
		var defaultReturn int
		return defaultReturn, errors.New("empty")
	}
	min := s[0]
	for _, each := range s {
		if each < min {
			min = each
		}
	}
	return min, nil
}

// Random 随机获取一个元素
func (s IntSlice) Random() (int, error) {
	if len(s) <= 0 {
		var defaultReturn int
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]int
func (s IntSlice) Shuffle() IntSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]int, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return IntSlice(value)
}

// Collect 获取[]int
func (s IntSlice) Collect() []int {
	return s
}
