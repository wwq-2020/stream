package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint64Slice uint64的Slice
type Uint64Slice []uint64

// Concat 拼接
func (s Uint64Slice) Concat(given []uint64) Uint64Slice {
	value := make([]uint64, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Uint64Slice(value)
}

// Drop 丢弃前n个
func (s Uint64Slice) Drop(n int) Uint64Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Uint64Slice(s[n:])
}

// Filter 过滤
func (s Uint64Slice) Filter(fn func(int, uint64) bool) Uint64Slice {
	value := make([]uint64, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Uint64Slice(value)
}

// First 获取第一个元素
func (s Uint64Slice) First() (uint64, error) {
	if len(s) <= 0 {
		var defaultReturn uint64
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Uint64Slice) Last() (uint64, error) {
	if len(s) <= 0 {
		var defaultReturn uint64
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Uint64Slice) Map(fn func(int, uint64) uint64) Uint64Slice {
	value := make([]uint64, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Uint64Slice(value)
}

// Reduce reduce
func (s Uint64Slice) Reduce(fn func(uint64, uint64, int) uint64, initial uint64) uint64 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Uint64Slice) Reverse() Uint64Slice {
	value := make([]uint64, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Uint64Slice(value)
}

// Unique 唯一
func (s Uint64Slice) Unique() Uint64Slice {
	value := make([]uint64, 0, len(s))
	seen := make(map[uint64]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Uint64Slice(value)
}

// Append 在尾部添加
func (s Uint64Slice) Append(given uint64) Uint64Slice {
	return append(s, given)
}

// Len 获取长度
func (s Uint64Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Uint64Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s Uint64Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Uint64Slice) Sort() Uint64Slice {
	value := make([]uint64, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Uint64Slice(value)
}

// All 是否所有元素满足条件
func (s Uint64Slice) All(fn func(int, uint64) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Uint64Slice) Any(fn func(int, uint64) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Uint64Slice) Paginate(size int) [][]uint64 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint64
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
func (s Uint64Slice) Preappend(given uint64) Uint64Slice {
	value := make([]uint64, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return Uint64Slice(value)
}

// Max 获取最大元素
func (s Uint64Slice) Max() (uint64, error) {
	if len(s) <= 0 {
		var defaultReturn uint64
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
func (s Uint64Slice) Min() (uint64, error) {
	if len(s) <= 0 {
		var defaultReturn uint64
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
func (s Uint64Slice) Random() (uint64, error) {
	if len(s) <= 0 {
		var defaultReturn uint64
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]uint64
func (s Uint64Slice) Shuffle() Uint64Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]uint64, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Uint64Slice(value)
}

// Collect 获取[]uint64
func (s Uint64Slice) Collect() []uint64 {
	return s
}
