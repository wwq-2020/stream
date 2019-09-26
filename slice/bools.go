package commons
	
import (
	"errors"
	"math/rand"
	"sort"
)

// BoolSlice bool的Slice
type BoolSlice []bool

// Concat 拼接
func (s BoolSlice) Concat(given []bool) BoolSlice {
	value := make([]bool, len(s)+len(given))
	copy(value, s)
	copy(value[len(s):], given)
	return BoolSlice(value)
}

// Drop 丢弃前n个
func (s BoolSlice) Drop(n int) BoolSlice {
	if n < 0 {
		n = 0
	}
	l := len(s) - n
	if l < 0 {
		n = len(s)
	}
	return BoolSlice(s[n:])
}

// Filter 过滤
func (s BoolSlice) Filter(fn func(int, bool) bool) BoolSlice {
	value := make([]bool, 0, len(s))
	for i, each := range s {
		if fn(i, each) {
			value = append(value, each)
		}
	}
	return BoolSlice(value)
}

// First 获取第一个元素
func (s BoolSlice) First() (bool, error) {
	if len(s) <= 0 {
		var defaultReturn bool
		return defaultReturn, errors.New("empty")
	}
	return s[0], nil
}

// Last 获取最后一个元素
func (s BoolSlice) Last() (bool, error) {
	if len(s) <= 0 {
		var defaultReturn bool
		return defaultReturn, errors.New("empty")
	}
	return s[len(s)-1], nil
}

// Map 对每个元素进行操作
func (s BoolSlice) Map(fn func(int, bool) bool) BoolSlice {
	value := make([]bool, len(s))
	for i, each := range s {
		value[i] = fn(i, each)
	}
	return BoolSlice(value)
}

// Reduce reduce
func (s BoolSlice) Reduce(fn func(bool, bool, int) bool, initial bool) bool {
	final := initial
	for i, each := range s {
		final = fn(final, each, i)
	}
	return final
}

// Reverse 逆序
func (s BoolSlice) Reverse() BoolSlice {
	value := make([]bool, len(s))
	for i, each := range s {
		value[len(s)-1-i] = each
	}
	return BoolSlice(value)
}

// Unique 唯一
func (s BoolSlice) Unique() BoolSlice {
	value := make([]bool, 0, len(s))
	seen := make(map[bool]struct{})
	for _, each := range s {
		if _, exist := seen[each]; exist {
			continue
		}		
		seen[each] = struct{}{}
		value = append(value, each)			
	}
	return BoolSlice(value)
}

// Append 在尾部添加
func (s BoolSlice) Append(given bool) BoolSlice {
	return append(s, given)
}

// Len 获取长度
func (s BoolSlice) Len() int {
	return len(s)
}

// IsEmpty 判断是否为空
func (s BoolSlice) IsEmpty() bool {
	return len(s) == 0
}

// IsNotEmpty 判断是否非空
func (s BoolSlice) IsNotEmpty() bool {
	return len(s) != 0
}

// Sort 排序
func (s BoolSlice) Sort() BoolSlice {
	value := make([]bool, len(s))
	copy(value, s)
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return BoolSlice(value)
}

// All 是否所有元素满足条件
func (s BoolSlice) All(fn func(int, bool) bool) bool {
	for i, each := range s {
		if !fn(i, each) {
			return false
		}
	}
	return true
}

// Any 是否有元素满足条件
func (s BoolSlice) Any(fn func(int, bool) bool) bool {
	for i, each := range s {
		if fn(i, each) {
			return true
		}
	}
	return false
}

// Paginate 分页
func (s BoolSlice) Paginate(size int) [][]bool {
	if size <= 0 {
		size = 1
	}
	var pages [][]bool
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
func (s BoolSlice) Preappend(given bool) BoolSlice {
	value := make([]bool, len(s)+1)
	value[0] = given
	copy(value[1:], s)
	return BoolSlice(value)
}

// Max 获取最大元素
func (s BoolSlice) Max() (bool, error) {
	if len(s) <= 0 {
		var defaultReturn bool
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
func (s BoolSlice) Min() (bool, error) {
	if len(s) <= 0 {
		var defaultReturn bool
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
func (s BoolSlice) Random() (bool, error) {
	if len(s) <= 0 {
		var defaultReturn bool
		return defaultReturn, errors.New("empty")
	}
	n := rand.Intn(len(s))
	return s[n], nil
}

// Shuffle 打乱[]bool
func (s BoolSlice) Shuffle() BoolSlice {
	if len(s) <= 0 {
		return s
	}
	value := make([]bool, len(s))
	copy(value, s)
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i] 
	})
	return BoolSlice(value)
}

// Collect 获取[]bool
func (s BoolSlice) Collect() []bool {
	return s
}
