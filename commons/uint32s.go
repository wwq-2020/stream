package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// Uint32Slice uint32的Slice
type Uint32Slice []uint32

// Concat 拼接
func (s Uint32Slice) Concat(given []uint32) Uint32Slice {
	value := make([]uint32, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return Uint32Slice(value)
}

// Drop 丢弃前n个
func (s Uint32Slice) Drop(n int) Uint32Slice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return Uint32Slice(s[n:])
}

// Filter 过滤
func (s Uint32Slice) Filter(fn func(int, uint32) bool) Uint32Slice {
	value := make([]uint32, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return Uint32Slice(value)
}

// First 获取第一个元素
func (s Uint32Slice) First() (uint32, error) {
	if len(s) <= 0 {
		var defaultReturn uint32
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s Uint32Slice) Last() (uint32, error) {
	if len(s) <= 0 {
		var defaultReturn uint32
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s Uint32Slice) Map(fn func(int, uint32) uint32) Uint32Slice {
	value := make([]uint32, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return Uint32Slice(value)
}

// Reduce reduce
func (s Uint32Slice) Reduce(fn func(uint32, uint32, int) uint32, initial uint32) uint32 {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s Uint32Slice) Reverse() Uint32Slice {
	value := make([]uint32, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return Uint32Slice(value)
}

// Unique 唯一
func (s Uint32Slice) Unique() Uint32Slice {
	value := make([]uint32, 0, len(s))
	seen := make(map[uint32]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return Uint32Slice(value)
}

// Append 在尾部添加
func (s Uint32Slice) Append(given uint32) Uint32Slice {
	return append(s, given)
}

// Len 获取长度
func (s Uint32Slice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s Uint32Slice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s Uint32Slice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s Uint32Slice) Sort() Uint32Slice {
	value := make([]uint32, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return Uint32Slice(value)
}

// All 是否所有元素满足条件
func (s Uint32Slice) All(fn func(int, uint32) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s Uint32Slice) Any(fn func(int, uint32) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s Uint32Slice) Paginate(size int) [][]uint32 {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint32
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
func (s Uint32Slice) Preappend(given uint32) Uint32Slice {
	value := make([]uint32, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s Uint32Slice) Max() (uint32, error) {
	if len(s) <= 0 {
		var defaultReturn uint32
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
func (s Uint32Slice) Min() (uint32, error) {
	if len(s) <= 0 {
		var defaultReturn uint32
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
func (s Uint32Slice) Random() (uint32, error) {
	if len(s) <= 0 {
		var defaultReturn uint32
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]uint32
func (s Uint32Slice) Shuffle() Uint32Slice {
	if len(s) <= 0 {
		return s
	}
	value := make([]uint32, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return Uint32Slice(value)
}

// Collect 获取[]uint32
func (s Uint32Slice) Collect() []uint32 {
	return s
}
