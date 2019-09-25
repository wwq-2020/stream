package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// UintSlice uint的Slice
type UintSlice []uint

// Concat 拼接
func (s UintSlice) Concat(given []uint) UintSlice {
	value := make([]uint, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return UintSlice(value)
}

// Drop 丢弃前n个
func (s UintSlice) Drop(n int) UintSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return UintSlice(s[n:])
}

// Filter 过滤
func (s UintSlice) Filter(fn func(int, uint) bool) UintSlice {
	value := make([]uint, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return UintSlice(value)
}

// First 获取第一个元素
func (s UintSlice) First() (uint, error) {
	if len(s) <= 0 {
		var defaultReturn uint
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s UintSlice) Last() (uint, error) {
	if len(s) <= 0 {
		var defaultReturn uint
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s UintSlice) Map(fn func(int, uint) uint) UintSlice {
	value := make([]uint, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return UintSlice(value)
}

// Reduce reduce
func (s UintSlice) Reduce(fn func(uint, uint, int) uint, initial uint) uint {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s UintSlice) Reverse() UintSlice {
	value := make([]uint, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return UintSlice(value)
}

// Unique 唯一
func (s UintSlice) Unique() UintSlice {
	value := make([]uint, 0, len(s))
	seen := make(map[uint]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return UintSlice(value)
}

// Append 在尾部添加
func (s UintSlice) Append(given uint) UintSlice {
	return append(s, given)
}

// Len 获取长度
func (s UintSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s UintSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsEmpty 判断是否非空
func (s UintSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s UintSlice) Sort() UintSlice {
	value := make([]uint, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return UintSlice(value)
}

// All 是否所有元素满足条件
func (s UintSlice) All(fn func(int, uint) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s UintSlice) Any(fn func(int, uint) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s UintSlice) Paginate(size int) [][]uint {
	if size <= 0 {
		size = 1
	}
	var pages [][]uint
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
func (s UintSlice) Preappend(given uint) UintSlice {
	value := make([]uint, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return s
}

// Max 获取最大元素
func (s UintSlice) Max() (uint, error) {
	if len(s) <= 0 {
		var defaultReturn uint
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
func (s UintSlice) Min() (uint, error) {
	if len(s) <= 0 {
		var defaultReturn uint
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
func (s UintSlice) Random() (uint, error) {
	if len(s) <= 0 {
		var defaultReturn uint
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]uint
func (s UintSlice) Shuffle() UintSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]uint, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return UintSlice(value)
}

// Collect 获取[]uint
func (s UintSlice) Collect() []uint {
	return s
}
