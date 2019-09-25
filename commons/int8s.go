package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Int8Slice int8的Slice
type Int8Slice []int8

// Concat 拼接
func (s Int8Slice) Concat(given []int8) Int8Slice {
	value := make([]int8, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Int8Slice(value)
}

// Drop 丢弃前n个
func (s Int8Slice) Drop(n int) Int8Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Int8Slice(s[n:])
}

// Filter 过滤
func (s Int8Slice) Filter(fn func(int, int8) bool) Int8Slice {
	value := make([]int8, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Int8Slice(value)
}

// First 获取第一个元素
func (s Int8Slice) First() (int8, error) {
	if len(s) <= 0 {
		var defaultReturn int8
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Int8Slice) Last() (int8, error) {
	if len(s) <= 0 {
		var defaultReturn int8
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Int8Slice) Map(fn func(int, int8) int8) Int8Slice {
	value := make([]int8, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Int8Slice(value)
}

// Reduce reduce
func (s Int8Slice) Reduce(fn func(int8, int8, int) int8, initial int8) int8 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Int8Slice) Reverse() Int8Slice {
	value := make([]int8, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Int8Slice(value)
}

// Unique 唯一
func (s Int8Slice) Unique() Int8Slice {
	value := make([]int8, 0, len(s))
	seen := make(map[int8]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Int8Slice(value)
}

// Append 在尾部添加
func (s Int8Slice) Append(given int8) Int8Slice {
	return append(s, given)
}

// Len 获取长度
func (s Int8Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Int8Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s Int8Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Int8Slice) Sort() Int8Slice {
	value := make([]int8, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Int8Slice(value)
}

// All 是否所有元素满足条件
func (s Int8Slice) All(fn func(int, int8) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Int8Slice) Any(fn func(int, int8) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Int8Slice) Paginate(size int) [][]int8 {
	if size <= 0 {
		size = 1
	}
	var pages [][]int8
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
func (s Int8Slice) Preappend(given int8) Int8Slice {
	value := make([]int8, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s Int8Slice) Max() (int8, error) {
	if len(s) <= 0 {
		var defaultReturn int8
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
func (s Int8Slice) Min() (int8, error) {
	if len(s) <= 0 {
		var defaultReturn int8
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
func (s Int8Slice) Random() (int8, error) {
	if len(s) <= 0 {
		var defaultReturn int8
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]int8
func (s Int8Slice) Shuffle() Int8Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]int8, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Int8Slice(value)
}

// Collect 获取[]int8
func (s Int8Slice) Collect() []int8 {
	return s
}
