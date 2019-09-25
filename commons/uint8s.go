package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint8Slice uint8的Slice
type Uint8Slice []uint8

// Concat 拼接
func (s Uint8Slice) Concat(given []uint8) Uint8Slice {
	value := make([]uint8, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Uint8Slice(value)
}

// Drop 丢弃前n个
func (s Uint8Slice) Drop(n int) Uint8Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Uint8Slice(s[n:])
}

// Filter 过滤
func (s Uint8Slice) Filter(fn func(int, uint8) bool) Uint8Slice {
	value := make([]uint8, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Uint8Slice(value)
}

// First 获取第一个元素
func (s Uint8Slice) First() (uint8, error) {
	if len(s) <= 0 {
		var defaultReturn uint8
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Uint8Slice) Last() (uint8, error) {
	if len(s) <= 0 {
		var defaultReturn uint8
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Uint8Slice) Map(fn func(int, uint8) uint8) Uint8Slice {
	value := make([]uint8, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Uint8Slice(value)
}

// Reduce reduce
func (s Uint8Slice) Reduce(fn func(uint8, uint8, int) uint8, initial uint8) uint8 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Uint8Slice) Reverse() Uint8Slice {
	value := make([]uint8, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Uint8Slice(value)
}

// Unique 唯一
func (s Uint8Slice) Unique() Uint8Slice {
	value := make([]uint8, 0, len(s))
	seen := make(map[uint8]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Uint8Slice(value)
}

// Append 在尾部添加
func (s Uint8Slice) Append(given uint8) Uint8Slice {
	return append(s, given)
}

// Len 获取长度
func (s Uint8Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Uint8Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s Uint8Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Uint8Slice) Sort() Uint8Slice {
	value := make([]uint8, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Uint8Slice(value)
}

// All 是否所有元素满足条件
func (s Uint8Slice) All(fn func(int, uint8) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Uint8Slice) Any(fn func(int, uint8) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Uint8Slice) Paginate(size int) [][]uint8 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint8
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
func (s Uint8Slice) Preappend(given uint8) Uint8Slice {
	value := make([]uint8, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s Uint8Slice) Max() (uint8, error) {
	if len(s) <= 0 {
		var defaultReturn uint8
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
func (s Uint8Slice) Min() (uint8, error) {
	if len(s) <= 0 {
		var defaultReturn uint8
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
func (s Uint8Slice) Random() (uint8, error) {
	if len(s) <= 0 {
		var defaultReturn uint8
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]uint8
func (s Uint8Slice) Shuffle() Uint8Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]uint8, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Uint8Slice(value)
}

// Collect 获取[]uint8
func (s Uint8Slice) Collect() []uint8 {
	return s
}
